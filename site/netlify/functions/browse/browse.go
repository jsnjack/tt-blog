package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// request.QueryStringParameters https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go
	url := request.QueryStringParameters["url"]
	if url == "" {
		return &events.APIGatewayProxyResponse{
			StatusCode:      404,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            "url parameter not found",
			IsBase64Encoded: false,
		}, nil
	}

	// Prepare payload
	payload := map[string]interface{}{"url": url, "headless": true}
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	// Send payload, create a session
	resp, err := http.Post("https://surfly.com/v2/sessions/?api_key="+os.Getenv("SURFLY_API_KEY"), "application/json",
		bytes.NewBuffer(jsonPayload))

	if err != nil {
		return nil, err
	}

	// Parse response
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	// Redirect user to surfly session
	return &events.APIGatewayProxyResponse{
		StatusCode:      201,
		Headers:         map[string]string{"Location": fmt.Sprintf("%v", res["headless_link"])},
		Body:            "",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
