package main

import (
	"log"
	"net/http"
	"text/template"

	"example.com/inertia-golang/models"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := models.LoadPage(title)
	t, _ := template.ParseFiles("templates/view.html")
	t.Execute(w, p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := models.LoadPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	t, _ := template.ParseFiles("templates/edit.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
