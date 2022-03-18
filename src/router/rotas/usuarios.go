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
	RequerAutenticacao: false,
},

{
	URI: "/usuarios/{usuarioID}",
	Metodo: http.MethodGet,                  //ok
	Funcao: controllers.BuscarUsuario,
	RequerAutenticacao: false,
},

{
	URI: "/usuarios/{usuarioId}",
	Metodo: http.MethodPut,
	Funcao: controllers.AtualizarUsuario,
	RequerAutenticacao: false,
},

{
	URI: "/usuarios/{usuariosID}",
	Metodo: http.MethodDelete,              //ok
	Funcao: controllers.DeletarUsuario,
	RequerAutenticacao: false,
},

}