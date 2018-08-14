package types

type Sqs struct {
	Classes  []string    `yaml:"classes"`
	FreeTier SqsFreeTier `yaml:"free_tier"`
}

type SqsFreeTier struct {
	Requests RequestsPerPeriod `yaml:"requests"`
}

type RequestsPerPeriod struct {
	Value  float32 `yaml:"value"`
	Period string  `yaml:"period"`
}

type Limits struct {
	RequestPayload RequestPayloadLimits `yaml:"request_payload"`
}

type RequestPayloadLimits struct {
	Value float32
	Units string
}

type Price struct {
	Requests RequestsPrices `yaml:"requests"`
	Data     DataPrices     `yaml:"data"`
}

type RequestsPrices struct {
	Per    int32           `yaml:"per"`
	Prices []RequestsPrice `yaml:"prices"`
}

type RequestsPrice struct {
	Price float32 `yaml:"price"`
	Class string  `yaml:"class"`
}

type DataPrices struct {
	Period string          `yaml:"period"`
	In     DataFlatPrice   `yaml:"in"`
	Out    DataLadderPrice `yaml:"out"`
}

type DataFlatPrice struct {
	Type  string  `yaml:"type"`
	Price float32 `yaml:"price"`
}

type DataLadderPrice struct {
	Type  string     `yaml:"type"`
	Units string     `yaml:"units"`
	Bands []DataBand `yaml:"bands"`
}

type DataBand struct {
	From  float32 `yaml:"from"`
	To    float32 `yaml:"to"`
	Price float32 `yaml:"price"`
}
