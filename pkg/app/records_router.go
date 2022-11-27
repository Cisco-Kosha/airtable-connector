package app

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kosha/airtable-connector/pkg/httpclient"
)

// ListRecords godoc
// @Summary Get airtable
// @Description Get Records
// @Tags Records
// @Accept  json
// @Produce  json
// @Success 200 {object} models.RecordsStruct
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

// getRecord godoc
// @Summary Get airtable
// @Description Get Records
// @Tags Records
// @Accept  json
// @Produce  json
// @Param baseId path string true "Base Id"
// @Param tableIdOrName path string true "Table Id"
// @Param recordId path string true "Base Id"
// @Success 200 {object} models.Record
// @Router /api/v0/{baseId}/{tableIdOrName}/{recordId} [get]
func (a *App) getRecords(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	base_id := vars["baseId"]
	table_id := vars["tableId"]
	recordId := vars["recordId"]
	generated_url := a.Cfg.GetAirtableURL() + "/" + base_id + "/" + table_id + "/" + recordId
	fmt.Println(generated_url)
	agents := httpclient.GetRecords(generated_url, a.Cfg.GetPersonalAccessToken())
	respondWithJSON(w, http.StatusOK, agents)
}

// createRecords= godoc
// @Summary Create a record for a project
// @Description Create a record for airtable
// @Description Please refer to https://airtable.com/developers/web/api/create-records for more parameter options.
// @Tags createRecords
// @Accept json
// @Produce json
// @Param project body models.Record false "Enter project risk properties"
// @Param baseId path string true "Base Id"
// @Param tableId path string true "Table Id"
// @Success 200
// @Router /api/v0/{baseId}/{tableId} [post]
func (a *App) createRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	base_id := vars["baseId"]
	table_id := vars["tableId"]
	generated_url := a.Cfg.GetAirtableURL() + "/" + base_id + "/" + table_id
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}

	res, err := httpclient.CreateRecord(generated_url, bodyBytes, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in createRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}

// deleteRecords= godoc
// @Summary delete a record for a project
// @Description delete a record for the airtable
// @Description Please refer to https://airtable.com/developers/web/api/create-records for more parameter options.
// @Tags deleteRecords
// @Accept json
// @Produce json
// @Param project body models.Record false "Enter project risk properties"
// @Param baseId path string true "Base Id"
// @Param tableId path string true "Table Id"
// @Param RecordId path string true "Record Id"
// @Success 200
// @Router /api/v0/{baseId}/{tableId}/{recordId} [delete]
func (a *App) deleteRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	base_id := vars["baseId"]
	table_id := vars["tableId"]
	record_id := vars["recordId"]
	generated_url := a.Cfg.GetAirtableURL() + "/" + base_id + "/" + table_id + "/" + record_id
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}

	res, err := httpclient.DeleteRecord(generated_url, bodyBytes, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in deleteRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}

// updateRecords= godoc
// @Summary update a record for a project
// @Description Update a record for the airtable
// @Description Please refer to https://airtable.com/developers/web/api/create-records for more parameter options.
// @Tags updateRecords
// @Accept json
// @Produce json
// @Param project body models.Record false "Enter project risk properties"
// @Param baseId path string true "Base Id"
// @Param tableId path string true "Table Id"
// @Param RecordId path string true "Record Id"
// @Success 200
// @Router /api/v0/{baseId}/{tableId}/{recordId} [put]
func (a *App) updateRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	base_id := vars["baseId"]
	table_id := vars["tableId"]
	record_id := vars["recordId"]
	generated_url := a.Cfg.GetAirtableURL() + "/" + base_id + "/" + table_id + "/" + record_id
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}

	res, err := httpclient.UpdateRecord(generated_url, bodyBytes, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in updateRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}
