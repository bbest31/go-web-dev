package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", urlQuery)
	http.HandleFunc("/post", bodyQuery)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func urlQuery(w http.ResponseWriter, req *http.Request) {
	// Returns the first value for the named component of the query
	// POST and PUT body params take precendence over the URL query string values.
	v := req.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}

func bodyQuery(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>
	`+v)
}

// visit this page: http://localhost:8080/?q=dog
