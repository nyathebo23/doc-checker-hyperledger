package organizations

import (
	"encoding/json"
	"net/http"

	OrganizationModel "feuli/backend/models/v1/rawresourcetypes"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		organizations, err := OrganizationModel.Index()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(organizations)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
	}
}
