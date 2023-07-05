package usuario

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {

	templates := template.Must(template.ParseGlob("home.html"))

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		usuario := Usuario{"teste", "teste@teste.com.br"}
		templates.ExecuteTemplate(writer, "home.html", usuario)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
