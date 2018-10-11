package calc

import (
	"path/filepath"
	"runtime"
)

func dataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	types_dir := filepath.Dir(filename)
	return filepath.Join(types_dir, "..", "data")
}

func configFile(filename string) string {
	return filepath.Join(dataDir(), filename)
}

func float_equals(a float64, b float64, tolerance_fraction float64) bool {
	return (a*(1+tolerance_fraction)) >= b && (a*(1-tolerance_fraction)) <= b
}
