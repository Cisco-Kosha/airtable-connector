package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	// specification routes
	a.Router.HandleFunc("/api/v0/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	// Records routes
	a.Router.HandleFunc("/api/v0/{baseId}/{tableId}", a.listRecords).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/api/v0/{baseId}/{tableId}/{recordId}", a.getRecords).Methods("GET", "OPTIONS")
	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
