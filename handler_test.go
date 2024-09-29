package main

import "testing"

func TestHandler(t *testing.T) {
	var requestId = "a2fe3b91-4efe-4a59-b916-ce1b6b004142"
	event := APIGatewayRequest{
		RequestContext: RequestContext{
			Identity: Identity{
				SourceIp:  "2a02:6b8:c02:900:0:f822:0:a8",
				UserAgent: "PostmanRuntime/7.28.1",
			},
			HTTPMethod:       "GET",
			RequestId:        requestId,
			RequestTime:      "24/Aug/2021:17:48:59 +0000",
			RequestTimeEpoch: 1_629_827_339,
			APIGateway: APIGateway{
				OperationContext: map[string]string{
					"name": "world",
				},
			},
		},
		HTTPMethod:      "GET",
		Path:            "/world",
		IsBase64Encoded: false,
	}
	ctx := ContextExample

	request, err := Handler(ctx, &event)

	if err != nil {
		t.Error("error is not nil")
	}

	assertEqual(t, request.StatusCode, 200)
	assertEqual(t, request.Body, "Hello, world!")
}

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("actual: <%s>, expected: <%s>", actual, expected)
	}
}