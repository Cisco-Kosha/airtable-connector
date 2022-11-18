package app

import (
	"net/http"

	"github.com/kosha/freshservice-connector/pkg/httpclient"
)

// getAgents godoc
// @Summary Get airtable
// @Description Get Records
// @Tags agents
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Records
// @Router /v0/{baseId}/{tableIdOrName} [get]
func (a *App) getrecords(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	agents := httpclient.GetRecords(a.Cfg.GetAirtableURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}
