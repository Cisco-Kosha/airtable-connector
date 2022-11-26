package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	// specification routes
	a.Router.HandleFunc("/api/v0/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	//	a.Router.HandleFunc("/api/v0/specification/test", a.testConnectorSpecification).Methods("GET", "OPTIONS")

	// Records routes
	a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableId}", a.listRecords).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableId}/{recordId}", a.getRecords).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableIdOrName}/{recordId}", a.listRecords).Methods("UPDATE", "OPTIONS")
	// a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableIdOrName}", a.getRecords).Methods("POST", "OPTIONS")
	// a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableIdOrName}/{recordId}", a.listRecords).Methods("DELETE", "OPTIONS")

	a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableIdOrName}/{recordId}/comments", a.listComments).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("https://api.airtable.com/v0/v0/meta/bases", a.getBases).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
