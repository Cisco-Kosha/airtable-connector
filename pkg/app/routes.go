package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV2 = "/api/v2"

	// specification routes
	a.Router.HandleFunc(apiV2+"/", a.listConnectorSpecification).Methods("GET", "OPTIONS")

	// Records routes
	a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableId}", a.listRecords).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableId}/{recordId}", a.getRecords).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableIdOrName}/{recordId}", a.listRecords).Methods("UPDATE", "OPTIONS")
	// a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableIdOrName}", a.getRecords).Methods("POST", "OPTIONS")
	// a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableIdOrName}/{recordId}", a.listRecords).Methods("DELETE", "OPTIONS")


	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
