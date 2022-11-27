package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	// specification routes
	a.Router.HandleFunc("/api/v0/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	// Records routes
	a.Router.HandleFunc("/api/v0/{baseId}/{tableId}", a.listRecords).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/api/v0/{baseId}/{tableId}", a.createRecords).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/api/v0/{baseId}/{tableId}/{recordId}", a.getRecords).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc("/api/v0/{baseId}/{tableId}/{recordId}", a.updateRecords).Methods("UPDATE", "OPTIONS")
	// a.Router.HandleFunc("/api/v0/{baseId}/{tableId}/{recordId}", a.deleteRecords).Methods("DELETE", "OPTIONS")

	// Comments routes
	a.Router.HandleFunc("/api/v0/{baseId}/{tableId}/{recordId}/comments", a.listComments).Methods("GET", "OPTIONS")
	// Bases routes
	a.Router.HandleFunc("/api/v0/meta/bases", a.getBases).Methods("GET", "OPTIONS")
	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
