package beneficiaries

import (
	"encoding/json"
	"net/http"

	BeneficiariesModel "feuli/backend/models/v1/beneficiaries"
)

func Index(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		beneficiaries, err := BeneficiariesModel.Index(clients)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(beneficiaries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
	}
}
