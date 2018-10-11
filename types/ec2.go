package types

type Ec2 struct {
	Regions []Ec2Region `yaml:"regions"`
}

type Ec2Region struct {
	Name      string       `yaml:"name"`
	Instances Ec2Instances `yaml:"instances"`
}

type Ec2Instances struct {
	Prices Ec2Prices `yaml:"prices"`
}

type Ec2Prices struct {
	InstanceTypes []Ec2InstanceType `yaml:"instance_types"`
}

type Ec2InstanceType struct {
	Name         string           `yaml:"name"`
	Os           string           `yaml:"os"`
	Reservations []Ec2Reservation `yaml:"reservations"`
}

type Ec2Reservation struct {
	Type   string  `yaml:"type"`
	Price  float64 `yaml:"price"`
	Per    string  `yaml:"per"`
	Bullet float64 `yaml:"bullet"`
	Term   Term    `yaml:"term"`
}

type Term struct {
	Length int32  `yaml:"length"`
	Period string `yaml:"period"`
}
