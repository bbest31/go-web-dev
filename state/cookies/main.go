package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Cookies are files written to the user browser to save state information.
These are passed with requests to a desired domain or url path and can hold
user session information or whatever other KV pair info you want.

This can be used instead of having a user uuid in the url query because
that would not be encrypted over http.

See Cookie object docs: https://golang.org/pkg/net/http/#Cookie
*/

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandledFunc("/multiple", multiple)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my=cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "go to dev tools /application/cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE #1:", c)

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE #2:", c2)

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
}

func multiple(w http.ResponseWriter, eq *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(w, "Cookies written -check your browser\n")
}
