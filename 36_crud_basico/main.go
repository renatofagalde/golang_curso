package main

import (
	"crud/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", controller.CriarUsuario).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":5000", router))
}
