package main

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"database/sql"
	"net/http"
	"./config"
	"log"
)

var(
	Port string
	
	tpl *template.Template
	
	DataPort string
	DataUser string
	DataPassword string
	DataBaseName string
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
	
	DataPort = config.Port
	DataUser = config.User
	DataPassword = config.Password
	DataBaseName = config.BaseName
	
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	
	log.Println("Connecting to database: ")
	log.Printf("Port: %s", DataPort)
	log.Printf("User: %s", DataUser)
	log.Printf("Password: %s", DataPassword)
	log.Printf("BaseName: %s", DataBaseName)
	
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/osadnicy")
	
	if err != nil {
		panic(err.Error())
	}
	
	http.ListenAndServe(":8080", nil)
	
	defer log.Println("Switching off...")
	defer log.Println("Disconnecting with base...")
	defer db.Close()
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