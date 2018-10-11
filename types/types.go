package types

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	yaml "gopkg.in/yaml.v2"
)

type Types struct {
	Sqs Sqs `yaml:"sqs"`
	Ec2 Ec2 `yaml:"ec2"`
}

func dataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	typesDir := filepath.Dir(filename)
	return filepath.Join(typesDir, "..", "data")
}

func configFile(filename string) string {
	return filepath.Join(dataDir(), filename)
}

func LoadConfigFile(file_path string) *Types {
	types := Types{}

	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(data, &types)

	return &types
}
