package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	parseToFile()
	// parseToStdout()
}

func parseToFile() {
	// Parse the file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	// Execute the file
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseToStdout() {
	// Parse the file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the file
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
