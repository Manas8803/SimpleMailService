# 📧 Simple Mail Service App 📧

This is a serverless application that provides a simple mail service functionality. It is built using AWS Lambda and is designed to be deployed on the cloud. 🚀

## 🔥 Features

- Send plain text emails 📃
- Send HTML-formatted emails 💻

## 🏗️ Architecture

The application is built using a serverless architecture, leveraging AWS Lambda functions and other AWS services. The primary components are:

1. **Lambda Function**: The core functionality of the application is implemented in an AWS Lambda function, written in Golang. ⚡
2. **API Gateway**: AWS API Gateway is used to expose the Lambda function as a RESTful API endpoint, allowing clients to trigger email sending. 🌐

## 🚀 Deployment

The application is deployed using the AWS Serverless Application Model (SAM) and the AWS CloudFormation service. The deployment process is automated using a CI/CD pipeline. 🛠️

## 🎯 Usage

To use the Simple Mail Service App, you can send HTTPS requests to the API Gateway endpoint. The API supports the following operation:

- `POST /`: Send a plain text or HTML-formatted email. 📤

### Request Format
```json
{
 "email": "<recipient_email>",
 "message": {
   "subject": "Hello from API",
   "body": "<h1>This is the email body content.</h1>"
 }
}
```

### Response Format
```json
{
 "message": "Email sent successfully or Error message"
}
