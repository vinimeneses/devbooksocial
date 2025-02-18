package main

import (
	"fmt"
	"log"
	"net/http"
	"socialnetworking/src/router"
)

func main() {
	fmt.Println("Rodando o projeto!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
