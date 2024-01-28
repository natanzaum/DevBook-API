package modelos

type Senha struct {
	SenhaAtual string `json:"senha_atual"`
	NovaSenha  string `json:"nova_senha"`
}
