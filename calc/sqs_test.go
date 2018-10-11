package calc

import (
	"testing"

	"github.com/cmsd2/aws_price_calc/types"
)

func TestSqsPriceRps(t *testing.T) {
	config := loadSqsTestConfigFile()

	price := SqsPriceRps(&config, 100, false, 4)
	if !float_equals(199.0, price, 0.01) {
		t.Errorf("Price was incorrect, got: %f, want: %f", price, 199.0)
	}
}

func TestSqsPrice(t *testing.T) {
	config := loadSqsTestConfigFile()

	price := SqsPrice(&config, 267840000, false, 4)
	if !float_equals(199.0, price, 0.01) {
		t.Errorf("Price was incorrect, got: %f, want: %f", price, 199.0)
	}
}

func loadSqsTestConfigFile() types.Sqs {
	yamlPath := configFile("Sqs.yaml")

	return types.LoadConfigFile(yamlPath).Sqs
}
