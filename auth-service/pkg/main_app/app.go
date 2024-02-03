package main_app

import (
	_ "auth-service/pkg/main_app/docs"
	"auth-service/pkg/main_app/user/service"
	"auth-service/pkg/storage"
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Registration API
//	@version		1.0
//	@description	This is a registration api for an application.
//	@BasePath		/api/v1

var gorillaAdp *gorillamux.GorillaMuxAdapter

func init() {

	//* Run app
	storage.ConnectDB()

	//* Initialse router
	router := service.SetupRoutes()
	// docs route : api/v1/docs/swagger/index.html#/user
	router.PathPrefix("").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	gorillaAdp = gorillamux.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	SwitchableAPIGatewayRequest_req := core.NewSwitchableAPIGatewayRequestV1(&req) //& Function, not a method.
	SwitchableAPIGatewayResponse_res, err := gorillaAdp.ProxyWithContext(ctx, *SwitchableAPIGatewayRequest_req)
	if err != nil {
		log.Println(err)
	}
	APIGatewayProxyResponse_res := SwitchableAPIGatewayResponse_res.Version1()
	return *APIGatewayProxyResponse_res, nil
}
func Run() {
	lambda.Start(Handler)
}
