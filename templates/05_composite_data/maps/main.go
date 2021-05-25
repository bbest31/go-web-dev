package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", map[int]string{1: "Apples", 2: "Oranges", 3: "Bananas"})
	if err != nil {
		log.Fatalln(err)
	}
}
