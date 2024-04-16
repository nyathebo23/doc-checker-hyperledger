package main

// ResourceType resource
type Organization struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	OrgType  string `json:"org_type"`
}

// ResourceTypeTransactionItem transaction item
type OrganizationTransactionItem struct {
	TXID         string            `json:"tx_id"`
	Organization ResourceTypeIndex `json:"organization"`
	Timestamp    int64             `json:"timestamp"`
}
