package handler

import (
	"crud/banco"
	"crud/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// criar usuário
func CriarUsuario(writer http.ResponseWriter, request *http.Request) {
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
	statement, erro := db.Prepare("insert into usuarios (nome,email) values (?,?);")
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao preparar insert: %s", erro)))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Erro ao executar insert"))
		return
	}

	id, erro := insercao.LastInsertId()
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Erro ao inserir usuario"))
		return
	}

	//status code
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("Usuário inserido. ID = %d", id)))

}
