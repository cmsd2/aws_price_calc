package api

import (
	"fmt"
	"encoding/json"
)

const SqsResourceType = "AWS::SQS::Queue"

type Request struct {
	Description string `json:"description"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Properties interface{} `json:"properties"`
}

type SqsProperties struct {
	MessagesPerMonth float64 `json:"messagesPerMonth"`
	MessageSizeKB float64 `json:"messageSizeKB"`
	IsFifo bool `json:"isFifo"`
	MessagesPerSecond float64 `json:"messagesPerSecond"`
}

func (e *Resource) UnmarshalJSON(js []byte) error {
	var data struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	err := json.Unmarshal(js, &data)
	if err != nil {
		return err
	}

	switch data.Type {
	case SqsResourceType:
		return e.unmarshalSqsQueueResource(js)
	default:
		return fmt.Errorf("resource %s has unsupported type %s", data.Name, data.Type)
	}
}

func (e *Resource) unmarshalSqsQueueResource(js []byte) error {
	var data struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Properties SqsProperties `json:"properties"`
	}

	err := json.Unmarshal(js, &data)
	if err != nil {
		return err
	}

	e.Name = data.Name
	e.Type = data.Type
	e.Properties = data.Properties

	return nil
}