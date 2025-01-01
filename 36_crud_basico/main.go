package main

import (
	"crud/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", handler.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", handler.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", handler.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", handler.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", handler.DeletarUsuario).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":5000", router))
}
