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

// BuscarPorID traz uma unica publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
  linha, erro := repositorio.db.Query(`
	select p.*, u.nick from
	publicacoes p inner join usuarios u
	on u.id = p.autor_id where p.id = ?`,
	publicacaoID,
)
if erro != nil {
	return modelos.Publicacao{}, erro
}
defer linha.Close()
var publicacao modelos.Publicacao

if linha.Next() {
  if erro = linha.Scan(
		&publicacao.ID,
		&publicacao.Titulo,
		&publicacao.Conteudo,
		&publicacao.AutorID,
		&publicacao.Curtidas,
		&publicacao.CriadoEm,
		&publicacao.AutorNick,
	); erro != nil {
		return modelos.Publicacao{}, erro
	}
}

return publicacao, nil
}
