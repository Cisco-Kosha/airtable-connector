package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV2 = "/api/v2"

	// specification routes
	a.Router.HandleFunc(apiV2+"/", a.listConnectorSpecification).Methods("GET", "OPTIONS")

	// Records routes
	a.Router.HandleFunc("https://api.airtable.com/v0/{baseId}/{tableId}", a.getrecords).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
