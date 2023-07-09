package handler

import (
	"crud/banco"
	"crud/model"
	"encoding/json"
	"fmt"
	"net/http"
)

// listar usuários
func BuscarUsuarios(writer http.ResponseWriter, request *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao conectar com o BD. Erro: %s", erro)))
		return
	}
	defer db.Close()

	cursor, erro := db.Query("select * from usuarios")
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao preparar consulta de usuários: %s", erro)))
		return
	}
	defer cursor.Close()

	var usuarios []model.Usuario
	for cursor.Next() {
		var usuario model.Usuario
		if erro := cursor.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			writer.Write([]byte("Erro ao escanear usuario"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	//convertendo usuários em usuario
	if erro := json.NewEncoder(writer).Encode(usuarios); erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao converter usuários para json. Erro: %s", erro)))
		return
	}

}
