package api

import (
	"encoding/json"
	"fmt"
)

// list of cloudformation types here:
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
const SqsResourceType = "AWS::SQS::Queue"
const Ec2ResourceType = "AWS::EC2::Instance"

const LinuxEc2OperatingSystem = "Linux"
const OnDemandEc2ReservationType string = "OnDemand"
const ReservedEc2ReservationType string = "Reserved"
const SpotEc2ReservationType string = "Spot"

type Request struct {
	Description string     `json:"description"`
	Resources   []Resource `json:"resources"`
}

type Resource struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Properties interface{} `json:"properties"`
}

type SqsProperties struct {
	MessagesPerMonth  float64 `json:"messagesPerMonth"`
	MessageSizeKB     float64 `json:"messageSizeKB"`
	IsFifo            bool    `json:"isFifo"`
	MessagesPerSecond float64 `json:"messagesPerSecond"`
}

type Ec2InstanceProperties struct {
	InstanceType    string  `json:"instanceType"`
	OperatingSystem string  `json:"operatingSystem"`
	ReservationType string  `json:"reservationType"`
	Utilisation     float64 `json:"utilisation"`
	Region          string  `json:"region"`
	Quantity        int32   `json:"quantity"`
	TermYears       int32   `json:"term_years"`
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
	case Ec2ResourceType:
		return e.unmarshalEc2InstanceResource(js)
	default:
		return fmt.Errorf("resource %s has unsupported type %s", data.Name, data.Type)
	}
}

func (e *Resource) unmarshalSqsQueueResource(js []byte) error {
	var data struct {
		Name       string        `json:"name"`
		Type       string        `json:"type"`
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

func (e *Resource) unmarshalEc2InstanceResource(js []byte) error {
	var data struct {
		Name       string                `json:"name"`
		Type       string                `json:"type"`
		Properties Ec2InstanceProperties `json:"properties"`
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
