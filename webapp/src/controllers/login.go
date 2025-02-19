package controllers

import "net/http"

func CarregarPaginaDeLogin(w http.ResponseWriter, r *http.Request) {
	_, erro := w.Write([]byte("Carregando a página de login"))
	if erro != nil {
		return
	}
}
