package util

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// LambdaHTTPV1 TODO
// This is for AWS Lambdas that are behind API Gateway version 1. (Version 2 is the latest.)
type LambdaHTTPV1 interface {
	Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (response *events.APIGatewayProxyResponse, err error)
}

// LambdaHTTPV2 TODO
// This is for AWS Lambdas that are behind API Gateway version 2.
type LambdaHTTPV2 interface {
	Handle(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (response *events.APIGatewayV2HTTPResponse, err error)
}

type wrappedHandlerV1 struct {
	LambdaHTTPV1
}

type wrappedHandlerV2 struct {
	LambdaHTTPV2
}

// TODO
func NewHandlerV1(typedHandler LambdaHTTPV1) lambda.Handler {
	return wrappedHandlerV1{
		LambdaHTTPV1: typedHandler,
	}
}

// TODO
func NewHandlerV2(typedHandler LambdaHTTPV2) lambda.Handler {
	return wrappedHandlerV2{
		LambdaHTTPV2: typedHandler,
	}
}

// TODO
func (handler wrappedHandlerV1) Invoke(ctx context.Context, request []byte) (response []byte, err error) {
	req := &events.APIGatewayProxyRequest{}
	err = json.Unmarshal(request, req)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal AWS Lambda request: %w", err)
	}

	resp, err := handler.Handle(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("handler failed to handle reqeust: %w", err)
	}

	response, err = json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response into AWS format: %w", err)
	}

	return response, nil
}

// TODO
func (handler wrappedHandlerV2) Invoke(ctx context.Context, request []byte) (response []byte, err error) {
	req := &events.APIGatewayV2HTTPRequest{}
	err = json.Unmarshal(request, req)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal AWS Lambda request: %w", err)
	}

	resp, err := handler.Handle(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("handler failed to handle reqeust: %w", err)
	}

	response, err = json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response into AWS format: %w", err)
	}

	return response, nil
}
