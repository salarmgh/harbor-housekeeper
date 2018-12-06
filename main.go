package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/imagecleaner", imageCleanerHandler)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("Running")
	err := http.ListenAndServe("0.0.0.0:8085", loggedRouter)
	if err != nil {
		log.Println(err)
	}
}
