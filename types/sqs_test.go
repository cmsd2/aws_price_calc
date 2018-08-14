package types

import "testing"
import "runtime"
import (
	"path/filepath"
)

func dataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	types_dir := filepath.Dir(filename)
	return filepath.Join(types_dir, "..", "data")
}

func configFile(filename string) string {
	return filepath.Join(dataDir(), filename)
}

func loadSqsTestConfigFile() Sqs {
	yaml_path := configFile("sqs.yaml")

	return LoadConfigFile(yaml_path).Sqs
}

func TestSqsTypes(t *testing.T) {
	loadSqsTestConfigFile()
}

func TestSqsRequestsPrice(t *testing.T) {
	sqs := loadSqsTestConfigFile()

	if sqs.Price.Requests.Per != 1000000.0 {
		t.Errorf("Per was incorrect, got: %f, want: %f", sqs.Price.Requests.Per, 1000000.0)
	}

	switch {
	case sqs.Price.Requests.Standard <= 0:
		t.Errorf("Standard was incorrect, got: %f, want: %s", sqs.Price.Requests.Standard, "in (0,1)")
	case sqs.Price.Requests.Standard >= 1:
		t.Errorf("Standard was incorrect, got: %f, want: %s", sqs.Price.Requests.Standard, "in (0,1)")
	}

	switch {
	case sqs.Price.Requests.Fifo <= 0:
		t.Errorf("Fifo was incorrect, got: %f, want: %s", sqs.Price.Requests.Fifo, "in (0,1)")
	case sqs.Price.Requests.Fifo >= 1:
		t.Errorf("Fifo was incorrect, got: %f, want: %s", sqs.Price.Requests.Fifo, "in (0,1)")
	}
}