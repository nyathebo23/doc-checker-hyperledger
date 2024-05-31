package document_saves

import (
	"encoding/json"
	"net/http"

	DocumentSaveModel "backend/models/v1/document_saves"
	"backend/hyperledger"
)

func Index(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		documentSave, err := DocumentSaveModel.Index(clients)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(documentSave)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
	}
}
