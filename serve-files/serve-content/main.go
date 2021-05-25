package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/bork", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/samoyed.png">`)

}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("samoyed.png")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer f.Close()

	// Stat gives us back file information
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	// This line serves the file to the writer
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
