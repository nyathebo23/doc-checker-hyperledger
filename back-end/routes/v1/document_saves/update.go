package document_saves

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"feuli/backend/models"

	DocumentSaveModel "feuli/backend/models/v1/document_saves"
	"feuli/backend/hyperledger"
)

func Update(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var opts DocumentSaveModel.UpdateOpts
		var documentSave models.DocumentSave
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&documentSave); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			opts.Replace = true
		}

		updatedDocumentSave, err := documentSavesModel.Update(clients, id, &documentSave, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(updatedDocumentSave)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
