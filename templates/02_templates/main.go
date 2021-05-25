package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// init() is called before main automatically by go.
func init() {
	// Must handles error checking so we parse our templates once and checks for issues
	// Parse all files in the specified path with the specific file extension
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// Can execute any template
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute a specfic template
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
