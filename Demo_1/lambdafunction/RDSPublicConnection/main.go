package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/jevans40/FullstackWeb/Demo_1/lambdafunction/models"

	"github.com/jevans40/FullstackWeb/Demo_1/lambdafunction/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type GetProfile struct {
	Name string `json:"Name"`
}

func HandleRequest(ctx context.Context, event GetProfile) (Response, error) {
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
	if event.Name != "" {
		account.Artist_name = event.Name
	}
	var users []models.ProfileModel
	db2.Where(account.Artist_name).Find(&users)
	usersRet, _ := json.Marshal(users)
	print(usersRet)
	var buf bytes.Buffer
	json.HTMLEscape(&buf, usersRet)
	return Response{StatusCode: 200,
		IsBase64Encoded: false,
		Body:            buf.String(),
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
