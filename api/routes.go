package api

import (
	"beer/web"
	"io/fs"
	"log"
	"net/http"
)

var routes = map[string]func(http.ResponseWriter, *http.Request){
	"GET /sse":         GetSSEHandler,
	"GET /api/v1/apps": GetAppsHandler,
}

func GetRouter(headless bool) *http.ServeMux {
	router := http.NewServeMux()

	for route, handler := range routes {
		router.HandleFunc(route, handler)
	}

	if headless {
		log.Println("Running in headless mode")
	} else {
		rootWebDir, _ := fs.Sub(&web.Build, "build")
		fileServer := http.StripPrefix("/", http.FileServer(http.FS(rootWebDir)))
		router.Handle("GET /", fileServer)
	}

	return router
}
