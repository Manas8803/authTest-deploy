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

type AuthTestProps struct {
	awscdk.StackProps
}

func LamdaStack(scope constructs.Construct, id string, props *AuthTestProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	sendEmail_handler := awslambda.NewFunction(stack, jsii.String("SendEmail"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("email-service.zip"), nil),
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("/email-service/build/main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"EMAIL":    jsii.String(os.Getenv("EMAIL")),
			"PASSWORD": jsii.String(os.Getenv("PASSWORD")),
		},
	})

	auth_handler := awslambda.NewFunction(stack, jsii.String("Auth"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("main.zip"), nil),
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("/auth-service/build/main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"SQLURI":            jsii.String(os.Getenv("SQLURI")),
			"JWT_SECRET_KEY":    jsii.String(os.Getenv("JWT_SECRET_KEY")),
			"JWT_LIFETIME":      jsii.String(os.Getenv("JWT_LIFETIME")),
			"SEND_TO_EMAIL_ARN": jsii.String(*sendEmail_handler.FunctionArn()),
		},
	})

	awsapigateway.NewLambdaRestApi(stack, jsii.String("authTest"), &awsapigateway.LambdaRestApiProps{
		Handler: auth_handler,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	LamdaStack(app, "AuthTest-Stack", &AuthTestProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

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
