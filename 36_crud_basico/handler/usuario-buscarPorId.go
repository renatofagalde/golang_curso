package handler

import (
	"crud/banco"
	"crud/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// buscar ID por usuário
func BuscarUsuario(writer http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao recuperar ID(%s). Erro: %s", parametros["id"], erro)))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao conectar com o BD. Erro: %s", erro)))
		return
	}
	defer db.Close()

	cursor, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao preparar consulta de usuários: %s", erro)))
		return
	}
	defer cursor.Close()

	var usuario model.Usuario
	if cursor.Next() {
		if erro := cursor.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			writer.Write([]byte("Erro ao escanear usuario"))
			return
		}
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	//convertendo usuários em usuario
	if erro := json.NewEncoder(writer).Encode(usuario); erro != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Erro ao converter usuários para json. Erro: %s", erro)))
		return
	}

}
