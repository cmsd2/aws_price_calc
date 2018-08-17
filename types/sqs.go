package types

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Sqs struct {
	Classes  []string    `yaml:"classes"`
	FreeTier SqsFreeTier `yaml:"free_tier"`
	Limits   Limits      `yaml:"limits"`
	Price    Price       `yaml:"price"`
}

type SqsFreeTier struct {
	Requests RequestsPerPeriod `yaml:"requests"`
}

type RequestsPerPeriod struct {
	Value  float64 `yaml:"value"`
	Period string  `yaml:"period"`
}

type Limits struct {
	RequestPayload RequestPayloadLimits `yaml:"request_payload"`
}

type RequestPayloadLimits struct {
	Value float64
	Units string
}

type Price struct {
	Requests RequestsPrices `yaml:"requests"`
	Data     DataPrices     `yaml:"data"`
}

type RequestsPrices struct {
	Per      float64   `yaml:"per"`
	Standard float64 `yaml:"standard"`
	Fifo     float64 `yaml:"fifo"`
}

type RequestsPrice struct {
	Price float64 `yaml:"price"`
	Class string  `yaml:"class"`
}

type DataPrices struct {
	Period string          `yaml:"period"`
	In     DataFlatPrice   `yaml:"in"`
	Out    DataLadderPrice `yaml:"out"`
}

type DataFlatPrice struct {
	Type  string  `yaml:"type"`
	Price float64 `yaml:"price"`
}

type DataLadderPrice struct {
	Type  string     `yaml:"type"`
	Units string     `yaml:"units"`
	Bands []DataBand `yaml:"bands"`
}

type DataBand struct {
	From  float64 `yaml:"from"`
	To    float64 `yaml:"to"`
	Price float64 `yaml:"price"`
	Poa   bool    `yaml:"poa"`
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