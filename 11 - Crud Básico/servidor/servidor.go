package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// BuscarUsuarios vai buscar apenas um registros do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM AvisosGeral")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os dados"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nivel, &usuario.Mensagem, &usuario.Mostrar); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader((http.StatusOK))
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuarios para JSON"))
	}
}

// BuscarUsuarios vai buscar os registros do banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	linha, erro := db.Query("SELECT * FROM AvisosGeral WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o registro"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nivel, &usuario.Mensagem, &usuario.Mostrar); erro != nil {
			w.Write([]byte("Erro ao buscar o usuario"))
			return
		}
	}

	w.WriteHeader((http.StatusOK))
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuarios para JSON"))
	}
}

// AtualizarUsuario vai atualizar o registro do banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parametro para inteiro"))
		return
	}

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
	defer db.Close()

	statement, erro := db.Prepare("UPDATE AvisosGeral SET nivel = ?, mensagem = ?, mostrar = ? WHERE id = ?")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao preparar statement: " + erro.Error()))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nivel, usuario.Mensagem, usuario.Mostrar, ID); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao atualizar o usuario: " + erro.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario vai Deleta o registro do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM AvisosGeral WHERE id = ?")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao preparar statement: " + erro.Error()))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuario: " + erro.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
