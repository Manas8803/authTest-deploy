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

type Data struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var payload Data
	log.Println("BODY : ", request.Body)

	if request.Body == "" {
		auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")
		to := []string{"manasahavegeta@gmail.com"}
		message := []byte("To: " + "manasahavegeta@gmail.com" + "\r\n" +
			"Subject: Error \r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
			"<html><body>" +
			"<h1>BODY : <strong>" + string(request.Body) + "</strong></h1>" +
			"</body></html>")
		if err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, message); err != nil {
			log.Println("Error in sending OTP:", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error : " + err.Error(),
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "BAD REQUEST : EMPTY JSON BODY",
		}, nil
	}
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")
	to := []string{payload.Email}
	message := []byte("To: " + payload.Email + "\r\n" +
		"Subject: OTP for Verification\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
		"<html><body>" +
		"<h1>Your OTP for Verification is <strong>" + payload.OTP + "</strong></h1>" +
		"</body></html>")

	if err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, message); err != nil {
		log.Println("Error in sending OTP:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error: " + err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "OTP sent successfully"}, nil
}

func main() {
	lambda.Start(Handler)
}
