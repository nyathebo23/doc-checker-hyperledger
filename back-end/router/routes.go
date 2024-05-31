package router

import (
	"feuli/backend/models"
	HomeHandler "feuli/backend/routes/home"
	StatusHandler "feuli/backend/routes/status"
)

func GetRoutes() models.Routes {
	return models.Routes{
		models.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: HomeHandler.Index},
		models.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: StatusHandler.Index},
	}
}
