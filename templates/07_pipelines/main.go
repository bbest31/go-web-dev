package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func double(n int) int {
	return n * 2
}

func sqr(n int) int {
	return n * n
}

// create a FuncMap to register functions.
// "uc" is the what the func will be called in the template
var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"dbl":      double,
	"sqr":      sqr,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 10)
	if err != nil {
		log.Fatalln(err)
	}
}
