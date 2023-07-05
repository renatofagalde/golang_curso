package main

import (
	"log"
	"net/http"
)

func main() {
	//URI - Identifcador do recurso

	http.HandleFunc("/home", home())

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func home() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Olá mundo!"))
	}
}
