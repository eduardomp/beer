package api

import (
	"log"
	"net/http"
)

func RunServer(headless bool) {

	//stack := CreateStack()

	server := http.Server{
		Addr:    ":9999",
		Handler: GetRouter(headless),
	}

	log.Println("Listening on http://localhost:9999/ ...")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
