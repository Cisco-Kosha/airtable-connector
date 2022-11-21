package app

import (
	"net/http"

	"github.com/kosha/airtable-connector/pkg/httpclient"
)


// getBases godoc
// @Summary Get airtable
// @Description Get Records
// @Tags agents
// @Accept  json
// @Produce  json
// @Success 200 {object} models.BasesStruct
// @Router /v0/meta/bases [get]
func (a *App) getBases(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	agents := httpclient.GetBases(a.Cfg.GetAirtableURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}