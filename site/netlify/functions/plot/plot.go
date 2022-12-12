package main

import (
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
	"c-": 1.0,
	"c":  2.0,
	"c+": 3.0,
	"c ": 3.0,
	"b-": 4.0,
	"b":  5.0,
	"b+": 6.0,
	"b ": 6.0,
	"a-": 7.0,
	"a":  8.0,
	"a+": 9.0,
	"a ": 9.0,
}

const topScoreName = "a+"

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// request.QueryStringParameters https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go

	// help request
	if len(request.QueryStringParameters) == 0 {
		helpMsg := `Usage: /plot?name=<name>&[legend=<legend,...>]&<skill>=<score[,score,...]>[&<skill>=<score[,score,...]>]...

<name> - the name of the person, e.g. "Amy Pond";
<skill> - the name of the skill, e.g. "quality";
<score> - the score of the skill, e.g. "A-", "A", "A+", "B-", "B", "B+", "C", "C+", "C-";
          multiple comma separated scores are supported, e.g. "A-,B+";
          the score is case insensitive, e.g. "a+" and "A+" are the same;
<legend> - (optional) the legend of the chart, if multiple comma separated scores are provided, e.g. "Q1 2022,Q2 2022"
		`
		return &events.APIGatewayProxyResponse{
			StatusCode:      400,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            helpMsg,
			IsBase64Encoded: false,
		}, nil
	}

	name, skills, scores, legend := dataExtractor(request.QueryStringParameters)
	data, err := generateSVG(name, skills, scores, legend)

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode:      503,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            err.Error(),
			IsBase64Encoded: false,
		}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "image/svg+xml"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func generateSVG(name string, skills []string, scores [][]float64, legend []string) ([]byte, error) {
	topScore := make([]float64, len(skills))
	for i := 0; i < len(skills); i++ {
		topScore[i] = evaluationMap[topScoreName]
	}

	options := make([]charts.OptionFunc, 0)
	options = append(
		options,
		charts.TitleTextOptionFunc(name),
		charts.RadarIndicatorOptionFunc(skills, topScore),
		charts.SVGTypeOption(),
	)
	if len(legend) > 0 {
		options = append(
			options,
			charts.LegendLabelsOptionFunc(legend),
		)
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
func dataExtractor(data map[string]string) (name string, skills []string, scores [][]float64, legend []string) {
	name = "Unknown person"
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
