package app

import (
	"net/http"

	"github.com/kosha/airtable-connector/pkg/httpclient"
)

// getBases godoc
// @Summary Get airtable
// @Description Get Bases
// @Tags Bases
// @Accept  json
// @Produce  json
// @Success 200 {object} models.BasesStruct
// @Router /api/v0/meta/bases [get]
func (a *App) getBases(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	generate_url := a.Cfg.GetAirtableURL() + "/meta/bases"
	bases := httpclient.GetBases(generate_url, a.Cfg.GetPersonalAccessToken())
	respondWithJSON(w, http.StatusOK, bases)
}
