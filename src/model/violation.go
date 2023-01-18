package model

type Violation struct {
	ID            string `json:"id"`
	ProductDID    string `json:"productDID"`
	TransactionID string `json:"transactionID"`
	SLA           struct {
		ID   string `json:"id"`
		Href string `json:"href"`
	} `json:"sla"`
	Rule struct {
		ID             string `json:"id"`
		Metric         string `json:"metric"`
		Unit           string `json:"unit"`
		ReferenceValue string `json:"referenceValue"`
		Operator       string `json:"operator"`
		Tolerance      string `json:"tolerance"`
		Consequence    string `json:"consequence"`
	} `json:"rule"`
	Violation struct {
		ActualValue string `json:"actualValue"`
	} `json:"violation"`
}
