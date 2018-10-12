package api

import (
	"encoding/json"
	"testing"
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

func TestUnmarshalEc2Resource(t *testing.T) {
	requestYaml := `{"resources":[{"name":"test1","type":"AWS::EC2::Instance","properties":{"region":"us-east-1","instanceType":"m5.large","operatingSystem":"linux","reservationType":"ondemand","quantity":1}}]}`

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

	if resource.Type != Ec2ResourceType {
		t.Errorf("expected %s, got %s", Ec2ResourceType, resource.Type)
	}

	ec2Properties, ok := resource.Properties.(Ec2InstanceProperties)
	if !ok {
		t.Errorf("failed to convert properties to Ec2InstanceProperties type")
	}

	if ec2Properties.Region != "us-east-1" {
		t.Errorf("expected us-east-1, got %s", ec2Properties.Region)
	}

	if ec2Properties.InstanceType != "m5.large" {
		t.Errorf("expected m5.large, got %s", ec2Properties.InstanceType)
	}

	if ec2Properties.OperatingSystem != "linux" {
		t.Errorf("expected linux, got %s", ec2Properties.OperatingSystem)
	}

	if ec2Properties.ReservationType != "ondemand" {
		t.Errorf("expected ondemand, got %s", ec2Properties.ReservationType)
	}
}
