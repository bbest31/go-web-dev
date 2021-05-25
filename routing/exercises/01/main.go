package main

import (
	"io"
	"net/http"
)

func base(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "HOME")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "DOG")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Brandon")
}

func main() {

	http.HandleFunc("/", base)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
