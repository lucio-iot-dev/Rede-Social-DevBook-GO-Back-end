package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacpes representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

//NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
   statement, erro := repositorio.db.Prepare(
		 "insert into publicacoes(titulo, conteudo, autor_id) values (?, ?, ?)",
		)
		if erro != nil {
			return 0, erro
		}
		defer statement.Close()

		resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
		if erro != nil {
			return 0, erro
		}

		ultimoIDInserido, erro := resultado.LastInsertId()
		if erro != nil {
			return 0, erro
		}
		return uint64(ultimoIDInserido), nil
}
