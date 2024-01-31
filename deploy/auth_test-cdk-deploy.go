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

type CdkWorkshopStackProps struct {
	awscdk.StackProps
}

func LamdaStack(scope constructs.Construct, id string, props *CdkWorkshopStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	lambda_func := awslambda.NewFunction(stack, jsii.String("Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("main.zip"), nil),
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("/app/build/main"),
		Environment: &map[string]*string{
			"SQLURI":         jsii.String(os.Getenv("SQLURI")),
			"JWT_SECRET_KEY": jsii.String(os.Getenv("JWT_SECRET_KEY")),
			"JWT_LIFETIME":   jsii.String(os.Getenv("JWT_LIFETIME")),
			"EMAIL":          jsii.String(os.Getenv("EMAIL")),
			"PASSWORD":       jsii.String(os.Getenv("PASSWORD")),
		},
	})

	awsapigateway.NewLambdaRestApi(stack, jsii.String("authTest"), &awsapigateway.LambdaRestApiProps{
		Handler: lambda_func,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	LamdaStack(app, "AuthTest-Stack", &CdkWorkshopStackProps{
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
