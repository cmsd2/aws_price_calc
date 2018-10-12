package types

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"

	"github.com/cmsd2/aws_price_calc/data"
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

func New() *Types {
	return &Types{}
}

func (config *Types) LoadConfigFile(file_name string) {
	yamlPath := path.Join("data", file_name)

	data, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(data, config)
}

func NewConfigFromFiles() *Types {
	config := New()

	yaml.Unmarshal([]byte(data.Ec2), config)
	yaml.Unmarshal([]byte(data.Sqs), config)

	return config
}
