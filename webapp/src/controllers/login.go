package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarPaginaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
