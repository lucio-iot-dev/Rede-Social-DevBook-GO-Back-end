package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
    rotas := rotasUsuarios
		rotas = append(rotas, rotaLogin)
		rotas = append(rotas, rotasPublicacoes...) //os tres pontos faz com que ele da um append para cada um dos itens dentro desse slice

		for _,rota := range rotas {

			if rota.RequerAutenticacao {
				r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
				).Methods(rota.Metodo)
			} else {
				r.HandleFunc(rota.URI, middlewares.Logger (rota.Funcao)).Methods(rota.Metodo) //HandFunc pede 3 parametros de configuração
			}
		}
		return r
}