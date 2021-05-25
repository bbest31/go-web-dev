package main

import (
	"io"
	"net/http"
)

type CustomBasicMux struct{}

type CustomHandler struct{}

func (ch CustomHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Kitty cat")
}

// Routing done in one handler - NOT VERY ELEGANT!!
func (ch CustomBasicMux) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "DOG")
	case "/cat":
		io.WriteString(res, "CAT")
	}
}

// This function can be typecast as a http.HandlerFunc to be passed into http.Handle()
func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Handled!!")
}

func main() {
	var h CustomBasicMux
	var c CustomHandler

	// http ServeMux allows you to assign handlers to certain routes
	// mux := http.NewServeMux()
	// mux.Handle("/dog/", h)
	// mux.Handle("/cat/", c)
	// http.ListenAndServe(":8080", mux)

	// You can also just use the DefaultServeMux object and define handles using http if you are not having any custom config.
	// Then just pass nil as the handler
	http.Handle("/dog/", h)
	http.Handle("/cat/", c)
	http.Handle("/handler", http.HandlerFunc(handle))

	// Can also use HandleFunc with a handler function
	// This should probably only be used when the response is trivial
	http.HandleFunc("/fish/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "FISH")
	})

	http.ListenAndServe(":8080", nil)
}
