package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("./assets/samoyed.png")
	if err != nil {
		log.Panic(err)
		// http.Error responds to the request with a plaintext error text and a status code
		http.Error(w, "error opening file", http.StatusInternalServerError)
		return
	}

	defer f.Close()

	io.WriteString(w, "call handled")
}

/*
Serves a static webpage on the home root
*/
func main() {
	// ListenAndServe returns an error so log.Fatal catches that error and logs it to our log file.
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", handle)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
