package api

type Response struct {
	Resources []ResponseResource `yaml:"resources"`
}

type ResponseResource struct {
	MonthlyCost float64 `yaml:"monthlyCost"`
}