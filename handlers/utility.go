package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile, WebAppRoot+"/templates/header.html", WebAppRoot+"/templates/footer.html")
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

// RenderGatedTemplate ...
func RenderGatedTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile, WebAppRoot+"/templates/gatedheader.html", WebAppRoot+"/templates/footer.html")
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}
