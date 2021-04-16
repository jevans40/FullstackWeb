package main

import (
	"context"

	"github.com/jevans40/FullstackWeb/Demo_1/lambdafunction/db"
	"github.com/jevans40/FullstackWeb/Demo_1/lambdafunction/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type CreateProfile struct {
	Artist_name string `json:"artist_name"`
	Likes       int    `json:"likes"`
	Description string `json:"description"`
}

func HandleRequest(ctx context.Context, event CreateProfile) (Response, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	if err != nil {
		return Response{StatusCode: 404,
			IsBase64Encoded: false,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"X-MyCompany-Func-Reply":       "world-handler",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "*",
				"Access-Control-Allow-Methods": "*",
			}}, err
	}
	defer db2.Close()
	db2.AutoMigrate(&models.ProfileModel{})
	account := &models.ProfileModel{}
	account.Artist_name = event.Artist_name
	account.Likes = event.Likes
	account.Picture = []byte{}
	account.Description = event.Description

	db2.Create(account)
	return Response{StatusCode: 200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"X-MyCompany-Func-Reply":       "world-handler",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
			"Access-Control-Allow-Methods": "*",
		}}, err
}
func main() {
	lambda.Start(HandleRequest)
}
