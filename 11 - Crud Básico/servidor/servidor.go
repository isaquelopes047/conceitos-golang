package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID       uint32 `json:"id"`
	Nivel    string `json:"nivel"`
	Mensagem string `json:"mensagem"`
	Mostrar  bool   `json:"mostrar"`
}

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	statement, erro := db.Prepare("INSERT INTO AvisosGeral (nivel, mensagem, mostrar) value (?, ?, ?)")
	if erro != nil {
		w.Write([]byte("Erro criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nivel, usuario.Mensagem, usuario.Mostrar)
	if erro != nil {
		w.Write([]byte("Erro criar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))
}
