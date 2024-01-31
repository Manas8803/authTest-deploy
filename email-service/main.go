// main.go

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/smtp"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var lambdaEvent LambdaEvent

	err := json.Unmarshal([]byte(event.Body), &lambdaEvent)
	if err != nil {
		log.Println("Error decoding JSON payload:", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Bad Request"}, nil
	}

	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")
	to := []string{lambdaEvent.Email}
	message := []byte("To: " + lambdaEvent.Email + "\r\n" +
		"Subject: OTP for Verification\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
		"<html><body>" +
		"<h1>Your OTP for Verification is <strong>" + lambdaEvent.OTP + "</strong></h1>" +
		"</body></html>")

	err = smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, message)
	if err != nil {
		log.Println("Error in sending OTP:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "OTP sent successfully"}, nil
}

func main() {
	lambda.Start(Handler)
}
