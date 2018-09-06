package api

type Response struct {
	Resources []ResponseResource `json:"resources"`
}

type ResponseResource struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	MonthlyCost float64 `json:"monthlyCost"`
}
