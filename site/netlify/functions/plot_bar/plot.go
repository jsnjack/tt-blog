package main

import (
	"encoding/base64"
	"errors"
	"sort"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/vicanso/go-charts/v2"
)

const helpMsg = `Usage: /plot2?type=<{bar}>&[title=<str>]&[legend=<str1[,...]>]&<key1>=<val1[,val2,...]>...&[format=<{png,svg}>]

<type> - type of the chart. Can be one of: bar
<title> - (optional) title of the chart
<legend> - (optional) the legend of the chart, used when multiple comma separated values are porvided for keys
<key> - name of a key which is used in the chart. Supports commas separated list of values
<format> - (optional) image format. Default is svg. Supported values are svg, png

Examples:
`

type ChartData struct {
	Title  string
	Legend []string
	Keys   []string
	Values []float64
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
		Legend: make([]string, 0),
		Keys:   make([]string, 0),
		Values: make([]float64, 0),
		Format: "svg",
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
		}
	}
	return cd, nil
}

func generateImage(cd *ChartData) ([]byte, error) {
	options := make([]charts.OptionFunc, 0)
	options = append(
		options,
		charts.TitleTextOptionFunc(cd.Title),
	)
	if len(cd.Legend) > 0 {
		options = append(options, charts.LegendLabelsOptionFunc(cd.Legend))
	}

	if cd.Format == "svg" {
		options = append(options, charts.SVGTypeOption())
	}

	switch cd.Type {
	case "bar":
		options = append(options, charts.XAxisDataOptionFunc(cd.Keys))
		p, err := charts.BarRender(
			cd.Values,
			options...,
		)

	default:
		return nil, errors.New("unsupported type: " + cd.Type)

	}

	if err != nil {
		return nil, err
	}

	buf, err := p.Bytes()
	if err != nil {
		return nil, err
	}

	// p, err := charts.BarRender(
	// 	values,
	// 	charts.XAxisDataOptionFunc([]string{
	// 		"Jan",
	// 		"Feb",
	// 		"Mar",
	// 		"Apr",
	// 		"May",
	// 		"Jun",
	// 		"Jul",
	// 		"Aug",
	// 		"Sep",
	// 		"Oct",
	// 		"Nov",
	// 		"Dec",
	// 	}),
	// 	charts.LegendLabelsOptionFunc([]string{
	// 		"Rainfall",
	// 		"Evaporation",
	// 	}, charts.PositionRight),
	// 	charts.MarkLineOptionFunc(0, charts.SeriesMarkDataTypeAverage),
	// 	charts.MarkPointOptionFunc(0, charts.SeriesMarkDataTypeMax,
	// 		charts.SeriesMarkDataTypeMin),
	// 	// custom option func
	// 	func(opt *charts.ChartOption) {
	// 		opt.SeriesList[1].MarkPoint = charts.NewMarkPoint(
	// 			charts.SeriesMarkDataTypeMax,
	// 			charts.SeriesMarkDataTypeMin,
	// 		)
	// 		opt.SeriesList[1].MarkLine = charts.NewMarkLine(
	// 			charts.SeriesMarkDataTypeAverage,
	// 		)
	// 	},
	// )

	return buf, nil
}
