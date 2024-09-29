package main

import (
	"context"
	"fmt"
)

func Handler(ctx context.Context, event *APIGatewayRequest) (*APIGatewayResponse, error) {
	operationContext := event.RequestContext.APIGateway.OperationContext
	name := operationContext["name"]

	return &APIGatewayResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, %s!", name),
	}, nil
}