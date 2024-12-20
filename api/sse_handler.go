package api

import (
	"beer/core/model"
	"beer/core/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetSSEHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	hostInfoChan := make(chan model.Host)
	connectionsChan := make(chan model.Connections)
	//appsChan := make(chan model.Apps)

	go services.GetHostInfoToChannel(hostInfoChan)
	go services.GetConnectionsToChannel(connectionsChan)

	for {
		select {
		case host := <-hostInfoChan:
			hostInfo, _ := json.Marshal(host)
			fmt.Fprintf(w, "event: hostInfo\n")
			fmt.Fprintf(w, "data: %s\n\n", hostInfo)
			w.(http.Flusher).Flush()
		case connections := <-connectionsChan:
			connectionsInfo, _ := json.Marshal(connections)
			fmt.Fprintf(w, "event: connectionsInfo\n")
			fmt.Fprintf(w, "data: %s\n\n", connectionsInfo)
			w.(http.Flusher).Flush()
		}
	}

}
