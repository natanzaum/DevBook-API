package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositorio de usuarios
type Repositorio struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Repositorio {
	return &Repositorio{db}
}

// Criar cria um usuário no banco de dados
func (u Repositorio) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare("insert into usuarios(nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// BuscarUsuarios busca usuários no banco de acordo com Nome ou Nick
func (u Repositorio) BuscarUsuarios(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick

	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarUsuarioPorID busca um usuário por ID no banco de dados
func (u Repositorio) BuscarUsuarioPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := u.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", ID)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// AtualizarUsuario edita as informações de um usuario no banco de dados
func (u Repositorio) AtualizarUsuario(ID uint64, usuario modelos.Usuario) error {
	statement, erro := u.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar edita as informações de um usuario no banco de dados
func (u Repositorio) Deletar(ID uint64) error {
	statement, erro := u.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (u Repositorio) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, erro := u.db.Query("select id, senha from usuarios where email = ?", email)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}
