package beneficiaries

import (
	"encoding/json"
	"net/http"

	"feuli/backend/models"

	BeneficiariesModel "feuli/backend/models/v1/beneficiaries"
)

func Store(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var beneficiary models.Beneficiary
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&beneficiary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newBeneficiary, err := BeneficiariesModel.Store(clients, beneficiary.FirstName, beneficiary.LastName, beneficiary.Email, beneficiary.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(newBeneficiary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
