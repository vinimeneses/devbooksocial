package router

import (
	"github.com/gorilla/mux"
	"socialnetworking/src/router/rotas"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
