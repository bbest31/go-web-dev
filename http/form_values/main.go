package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type myHandler int

// The myHandler type is now consider to be an http.Handler because it implements ServeHTTP
func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// Grab the form values from the body
	postForm := r.PostForm
	for key, value := range postForm {
		fmt.Printf("Key = %v | Value = %v\n", key, value)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var m myHandler
	http.ListenAndServe(":8080", m)
}
