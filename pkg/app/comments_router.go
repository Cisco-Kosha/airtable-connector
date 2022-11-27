package app

import (
	"net/http"

	"github.com/kosha/airtable-connector/pkg/httpclient"
)

// getComments godoc
// @Summary Get airtable
// @Description Get Comments
// @Tags agents
// @Accept  json
// @Produce  json
// @Param baseId path string true "Base Id"
// @Param tableIdOrName path string true "Table Id"
// @Param recordId path string true "Base Id"
// @Success 200 {object} models.CommentStruct
// @Router /api/v0/{baseId}/{tableIdOrName}/{recordId}/comments [get]
func (a *App) listComments(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	agents := httpclient.ListComments(a.Cfg.GetAirtableURL(), a.Cfg.GetPersonalAccessToken())

	respondWithJSON(w, http.StatusOK, agents)
}
