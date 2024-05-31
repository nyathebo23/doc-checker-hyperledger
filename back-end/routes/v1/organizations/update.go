package organizations

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"feuli/backend/models"

	OrganizationModel "feuli/backend/models/v1/organizations"
)

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var opts OrganizationModel.UpdateOpts
		var organization models.Organization
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&organization); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			opts.Replace = true
		}

		updateOrganization, err := OrganizationModel.Update(id, &organization, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(updateOrganization)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
