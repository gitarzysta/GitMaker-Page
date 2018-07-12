package main

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"database/sql"
	"net/http"
	"./config"
	"bytes"
	"log"
)

var(
	Port string
	
	tpl *template.Template
	
	DataPort string
	DataUser string
	DataPassword string
	DataBaseName string
	
	buffer bytes.Buffer
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html")) 
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
	
	buffer.WriteString(DataUser)
	buffer.WriteString(":@tcp(127.0.0.1:3306)/")
	buffer.WriteString(DataBaseName)

	db, err := sql.Open("mysql", buffer.String())
	
	if err != nil {
		panic(err.Error())
	}else {
		log.Println("Succefully connected with database!")
	}
	
	http.ListenAndServe(":8080", nil)
	
	defer log.Println("Switching off...")
	defer log.Println("Disconnecting with base...")
	defer db.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	
	if err != nil {
		log.Println(err.Error())
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "register.html", nil)
	
	if err != nil {
		log.Println(err.Error())
	}
	
	if template.HTMLEscapeString(r.Form.Get("nicknameField")) == "" {
		log.Println("NOPE")
	}
	
	log.Println("username:", template.HTMLEscapeString(r.Form.Get("nicknameField")))
}