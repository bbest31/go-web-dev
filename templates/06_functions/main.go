package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// create a FuncMap to register functions.
// "uc" is the what the func will be called in the template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type person struct {
	Name string
	Age  uint
}

func main() {
	p := person{"Brandon", 27}
	p2 := person{"Emma", 23}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", []person{p, p2})
	if err != nil {
		log.Fatalln(err)
	}
}
