package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositório de usuarios
type Usuarios struct {
	 db *sql.DB
}
// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
    return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	 statement, erro := repositorio.db.Prepare(
		 "insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
		)
		if erro != nil {
			  return 0, erro
		}
		defer statement.Close()

		resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
		if erro != nil {
			return 0, erro
	}
	 
	ultimoIDInserido, erro := resultado.LastInsertId()
	 if erro != nil {
		return 0, erro
   }

  return uint64(ultimoIDInserido), nil

}