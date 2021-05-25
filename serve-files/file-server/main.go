package main

import (
	"io"
	"net/http"
)

func main() {
	// Strip Prefix strips the /resources url prefix so the handler accessing the URL just gets what comes after
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/samoyed.png">`)
}
