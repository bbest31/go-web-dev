package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// log.Fatal(http.ListenAndServe(":8080", nil))

	log.Fatal(http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil))
}
