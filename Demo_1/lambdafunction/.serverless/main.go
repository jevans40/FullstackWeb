package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type AuthEvent struct {
	AuthToken string `json:"access_token"`
	Message   string `json:"message"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event AuthEvent) (Response, error) {
	var buf bytes.Buffer
	var body []byte

	auth, name, err := AuthHandler(event.AuthToken)
	if auth {
		body, err = json.Marshal(map[string]interface{}{
			"message": name + " says " + event.Message,
		})
		if err != nil {
			return Response{StatusCode: 404,
				Headers: map[string]string{
					"Content-Type":                 "application/json",
					"X-MyCompany-Func-Reply":       "world-handler",
					"Access-Control-Allow-Origin":  "*",
					"Access-Control-Allow-Headers": "*",
					"Access-Control-Allow-Methods": "*",
				}}, err
		}
	}
	if err != nil {
		return Response{StatusCode: 404,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"X-MyCompany-Func-Reply":       "world-handler",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "*",
				"Access-Control-Allow-Methods": "*",
			}}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"X-MyCompany-Func-Reply":       "world-handler",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
			"Access-Control-Allow-Methods": "*",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
