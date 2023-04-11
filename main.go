package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, "body")
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
