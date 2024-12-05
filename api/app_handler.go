package api

import (
	"beer/core/services"
	"encoding/json"
	"log"
	"net/http"
)

func GetAppsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAppsHandler - Getting apps information...")

	fetchProcesses := r.URL.Query().Get("fetchProcesses")
	apps := services.GetApps(fetchProcesses == "true")
	response, _ := json.Marshal(apps)
	w.Write(response)
}
