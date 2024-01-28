package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Repositorio representa um repositorio de publicacoes
type RepositorioPublicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *RepositorioPublicacoes {
	return &RepositorioPublicacoes{db}
}

// Criar salva uma publicação no banco de dados
func (repo RepositorioPublicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repo.db.Prepare("insert into publicacoes(titulo, conteudo, autor_id) values(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// BuscarPublicacao retorna uma publicação
func (repo RepositorioPublicacoes) BuscarPublicacao(usuarioID uint64) (modelos.Publicacao, error) {
	linhas, erro := repo.db.Query("select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?", usuarioID)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID,
			&publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// BuscarPublicações retorna todas as publicações dos seguidores de um usuário
func (repo RepositorioPublicacoes) BuscarPublicacoes(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repo.db.Query(`
	select distinct p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id
	where u.id = ? or s.seguidor_id = ?
	`, usuarioID, usuarioID)

	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID,
			&publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar atualiza uma publicação no banco de dados
func (repo RepositorioPublicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repo.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Deletar deleta uma publicação do banco de dados
func (repo RepositorioPublicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repo.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPublicacaoPorUsuario retorna todas as publicações de um usuário
func (repo RepositorioPublicacoes) BuscarPublicacaoPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repo.db.Query(`
	select p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	where p.autor_id= ?
	`, usuarioID)

	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID,
			&publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// CurtirPublicacao adiciona uma curtida a publicação
func (repo RepositorioPublicacoes) CurtirPublicacao(publicacaoID uint64) error {
	statement, erro := repo.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// DescurtirPublicacao subtrai uma curtida da publicação
func (repo RepositorioPublicacoes) DescurtirPublicacao(publicacaoID uint64) error {
	statement, erro := repo.db.Prepare(`update publicacoes set curtidas = 
	CASE WHEN curtidas > 0 THEN curtidas - 1
	ELSE 0 END
	where id = ?`)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
