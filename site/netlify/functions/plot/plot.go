package main

import (
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

	for key, value := range data {
		switch key {
		case "name":
			name = value
		case "legend":
			legend = strings.Split(value, ",")
		default:
			skills = append(skills, cases.Title(language.English, cases.NoLower).String(key))
			evaluations = append(evaluations, value)
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
