package network

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Data struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
type Payload_Body struct {
	Body string `json:"body"`
}

func SendOtpByEmail(email string, otp string) {

	sess, err := session.NewSession()
	if err != nil {
		log.Println("Error in creating session : ", err.Error())
		return
	}

	client := lambda.New(sess)
	data, err := json.Marshal(Data{Email: email, OTP: otp})
	if err != nil {
		log.Println("Error in marshalling data : ", err.Error())
	}

	body := Payload_Body{Body: string(data)}

	payload, err := json.Marshal(body)
	if err != nil {
		log.Println("Error in marshalling payload : ", err.Error())
	}

	input := &lambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("SEND_TO_EMAIL_ARN")),
		Payload:        payload,
		InvocationType: aws.String("Event"),
	}

	result, err := client.Invoke(input)
	if err != nil {
		log.Println("Error invoking Lambda function:", err)
	} else {
		log.Println("Lambda function invoked successfully:", result)
	}
}

func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding JSON response:", err)
	}
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	response := map[string]string{"error": message}
	RespondWithJSON(w, status, response)
}
