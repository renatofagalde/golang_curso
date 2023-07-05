package main

import (
	"log"
	"net/http"
)

func main() {
	//URI - Identifcador do recurso

	log.Fatal(http.ListenAndServe(":5000", nil))
}
