package beneficiaries

import (
	"net/http"

	"github.com/gorilla/mux"

	BeneficiariesModel "feuli/backend/models/v1/beneficiaries"
)

func Destroy(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		if err := BeneficiariesModel.Destroy(clients, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
	}
}
