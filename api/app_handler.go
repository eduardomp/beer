package api

import (
	"beer/core/model"
	"beer/core/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetAppsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAppsHandler - Getting apps information...")
	apps := services.GetApps()
	response, _ := json.Marshal(apps)
	w.Write(response)
}

func GetSSEHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	hostInfoChan := make(chan model.Host)
	//appsChan := make(chan model.Apps)

	go services.GetHostInfoToChannel(hostInfoChan)

	for {
		select {
		case host := <-hostInfoChan:
			hostInfo, _ := json.Marshal(host)
			fmt.Fprintf(w, "event: hostInfo\n")
			fmt.Fprintf(w, "data: %s\n\n", hostInfo)
			w.(http.Flusher).Flush()
		}
	}

}
