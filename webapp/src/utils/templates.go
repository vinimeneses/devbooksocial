package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	if erro := templates.ExecuteTemplate(w, template, dados); erro != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
	}
}
