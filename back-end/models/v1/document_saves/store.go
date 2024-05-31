package document_saves

import (
	"time"
	"encoding/json"
	"github.com/satori/go.uuid"

	"feuli/backend/models"
	"feuli/backend/hyperledger"
)

func Store(clients *hyperledger.Clients, orgId string, beneficiaryId string, filename string, hash string) (documentSave *models.DocumentSave, err error) {
	
	documentSave, err = models.NewDocumentSave(orgId, beneficiaryId, filename, hash)
	if err != nil {
		return
	}

	packet, err := json.Marshal(documentSave)
	if err != nil {
		return
	}

	if _, err = clients.Invoke("enspy", "document_save", "store", [][]byte{
		packet,
	}); err != nil {
		return
	}

	return
}
