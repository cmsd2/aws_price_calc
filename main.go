package main

import (
	"errors"
	"log"

	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cmsd2/aws_price_calc/api"
	"github.com/cmsd2/aws_price_calc/calc"
	"github.com/cmsd2/aws_price_calc/types"
	"github.com/davyzhang/agw"
	"github.com/gorilla/mux"
)

//go:generate go run scripts/includetxt.go

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")

	config *types.Types
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil

}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request api.Request

	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.(*agw.LPResponse).WriteBody(map[string]string{
			"error": "Bad Request",
		}, false)
		return
	}

	response, err := calcResources(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.(*agw.LPResponse).WriteBody(map[string]string{
			"error":   "Bad Request",
			"message": err.Error(),
		}, false)
	}

	w.(*agw.LPResponse).WriteBody(response, false)
}

func calcResources(request api.Request) (*api.Response, error) {
	response := new(api.Response)

	for i := range request.Resources {
		resource := request.Resources[i]

		responseResource, err := calcResourceCost(resource)
		if err != nil {
			return nil, err
		}

		response.Resources = append(response.Resources, *responseResource)
	}

	return response, nil
}

func calcResourceCost(resource api.Resource) (*api.ResponseResource, error) {
	responseResource := new(api.ResponseResource)
	responseResource.Name = resource.Name
	responseResource.Type = resource.Type

	switch resource.Type {
	case api.SqsResourceType:
		sqsProperties, ok := resource.Properties.(api.SqsProperties)
		if !ok {
			return nil, fmt.Errorf("unexpected deserialised type")
		}

		err := calcSqsQueueCost(resource, sqsProperties, responseResource)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("resource %s has unsupported type %s", resource.Name, resource.Type)
	}
	return responseResource, nil
}

func calcSqsQueueCost(resource api.Resource, sqsProperties api.SqsProperties, responseResource *api.ResponseResource) error {
	var price float64

	if sqsProperties.MessagesPerSecond != 0 {
		price = calc.SqsPriceRps(&config.Sqs, sqsProperties.MessagesPerSecond, sqsProperties.IsFifo, sqsProperties.MessageSizeKB)
	} else {
		price = calc.SqsPrice(&config.Sqs, sqsProperties.MessagesPerMonth, sqsProperties.IsFifo, sqsProperties.MessageSizeKB)
	}

	responseResource.MonthlyCost = price

	return nil
}

func loadConfig() *types.Types {
	yaml_path := path.Join("data", "Sqs.yaml")

	return types.LoadConfigFile(yaml_path)
}

func main() {
	config = loadConfig()
	mux := mux.NewRouter()
	mux.HandleFunc("/calc", handleCalc)
	lambda.Start(agw.Handler(mux))
}
