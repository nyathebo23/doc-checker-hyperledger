package document_saves

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	DocumentSaveModel "feuli/backend/models/v1/document_saves"
	"feuli/backend/hyperledger"
)

func Show(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		documentSave, err := DocumentSaveModel.Show(clients, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(documentSave)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
