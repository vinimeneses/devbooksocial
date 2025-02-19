package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeLogin,
		RequerAutenticacao: false,
	},
	//{
	//	URI:                "/",
	//	Metodo:             http.MethodGet,
	//	Funcao:             controllers.CarregarPaginaDeLogin,
	//	RequerAutenticacao: false,
	//},
	//{
	//	URI:                "/",
	//	Metodo:             http.MethodGet,
	//	Funcao:             controllers.CarregarPaginaDeLogin,
	//	RequerAutenticacao: false,
	//},
}
