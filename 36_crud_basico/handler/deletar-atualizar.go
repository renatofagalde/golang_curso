package handler

import (
	"crud/banco"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// deletar usuario
func DeletarUsuario(writer http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao recuperar ID(%s). Erro: %s", parametros["id"], erro)))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		writer.Write([]byte("Erro ao conectar com o BD"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id=?")
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao preparar delete: %s", erro)))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Erro ao executar insert"))
		return
	}

	//status code
	writer.WriteHeader(http.StatusNoContent)

}
