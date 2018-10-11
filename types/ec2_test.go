package types

import "testing"

func loadEc2TestConfigFile() Ec2 {
	yamlPath := configFile("Ec2.yaml")

	return LoadConfigFile(yamlPath).Ec2
}

func TestEc2Types(t *testing.T) {
	loadEc2TestConfigFile()
}
