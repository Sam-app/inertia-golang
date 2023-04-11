package main

import (
	"fmt"

	"example.com/inertia-golang/models"
)

func main() {
	p1 := &models.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()

	p2, _ := models.LoadPage("TestPage")
	fmt.Println(string(p2.Body))
}
