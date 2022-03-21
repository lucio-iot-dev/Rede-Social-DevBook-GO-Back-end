package rotas

import (
	"api/src/controllers"
	"net/http"
)

//slice de rotas
var rotasUsuarios = []Rota {  
	{
	URI: "/usuarios",
	Metodo: http.MethodPost,                 //ok
	Funcao: controllers.CriarUsuario,
	RequerAutenticacao: false,
},

{
	URI: "/usuarios",
	Metodo: http.MethodGet,
	Funcao: controllers.BuscarUsuarios,
	RequerAutenticacao: true,
},

{
	URI: "/usuarios/{usuarioID}",
	Metodo: http.MethodGet,                  //ok
	Funcao: controllers.BuscarUsuario,
	RequerAutenticacao: true,
},

{
	URI: "/usuarios/{usuarioId}",
	Metodo: http.MethodPut,
	Funcao: controllers.AtualizarUsuario,
	RequerAutenticacao: true,
},

{
	URI: "/usuarios/{usuariosID}",
	Metodo: http.MethodDelete,              //ok
	Funcao: controllers.DeletarUsuario,
	RequerAutenticacao: true,
},
{
	URI: "/usuarios/{usuarioId}/seguir",
	Metodo: http.MethodPost,
	Funcao: controllers.SeguirUsuario,
	RequerAutenticacao: true,
},
{
	URI: "/usuarios/{usuarioId}/parar-de-seguir",
	Metodo: http.MethodPost,
	Funcao: controllers.PararDeSeguirUsuario,
	RequerAutenticacao: true,
},

}