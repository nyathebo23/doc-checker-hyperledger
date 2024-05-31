package organizations

import (
	"net/http"

	"github.com/gorilla/mux"

	OrganizationModel "feuli/backend/models/v1/organizations"
)

func Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		if err := OrganizationModel.Destroy(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
	}
}
