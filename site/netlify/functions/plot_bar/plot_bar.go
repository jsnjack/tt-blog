package main

import (
	"encoding/base64"
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vicanso/go-charts/v2"
)

const helpMsg = `Usage: /plot_bar?type=<{bar}>&[title=<str>]&[legend=<str1[,...]>]&<key1>=<val1[,val2,...]>...&[format=<{png,svg}>]

<type> - type of the chart. Can be one of: bar
<title> - (optional) title of the chart
<legend> - (optional) the legend of the chart. Describe comma separated <values> of <keys>
<key> - name of a key which is used in the chart. Supports commas separated list of values
<format> - (optional) image format. Default is svg. Supported values are svg, png

Examples:
/plot_bar?type=bar&title=Haproxy+response+duration&legend=min,max,p99&haproxy18=10,20,19&haproxy20=9,20,17&format=png
`

type ChartData struct {
	Title  string
	Legend []string
	Keys   []string
	Values [][]float64
	Format string
	Type   string
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// request.QueryStringParameters https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go

	// help request
	if len(request.QueryStringParameters) == 0 {
		return &events.APIGatewayProxyResponse{
			StatusCode:      400,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            helpMsg,
			IsBase64Encoded: false,
		}, nil
	}

	cd, err := parseQuery(request.QueryStringParameters)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode:      400,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            err.Error(),
			IsBase64Encoded: false,
		}, nil
	}
	imgBytes, err := generateImage(cd)

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode:      503,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            err.Error(),
			IsBase64Encoded: false,
		}, err
	}

	if cd.Format == "png" {
		return &events.APIGatewayProxyResponse{
			StatusCode:      200,
			Headers:         map[string]string{"Content-Type": "image/png"},
			Body:            base64.StdEncoding.EncodeToString(imgBytes),
			IsBase64Encoded: true,
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "image/svg+xml"},
		Body:            string(imgBytes),
		IsBase64Encoded: false,
	}, nil
}

// parseQuery extracts data from the provided query string parameters
func parseQuery(data map[string]string) (*ChartData, error) {
	cd := &ChartData{
		Title:  "",
		Legend: make([]string, 0),
		Keys:   make([]string, 0),
		Values: make([][]float64, 0),
		Format: "svg",
		Type:   "bar",
	}

	// Sort the keys to preserve the order of `keys`
	keySlice := make([]string, 0)
	for k := range data {
		keySlice = append(keySlice, k)
	}

	sort.Strings(keySlice)

	for _, key := range keySlice {
		normalizedKey := strings.ToLower(key)
		switch normalizedKey {
		case "title":
			cd.Title = data[key]
		case "legend":
			cd.Legend = strings.Split(data[key], ",")
		case "type":
			cd.Type = data[key]
		case "format":
			cd.Format = data[key]
		default:
			// If the key is not one of the above, add it to the keys list
			cd.Keys = append(cd.Keys, key)
			cd.Values = append(cd.Values, []float64{})
			values := strings.Split(data[key], ",")
			for _, v := range values {
				normalizedValue := strings.TrimSpace(v)
				floatValue, err := strconv.ParseFloat(normalizedValue, 64)
				if err != nil {
					return nil, err
				}
				cd.Values[len(cd.Values)-1] = append(cd.Values[len(cd.Values)-1], floatValue)
			}
		}
	}
	return cd, nil
}

// generateImage generates an image from the provided ChartData
func generateImage(cd *ChartData) ([]byte, error) {
	options := make([]charts.OptionFunc, 0)
	options = append(
		options,
		charts.TitleTextOptionFunc(cd.Title),
	)
	if len(cd.Keys) > 0 {
		options = append(options, charts.LegendOptionFunc(charts.LegendOption{
			Data: cd.Keys,
			Left: "100%",
			Padding: charts.Box{
				Bottom: 30,
				Top:    10,
			},
		}))
	}

	if cd.Format == "svg" {
		options = append(options, charts.SVGTypeOption())
	}

	switch cd.Type {
	case "bar":
		options = append(options, charts.XAxisDataOptionFunc(cd.Legend))
		p, err := charts.BarRender(
			cd.Values,
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
	default:
		return nil, errors.New("unsupported type: " + cd.Type)

	}
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
