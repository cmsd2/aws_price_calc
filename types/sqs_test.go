package types

import "testing"
import "runtime"
import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func DataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	types_dir := filepath.Dir(filename)
	return filepath.Join(types_dir, "..", "data")
}

func ConfigFile(filename string) string {
	return filepath.Join(DataDir(), filename)
}

func TestSqsTypes(t *testing.T) {
	yaml_path := ConfigFile("sqs.yaml")

	sqs := Sqs{}

	data, err := ioutil.ReadFile(yaml_path)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(data, &sqs)
}
