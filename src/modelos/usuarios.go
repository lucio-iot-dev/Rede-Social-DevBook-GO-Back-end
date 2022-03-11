package modelos

import "time"
// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id,omitempty"`//Quando for passar o usuario para um json e o campo id tiver em branco ele não vai passar isso para o json ou seja ele vai tirar o campo id do json
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}