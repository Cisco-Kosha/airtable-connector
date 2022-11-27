package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kosha/airtable-connector/pkg/httpclient"
)

// getRecords godoc
// @Summary Get airtable
// @Description Get Records
// @Tags agents
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Records
// @Param baseId path string true "Base Id"
// @Param tableId path string true "Table Id"
// @Router /api/v0/{baseId}/{tableId} [get]
func (a *App) listRecords(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	base_id := vars["baseId"]
	table_id := vars["tableId"]
	generated_url := a.Cfg.GetAirtableURL() + "/" + base_id + "/" + table_id
	agents := httpclient.ListRecords(generated_url, a.Cfg.GetPersonalAccessToken())
	respondWithJSON(w, http.StatusOK, agents)
}

// getRecords godoc
// @Summary Get airtable
// @Description Get Records
// @Tags agents
// @Accept  json
// @Produce  json
// @Param baseId path string true "Base Id"
// @Param tableIdOrName path string true "Table Id"
// @Param recordId path string true "Base Id"
// @Success 200 {object} models.Records
// @Router /api/v0/{baseId}/{tableIdOrName}/{recordId} [get]
func (a *App) getRecords(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	base_id := vars["baseId"]
	table_id := vars["tableIdOrName"]
	recordId := vars["recordId"]
	generated_url := a.Cfg.GetAirtableURL() + "/" + base_id + "/" + table_id + "/" + recordId
	fmt.Println(generated_url)
	agents := httpclient.GetRecords(generated_url, a.Cfg.GetPersonalAccessToken())
	respondWithJSON(w, http.StatusOK, agents)
}
