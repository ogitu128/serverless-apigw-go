package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		println("Usage: go run . <json event> [<json context>]")
		return
	}

	eventJson := os.Args[1]
	event := &APIGatewayRequest{}
	if err := json.Unmarshal([]byte(eventJson), event); err != nil {
		fmt.Printf("Invalid json: %s\n", eventJson)
		return
	}

	var ctx context.Context
	if len(os.Args) == 3 {
		ctxJson := os.Args[2]
		parsedCtx := &ParsedContext{}

		if err := json.Unmarshal([]byte(ctxJson), parsedCtx); err != nil {
			fmt.Printf("Invalid json: %s\n", ctxJson)
			return
		}

		ctx = NewContext(parsedCtx.FunctionName, parsedCtx.FunctionVersion, parsedCtx.MemoryLimitInMb, parsedCtx.RequestID)
	} else {
		ctx = ContextExample
	}

	response, err := Handler(ctx, event)
	if err != nil {
		fmt.Printf("Handler error: %s\n", err)
	}

	result, _ := json.Marshal(response)
	println(string(result))
}