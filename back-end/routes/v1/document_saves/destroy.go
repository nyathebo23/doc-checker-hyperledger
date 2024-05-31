package document_saves

import (
	"net/http"

	"github.com/gorilla/mux"

	DocumentSaveModel "feuli/backend/models/v1/document_saves"
	"feuli/backend/hyperledger"
)

func Destroy(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		if err := DocumentSaveModel.Destroy(clients, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
	}
}
