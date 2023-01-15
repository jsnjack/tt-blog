package main

import (
	"encoding/base64"
	"sort"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/vicanso/go-charts/v2"
)

// evaluationMap is a map of evaluation to score. Note, browsers encode + as a space
var evaluationMap = map[string]float64{
	"c":  1.0,
	"b":  2.0,
	"a":  3.0,
	"a+": 4.0,
	"a ": 4.0,
}

const topScoreName = "a+"

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// request.QueryStringParameters https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go

	// help request
	if len(request.QueryStringParameters) == 0 {
		helpMsg := `Usage: /plot?name=<name>&[legend=<legend,...>]&<skill>=<score[,score,...]>[&<skill>=<score[,score,...]>]...

<name> - the name of the person, e.g. "Amy Pond";
<skill> - the name of the skill, e.g. "quality";
<score> - the score of the skill, e.g. "A", "A+", "B", "C";
          multiple comma separated scores are supported, e.g. "A,B";
          the score is case insensitive, e.g. "a+" and "A+" are the same;
<legend> - (optional) the legend of the chart, if multiple comma separated scores are provided, e.g. "Q1 2022,Q2 2022"
<type> - (optional) image format. Default is svg. Supported values are svg, png

Examples:
plot?name=Amy+Pond&ownership=a&structured=a&trust=a&strategic=a+&independence=a&quality=b&effectiveness=b&teamwork=a
plot?name=Amy+Pond&ownership=a,a&structured=a,a+&trust=a,a&strategic=a+,a&independence=a,a&quality=b,a&effectiveness=b,a&teamwork=a,a&legend=2021,2022&type=png
		`
		return &events.APIGatewayProxyResponse{
			StatusCode:      400,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            helpMsg,
			IsBase64Encoded: false,
		}, nil
	}

	name, skills, scores, legend, format := dataExtractor(request.QueryStringParameters)
	data, err := generateImage(name, skills, scores, legend, format)

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode:      503,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            err.Error(),
			IsBase64Encoded: false,
		}, err
	}

	if format == "png" {
		return &events.APIGatewayProxyResponse{
			StatusCode:      200,
			Headers:         map[string]string{"Content-Type": "image/png"},
			Body:            base64.StdEncoding.EncodeToString(data),
			IsBase64Encoded: true,
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "image/svg+xml"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func generateImage(name string, skills []string, scores [][]float64, legend []string, format string) ([]byte, error) {
	topScore := make([]float64, len(skills))
	for i := 0; i < len(skills); i++ {
		topScore[i] = evaluationMap[topScoreName]
	}

	options := make([]charts.OptionFunc, 0)
	options = append(
		options,
		charts.TitleTextOptionFunc(name),
		charts.RadarIndicatorOptionFunc(skills, topScore),
	)
	if len(legend) > 0 {
		options = append(options, charts.LegendLabelsOptionFunc(legend))
	}

	if format == "svg" {
		options = append(options, charts.SVGTypeOption())
	}

	p, err := charts.RadarRender(
		scores,
		options...,
	)

	if err != nil {
		return nil, err
	}

	buf, err := p.Bytes()
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// dataExtractor extracts data from the provided query string parameters
func dataExtractor(data map[string]string) (name string, skills []string, scores [][]float64, legend []string, format string) {
	name = "Unknown person"
	format = "svg"
	skills = make([]string, 0)
	evaluations := make([]string, 0)
	legend = make([]string, 0)

	// Sort the keys to preserve the order of the skills
	keySlice := make([]string, 0)
	for k := range data {
		keySlice = append(keySlice, k)
	}

	sort.Strings(keySlice)

	for _, key := range keySlice {
		switch key {
		case "name":
			name = data[key]
		case "legend":
			legend = strings.Split(data[key], ",")
		case "type":
			format = data[key]
		default:
			skills = append(skills, cases.Title(language.English, cases.NoLower).String(key))
			evaluations = append(evaluations, data[key])
		}
	}

	seriesNumber := 0
	for _, ev := range evaluations {
		num := len(strings.Split(ev, ","))
		if num > seriesNumber {
			seriesNumber = num
		}
	}
	scores = make([][]float64, seriesNumber)

	for _, ev := range evaluations {
		for idx, val := range strings.Split(ev, ",") {
			cleandVal := strings.TrimLeft(strings.ToLower(val), " ")
			scores[idx] = append(scores[idx], evaluationMap[cleandVal])
		}
	}
	return
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
