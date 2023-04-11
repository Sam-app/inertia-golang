package main

import (
	"log"
	"net/http"
	"text/template"

	"example.com/inertia-golang/models"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := models.LoadPage(title)
	renderTemplate(w, "templates/view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := models.LoadPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	renderTemplate(w, "templates/edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
