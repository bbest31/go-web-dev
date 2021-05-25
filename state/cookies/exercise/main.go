package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", counter)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func counter(w http.ResponseWriter, req *http.Request) {

	// Get cookie "visits"
	c, err := req.Cookie("visits")
	if err != nil {
		// If not there then init with 1
		http.SetCookie(w, &http.Cookie{
			Name:  "visits",
			Value: "1",
		})
		fmt.Fprintln(w, "User visit counter: 1")
		return
	}

	// If present then increment
	visits, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, "Unable to convert visits to int, please clear cookie cache!", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "visits", Value: strconv.Itoa(visits + 1)})
	fmt.Fprintln(w, "User visit counter: ", visits+1)

}

func expire(w http.ResponseWriter, req *http.Request) {
	// expire a cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "visits",
		Value:  "0",
		MaxAge: -1, // expires the cookie
	})
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
