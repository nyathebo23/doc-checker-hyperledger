package document_saves

import (
	"encoding/json"
	"net/http"

	"feuli/backend/models"

	DocumentSaveModel "feuli/backend/models/v1/document_saves"
	"feuli/backend/hyperledger"
)

func Store(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var documentSave models.DocumentSave
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&documentSave)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newdocumentSave, err := documentSavesModel.Store(
			clients, 
			documentSave.OrganizationId,
			documentSave.BeneficiaryId,
			documentSave.Filename,
			documentSave.Hash,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(newdocumentSave)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
