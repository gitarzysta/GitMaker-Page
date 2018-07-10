package main

import (
	"html/template"
	"net/http"
	"./config"
	"log"
)

var(
	Port string
	
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml")) 
}

func main() {
	log.Println("Starting GitMaker's website!!!")
	
	err := config.LoadConfiguration()
	
	if err != nil {
		log.Println(err.Error())
		return
	}
	
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	
	if err != nil {
		log.Println(err.Error())
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "register.gohtml", nil)
	
	if err != nil {
		log.Println(err.Error())
	}
}