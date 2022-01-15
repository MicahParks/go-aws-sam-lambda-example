package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/MicahParks/go-aws-sam-lambda-example/util"
)

type lambdaOneHandler struct {
	logger *log.Logger
}

// New creates a new handler for Lambda one.
func New(logger *log.Logger) lambda.Handler {
	return util.NewHandlerV1(lambdaOneHandler{
		logger: logger,
	})
}

// Handle implements util.LambdaHTTPV1 interface. It contains the logic for the handler.
func (handler lambdaOneHandler) Handle(ctx context.Context, _ *events.APIGatewayProxyRequest) (response *events.APIGatewayProxyResponse, err error) {
	response = &events.APIGatewayProxyResponse{}

	select {
	case <-ctx.Done():
		handler.logger.Println("Context expired.")
		err := ctx.Err()
		handler.logger.Printf("Error: %s", err.Error())
		response.StatusCode = http.StatusInternalServerError
		response.Body = "Failed due to expired context on startup."
	default:
		handler.logger.Println("Context not expired.")
		response.StatusCode = http.StatusOK
		response.Body = "Successfully returned response."
	}

	return response, nil
}
