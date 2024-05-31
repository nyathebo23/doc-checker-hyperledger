package beneficiaries

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"feuli/backend/models"

	BeneficiariesModel "feuli/backend/models/v1/beneficiaries"
)

func Update(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var opts BeneficiariesModel.UpdateOpts
		var beneficiary models.Beneficiary
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&beneficiary); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			opts.Replace = true
		}

		updatedBeneficiary, err := BeneficiariesModel.Update(clients, id, &beneficiary, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(updatedBeneficiary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
