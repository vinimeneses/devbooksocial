package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarPaginaDeLogin(w http.ResponseWriter, _ *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, _ *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, _ *http.Request) {
	utils.ExecutarTemplate(w, "home.html", nil)
}
