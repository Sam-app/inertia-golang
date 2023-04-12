package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/inertia-golang/static"
	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/mix"
)

var (
	url          = "http://localhost:8080"
	rootTemplate = "resources/views/app.gohtml"
	// version      = ""
)

func main() {
	mixManager := mix.New("", "./static", "")

	var version string
	var err error

	version, err = mixManager.Hash("")
	if err != nil {
		log.Fatal(err)
	}

	inertiaManager := inertia.New(url, rootTemplate, version)

	mux := http.NewServeMux()
	mux.Handle("/", inertiaManager.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inertiaManager.Share("title", "Inertia Golang")
		inertiaManager.ShareFunc("asset", func(path string) (string, error) {
			return url + "/" + path, nil
		})
		err := inertiaManager.Render(w, r, "home/Index", map[string]interface{}{
			"title": "Inertia Golang",
		})
		if err != nil {
			fmt.Println(err, "inertia render error")
		}

	})))
	var fileServer http.Handler

	staticFS := http.FS(static.Files)
	fileServer = http.FileServer(staticFS)

	mux.Handle("/js/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
