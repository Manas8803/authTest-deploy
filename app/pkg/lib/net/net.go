package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func SendOtpByEmail(email, otp string) {
	//* Invoke the target Lambda function
	sess := session.Must(session.NewSession())
	client := lambda.New(sess)
	payloadData := map[string]string{"email": email, "otp": otp}
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		log.Println("Error marshaling payload:", err)
	}
	input := &lambda.InvokeInput{
		FunctionName: aws.String("AuthTest-Stack1-SendEmail14C26199-BjmgAo7bktVO"),
		Payload:      payloadBytes,
	}
	log.Println("Payload:", string(payloadBytes))

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
