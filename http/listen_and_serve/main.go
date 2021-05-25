package main

import (
	"fmt"
	"net/http"
)

type myHandler int

// The myHandler type is now consider to be an http.Handler because it implements ServeHTTP
func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want in this function")

	// ResponseWriter usage
	w.Header().Set("Brandon-Key", "This is from me")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var m myHandler
	http.ListenAndServe(":8080", m)
}
