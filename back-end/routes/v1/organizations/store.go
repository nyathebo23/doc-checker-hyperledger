package organizations

import (
	"encoding/json"
	"net/http"

	"feuli/backend/models"

	OrganizationModel "feuli/backend/models/v1/organizations"
)

func Store(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var organization models.Organization
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&organization)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newOrg, err := OrganizationModel.Store(
			clients,
			organization.Title,
			organization.Name,
			organization.IsActive,
			organization.OrgType
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(newOrg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
