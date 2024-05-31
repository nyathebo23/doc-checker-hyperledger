package v1

import (
	"net/http"

	"feuli/backend/models"
	BeneficiariesHandler "feuli/backend/routes/v1/beneficiaries"
	OrganizationsHandler "feuli/backend/routes/v1/organizations"
	DocumentSavesHandler "feuli/backend/routes/v1/document_saves"
	"feuli/backend/hyperledger"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}

func GetRoutes(clients *hyperledger.Clients) map[string]models.SubRoutePackage {
	return map[string]models.SubRoutePackage{
		"/v1": {
			Middleware: Middleware(),
			Routes: models.Routes{
				// Users
				models.Route{Name: "BeneficiariesIndex", Method: "GET", Pattern: "/beneficiaries", HandlerFunc: BeneficiariesHandler.Index()},
				models.Route{Name: "BeneficiariesStore", Method: "POST", Pattern: "/beneficiaries", HandlerFunc: BeneficiariesHandler.Store()},
				models.Route{Name: "BeneficiariesReplace", Method: "PUT", Pattern: "/beneficiaries/{id}", HandlerFunc: BeneficiariesHandler.Update()},
				models.Route{Name: "BeneficiariesUpdate", Method: "PATCH", Pattern: "/beneficiaries/{id}", HandlerFunc: BeneficiariesHandler.Update()},
				models.Route{Name: "BeneficiariesDestroy", Method: "DELETE", Pattern: "/beneficiaries/{id}", HandlerFunc: BeneficiariesHandler.Destroy()},
				// Organizations
				models.Route{Name: "OrganizationsIndex", Method: "GET", Pattern: "/organizations", HandlerFunc: OrganizationsHandler.Index()},
				models.Route{Name: "OrganizationsStore", Method: "POST", Pattern: "/organizations", HandlerFunc: OrganizationsHandler.Store()},
				models.Route{Name: "OrganizationsReplace", Method: "PUT", Pattern: "/organizations/{id}", HandlerFunc: OrganizationsHandler.Update()},
				models.Route{Name: "OrganizationsUpdate", Method: "PATCH", Pattern: "/organizations/{id}", HandlerFunc: OrganizationsHandler.Update()},
				models.Route{Name: "OrganizationsDestroy", Method: "DELETE", Pattern: "/organizations/{id}", HandlerFunc: OrganizationsHandler.Destroy()},
				// DocumentSaves
				models.Route{Name: "DocumentSavesIndex", Method: "GET", Pattern: "/documentsaves", HandlerFunc: DocumentSavesHandler.Index(clients)},
				models.Route{Name: "DocumentSavesStore", Method: "POST", Pattern: "/documentsaves", HandlerFunc: DocumentSavesHandler.Store(clients)},
				models.Route{Name: "DocumentSavesReplace", Method: "PUT", Pattern: "/documentsaves/{id}", HandlerFunc: DocumentSavesHandler.Update(clients)},
				models.Route{Name: "DocumentSavesUpdate", Method: "PATCH", Pattern: "/documentsaves/{id}", HandlerFunc: DocumentSavesHandler.Update(clients)},
				models.Route{Name: "DocumentSavesDestroy", Method: "DELETE", Pattern: "/documentsaves/{id}", HandlerFunc: DocumentSavesHandler.Destroy(clients)},
				models.Route{Name: "DocumentSavesShow", Method: "GET", Pattern: "/documentsaves/{id}", HandlerFunc: DocumentSavesHandler.Show(clients)},
			},
		},
	}
}
