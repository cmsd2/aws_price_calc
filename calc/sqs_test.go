package calc

import "testing"

func TestSqsPrice(t *testing.T) {
	config := loadSqsTestConfigFile()

	price := SqsPriceRps(&config, 100, false, 4)
	if !float_equals(199.0, price, 0.01) {
		t.Errorf("Price was incorrect, got: %f, want: %f", price, 199.0)
	}
}

func TestSqsRawPrice(t *testing.T) {
	config := loadSqsTestConfigFile()

	price := SqsPrice(&config, 267840000, false, 4)
	if !float_equals(199.0, price, 0.01) {
		t.Errorf("Price was incorrect, got: %f, want: %f", price, 199.0)
	}
}