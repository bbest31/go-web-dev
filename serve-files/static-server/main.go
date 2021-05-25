package main

import (
	"net/http"
)

/*
Serves a static webpage on the home root
*/
func main() {
	// If there is an index.html file in the Dir folder then that is automatically served.
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
