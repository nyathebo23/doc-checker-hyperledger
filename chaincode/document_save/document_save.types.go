package main

// ResourceType resource
type DocumentSave struct {
	ID             string `json:"id"`
	OrganizationId string `json:"organization_id"`
	BeneficiaryId  string `json:"beneficiary_id"`
	Filename       string `json:"filename"`
	Hash           string `json:"hash"`
	SaveDate       int64  `json:"save_date"`
}

// ResourceTypeTransactionItem transaction item
type DocumentSaveTransactionItem struct {
	TXID         string            `json:"tx_id"`
	DocumentSave DocumentSave  `json:"document_save"`
	Timestamp    int64             `json:"timestamp"`
}
