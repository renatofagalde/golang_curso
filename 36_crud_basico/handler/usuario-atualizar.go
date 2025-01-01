package handler

import (
	"crud/banco"
	"crud/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

// atualizar usuprio
func AtualizarUsuario(writer http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao recuperar ID(%s). Erro: %s", parametros["id"], erro)))
		return
	}

	payload, erro := io.ReadAll(request.Body)
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao conectar com o BD. Erro: %s", erro)))
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(payload, &usuario); erro != nil {
		writer.Write([]byte("Erro ao converter usuário para struct"))
	}
	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		writer.Write([]byte("Erro ao conectar com o BD"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome=?,email=? where id=?")
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao preparar insert: %s", erro)))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Erro ao executar update"))
		return
	}

	//status code
	writer.WriteHeader(http.StatusNoContent)

}
