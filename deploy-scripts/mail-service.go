package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type MailServiceStackProps struct {
	awscdk.StackProps
}

func NewMailServiceStack(scope constructs.Construct, id string, props *MailServiceStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	sendEmail_handler := awslambda.NewFunction(stack, jsii.String("MailService-app"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../app"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"EMAIL":    jsii.String(os.Getenv("EMAIL")),
			"PASSWORD": jsii.String(os.Getenv("PASSWORD")),
		},
		FunctionName: jsii.String("MailService-Send_Mail_Lambda"),
	})
	awsapigateway.NewLambdaRestApi(stack, jsii.String("MailService-Api-Gateway"), &awsapigateway.LambdaRestApiProps{
		Handler: sendEmail_handler,
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewMailServiceStack(app, "MailServiceStack", &MailServiceStackProps{
		awscdk.StackProps{
			Env:       env(),
			StackName: jsii.String("MailServiceStack"),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file : ", err)
	}
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
