package plot

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/vicanso/go-charts/v2"
)

var evaluationMap = map[string]float64{
	"C-": 1.0,
	"C":  2.0,
	"C+": 3.0,
	"B-": 4.0,
	"B":  5.0,
	"B+": 6.0,
	"A-": 7.0,
	"A":  8.0,
	"A+": 9.0,
}

const topScoreName = "A+"

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// request.QueryStringParameters https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go
	name := "Unknown person"
	skills := make([]string, 0)
	values := make([]float64, 0)
	for key, value := range request.QueryStringParameters {
		switch key {
		case "name":
			name = value
		default:
			skills = append(skills, key)
			values = append(values, evaluationMap[value])
		}
	}

	data, err := generateSVG(name, skills, values)
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

func generateSVG(name string, skills []string, values []float64) ([]byte, error) {
	collection := make([][]float64, 0)
	collection = append(collection, values)

	topScore := make([]float64, len(skills))
	for i := 0; i < len(skills); i++ {
		topScore[i] = evaluationMap[topScoreName]
	}

	p, err := charts.RadarRender(
		collection,
		charts.TitleTextOptionFunc(name),
		charts.RadarIndicatorOptionFunc(skills, topScore),
		charts.SVGTypeOption(),
		charts.PaddingOptionFunc(charts.Box{Top: 0, Right: 0, Bottom: 0, Left: 0}),
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

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}