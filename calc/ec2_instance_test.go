package calc

import (
	"testing"

	"github.com/cmsd2/aws_price_calc/types"
)

func TestEc2InstancePrice(t *testing.T) {
	config := types.NewConfigFromFiles().Ec2

	data := Ec2InstancePriceData{
		Region:       "us-east-1",
		InstanceType: "m5.large",
		Reservation:  "ondemand",
		Os:           "linux",
	}
	termPrice, err := Ec2InstancePrice(&config, data)

	if err != nil {
		t.Errorf("Unexpected error returned, got: %s", err)
		return
	}

	if !float_equals(0.0, termPrice.Bullet, 0.01) {
		t.Errorf("Upfront price was incorrect, got: %f, want: %f", termPrice.Bullet, 0.0)
	}

	if !float_equals(0.096, termPrice.Price, 0.01) {
		t.Errorf("Periodic price was incorrect, got: %f, want: %f", termPrice.Price, 0.096)
	}
}
