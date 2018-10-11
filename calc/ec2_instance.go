package calc

import (
	"fmt"

	"github.com/cmsd2/aws_price_calc/types"
)

type TermPrice struct {
	Bullet     float64
	Price      float64
	Per        string
	TermLength int32
	TermPeriod string
}

type Ec2InstancePriceData struct {
	Region       string
	InstanceType string
	Reservation  string
	TermLength   int32
	TermPeriod   string
	Os           string
}

func Ec2InstancePrice(config *types.Ec2, data Ec2InstancePriceData) (*TermPrice, error) {
	for i := range config.Regions {
		if data.Region == config.Regions[i].Name {
			result, err := Ec2InstancePriceForRegion(&config.Regions[i], &data)
			return result, err
		}
	}

	return nil, fmt.Errorf("could not find ec2 price data for region %s", data.Region)
}

func Ec2InstancePriceForRegion(config *types.Ec2Region, data *Ec2InstancePriceData) (*TermPrice, error) {
	instanceTypes := config.Instances.Prices.InstanceTypes
	for i := range instanceTypes {
		if instanceTypes[i].Os == data.Os && instanceTypes[i].Name == data.InstanceType {
			result, err := Ec2InstancePriceForInstanceType(&instanceTypes[i], data)
			return result, err
		}
	}

	return nil, fmt.Errorf("could not find ec2 price data for instance type %s with os %s in region %s", data.InstanceType, data.Os, data.Region)
}

func Ec2InstancePriceForInstanceType(config *types.Ec2InstanceType, data *Ec2InstancePriceData) (*TermPrice, error) {

	for i := range config.Reservations {
		if config.Reservations[i].Type == data.Reservation {
			var result TermPrice
			result.Bullet = config.Reservations[i].Bullet
			result.Per = config.Reservations[i].Per
			result.Price = config.Reservations[i].Price
			result.TermLength = config.Reservations[i].Term.Length
			result.TermPeriod = config.Reservations[i].Term.Period
			return &result, nil
		}
	}

	return nil, fmt.Errorf("could not find ec2 price data for reservation type %s, instance type %s with os %s in region %s", data.Reservation, data.InstanceType, data.Os, data.Region)
}
