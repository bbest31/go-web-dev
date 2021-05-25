package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type person struct {
	Name string
	Age  uint
}

func (p person) Greet() string {
	return fmt.Sprintf("Hello my name is %v\n", p.Name)
}

func main() {
	p := person{"Brandon", 27}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
