package api

import (
	"beer/core/services"
	"encoding/json"
	"log"
	"net/http"
)

func GetAppsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAppsHandler - Getting apps information...")
	apps := services.GetApps()
	response, _ := json.Marshal(apps)
	w.Write(response)
}
