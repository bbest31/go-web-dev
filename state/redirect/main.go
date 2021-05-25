package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method, "\n\n")
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther) // 303 redirects and change the request method to GET
	// alternative method
	http.Redirect(w, req, "/", http.StatusSeeOther)
	// See HTTP 3XX status codes for redirect differences
}
