package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Request struct {
	Email   string  `json:"email"`
	Content Message `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

var cors_headers = map[string]string{
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Headers": "Content-Type",
	"Access-Control-Allow-Methods": "OPTIONS,GET,PUT,POST,DELETE",
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(request.Body) == 0 {
		resp := Response{
			Message: "Request body cannot be empty",
		}
		respBody, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    cors_headers,
			Body:       string(respBody),
		}, nil
	}

	var temp map[string]interface{}
	err := json.Unmarshal([]byte(request.Body), &temp)
	if err != nil {
		resp := Response{
			Message: "Internal Server Error: " + err.Error(),
		}
		respBody, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    cors_headers,
			Body:       string(respBody),
		}, nil
	}

	if _, ok := temp["email"].(string); !ok {
		resp := Response{
			Message: "Email cannot be empty",
		}
		respBody, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    cors_headers,
			Body:       string(respBody),
		}, nil
	}

	if _, ok := temp["message"]; !ok {
		resp := Response{
			Message: "Message cannot be empty",
		}
		respBody, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    cors_headers,
			Body:       string(respBody),
		}, nil
	}

	var payload Request
	err = json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		resp := Response{
			Message: "Internal Server Error: " + err.Error(),
		}
		respBody, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    cors_headers,
			Body:       string(respBody),
		}, nil
	}

	if strings.Trim(payload.Content.Subject, " ") == "" || strings.Trim(payload.Content.Body, " ") == "" {
		resp := Response{
			Message: "Content body and subject cannot be empty",
		}
		respBody, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    cors_headers,
			Body:       string(respBody),
		}, nil
	}

	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")
	to := []string{payload.Email}
	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n"+
		"<html><body style=\"color:green;\">%s</body></html>",
		payload.Email, payload.Content.Subject, payload.Content.Body))

	if err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, message); err != nil {
		log.Println("Error in sending email:", err)
		resp := Response{
			Message: "Internal Server Error: " + err.Error(),
		}
		respBody, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    cors_headers,
			Body:       string(respBody),
		}, nil
	}

	resp := Response{
		Message: "Email sent successfully",
	}
	respBody, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    cors_headers,
		Body:       string(respBody),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
