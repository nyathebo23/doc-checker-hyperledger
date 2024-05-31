package document_saves

import (
	"encoding/json"

	"feuli/backend/hyperledger"
	"feuli/backend/models"
)

func Index(clients *hyperledger.Clients, orgId string) (documentSaves *models.DocumentSaves, err error) {
	documentSaves = new(models.DocumentSaves)

	res, err := clients.Query("enspy", "mainchannel", "document_save", "queryString", [][]byte{
		[]byte("{\"selector\":{ \"organizationId\": { \"$eq\":\"" + orgId + "\" } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, documentSaves); err != nil {
		return
	}

	return
}
