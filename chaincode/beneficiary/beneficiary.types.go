package main

// Resource resource
type Beneficiary struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// ResourceTransactionItem the transaction item
type ResourceTransactionItem struct {
	TXID        string      `json:"tx_id"`
	Beneficiary Beneficiary `json:"beneficiary"`
	Timestamp   int64       `json:"timestamp"`
}
