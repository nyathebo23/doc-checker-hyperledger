package document_saves

import (
	"encoding/json"
	"errors"

	"feuli/backend/hyperledger"
	"feuli/backend/models"
)

func Show(clients *hyperledger.Clients, id string) (documentSave *models.DocumentSave, err error) {

	documentSaves := new(models.DocumentSaves)

	res, err := clients.Query("enspy", "mainchannel", "document_save", "queryString", [][]byte{
		[]byte("{\"selector\":{ \"id\": { \"$eq\":\"" + id + "\" } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, documentSaves); err != nil {
		return
	}

	list := *documentSaves

	if len(list) == 0 {
		err = errors.New("unable to find document save with id " + id)
		return
	}

	documentSave = &list[0]

	return
}
