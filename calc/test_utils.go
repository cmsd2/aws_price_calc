package calc

import (
	"runtime"
	"path/filepath"
	"github.com/cmsd2/aws_price_calc/types"
)

func dataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	types_dir := filepath.Dir(filename)
	return filepath.Join(types_dir, "..", "data")
}

func configFile(filename string) string {
	return filepath.Join(dataDir(), filename)
}

func loadSqsTestConfigFile() types.Sqs {
	yaml_path := configFile("Sqs.yaml")

	return types.LoadConfigFile(yaml_path).Sqs
}

func float_equals(a float64, b float64, tolerance_fraction float64) bool {
	return (a * (1 + tolerance_fraction)) >= b && (a * (1 - tolerance_fraction)) <= b
}