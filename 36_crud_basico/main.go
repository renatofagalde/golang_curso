package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
