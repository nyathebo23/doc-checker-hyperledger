package models

import "time"

type DocumentSaves []DocumentSave

type DocumentSave struct {
	ID             string `json:"id"`
	OrganizationId string `json:"organization_id"`
	BeneficiaryId  string `json:"beneficiary_id"`
	Filename       string `json:"filename"`
	Hash           string `json:"hash"`
	SaveDate       *time.Time  `json:"save_date"`
}


func NewDocumentSave(organizationId string, beneficiaryId string, filename string, hash string) (documentSave *DocumentSave, err error) {
	documentSave = new(DocumentSave)

	if documentSave.ID, err = genUUID(); err != nil {
		return
	}

	documentSave.FileName = name
	documentSave.OrganizationId = organizationId
	documentSave.BeneficiaryId = beneficiaryId
	documentSave.hash = hash

	t := time.Now()
	documentSave.SaveDate = &t

	return
}
