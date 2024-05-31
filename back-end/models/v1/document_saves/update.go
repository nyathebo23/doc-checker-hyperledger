package document_saves

import (
	"encoding/json"
	"feuli/backend/models"
	"feuli/backend/hyperledger"
)

type UpdateOpts struct {
	Replace bool
}

func Update(clients *hyperledger.Clients, id string, docSave *models.DocumentSave, opts *UpdateOpts) (*models.DocumentSave, error) {
	if !opts.Replace {
		existingDocSave, err := Show(clients, id)
		if err != nil {
			return nil, err
		}

		if docSave.Filename != "" && existingDocSave.Filename != docSave.Filename {
			existingDocSave.Filename = docSave.Filename
		}

	}

	packet, err := json.Marshal(docSave)
	if err != nil {
		return nil, err
	}

	if _, err = clients.Invoke("enspy", "documentSave", "update", [][]byte{
		[]byte(id),
		packet,
	}); err != nil {
		return nil, err
	}

	return docSave, nil
}
