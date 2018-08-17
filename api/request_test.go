package api

import (
	"testing"
	"encoding/json"
)

func TestUnmarshalSqsResource(t *testing.T) {
	requestYaml := `{"resources":[{"name":"test1","type":"AWS::SQS::Queue","properties":{"messagesPerMonth":10,"messageSizeKB":4}}]}`

	var request Request

	err := json.Unmarshal([]byte(requestYaml), &request)
	if err != nil {
		t.Errorf("unexpected error parsing yaml: %s", err)
	}

	if len(request.Resources) != 1 {
		t.Errorf("expected 1 resource, got %d", len(request.Resources))
	}

	resource := request.Resources[0]

	if resource.Name != "test1" {
		t.Errorf("expected 'test1', got %s", resource.Name)
	}

	if resource.Type != SqsResourceType {
		t.Errorf("expected %s, got %s", SqsResourceType, resource.Type)
	}

	sqsProperties, ok := resource.Properties.(SqsProperties)
	if !ok {
		t.Errorf("failed to convert properties to SqsProperties type")
	}

	if sqsProperties.MessagesPerMonth != 10 {
		t.Errorf("expected 10, got %f", sqsProperties.MessagesPerMonth)
	}

	if sqsProperties.MessageSizeKB != 4 {
		t.Errorf("expected 4, got %f", sqsProperties.MessageSizeKB)
	}
}
