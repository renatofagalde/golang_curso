package main

import (
	"handlerhtml/usuario"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {

	templates := template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		usuario := usuario.Usuario{Nome: "teste", Email: "teste@teste.com.br"}
		templates.ExecuteTemplate(writer, "home.html", usuario)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
