package main

import (
	"context"
)

type ParsedContext struct {
	FunctionName    string `json:"functionName"`
	FunctionVersion string `json:"functionVersion"`
	MemoryLimitInMb string `json:"memoryLimitInMB"`
	RequestID       string `json:"requestId"`
}

type FunctionContext struct {
	context.Context
	entries map[interface{}]interface{}
}

func (ctx FunctionContext) Value(key interface{}) interface{} {
	return ctx.entries[key]
}

func NewContext(functionName string, functionVersion string, memoryLimitInMb string, requestID string) context.Context {
	result := FunctionContext{context.Background(), make(map[interface{}]interface{})}

	result.entries["lambdaRuntimeFunctionName"] = functionName
	result.entries["lambdaRuntimeFunctionVersion"] = functionVersion
	result.entries["lambdaRuntimeMemoryLimit"] = memoryLimitInMb
	result.entries["lambdaRuntimeRequestID"] = requestID

	return result
}

var ContextExample = NewContext("d4eo2faf62**********", "d4e3vrugh3**********", "128", "6e8356f9-489b-4c7b-8ba6-c8cd74f25455")