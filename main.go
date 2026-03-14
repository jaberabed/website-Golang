package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//Route handler for home page (root URL)
	http.HandleFunc("/", home)
    //Route handler for about page 
	http.HandleFunc("/about", about)
	//Route handler for contact page 
	http.HandleFunc("/contact", contact)
	//Route handler for projects page 
	http.HandleFunc("/projects", projects)
	//Route handler for skills page 
	http.HandleFunc("/skills", skills)
	//Start the server on port 8080

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}

func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmplPath := "templates/" + tmpl
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) 
	}
	t.Execute(w, nil)
}