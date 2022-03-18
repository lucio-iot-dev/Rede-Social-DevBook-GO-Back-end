package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
		if erro != nil {
		  respostas.Erro(w, http.StatusUnprocessableEntity, erro)
			return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
  }

  respostas.JSON(w, http.StatusCreated, usuario)
 
}
// BuscarUsuario busca um usuário salvo no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
   nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	 db, erro := banco.Conectar()
	 if erro != nil {
		 respostas.Erro(w, http.StatusInternalServerError, erro)
		 return
	 }
	 defer db.Close()

	 repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	 usuarios, erro := repositorio.Buscar(nomeOuNick)
	 if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}
// BuscarUsuario busca um usuário salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
     respostas.Erro(w, http.StatusBadRequest, erro)
		 return	
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
		
}
// AtualizarUsuario altera as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
   parametros := mux.Vars(r)
	 usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	 if erro != nil {
		 respostas.Erro(w, http.StatusBadRequest, erro)
		 return
	 }
	 corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	 respostas.Erro(w, http.StatusUnprocessableEntity, erro)
	 return
	 
   	 var usuario modelos.Usuario
	 if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		 respostas.Erro(w, http.StatusBadRequest, erro)
		 return
	 }

	 if erro = usuario.Preparar("edicao"); erro != nil {
		 respostas.Erro(w, http.StatusBadRequest, erro)
		 return
	 }

	 db, erro := banco.Conectar()
	 if erro != nil {
		 respostas.Erro(w, http.StatusInternalServerError, erro)
		 return
	 }
	 defer db.Close()

	 repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	 if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		 respostas.Erro(w, http.StatusInternalServerError, erro)
		 return
	 }

	 respostas.JSON(w, http.StatusNoContent, nil)

}
// DeletarUsuario exclui as informações de um usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
  parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuariosID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)


}