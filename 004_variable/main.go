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
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Be true to yourself, live a life of enrichment and share it with people you love.`)
	if err != nil {
		log.Fatalln(err)
	}
}
