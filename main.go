package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.PathPrefix("/node_modules/").Handler(http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/form-project", formProject).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")
	route.HandleFunc("/contact", contact).Methods("GET")

	fmt.Println("server on")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("view/index.html")

	if err != nil {
		w.Write([]byte("messege: " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func formProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("view/add.html")

	if err != nil {
		w.Write([]byte("messege: " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	var tech []string
	for key, values := range r.Form {
		for _, value := range values {
			if key == "technologies" {
				tech = append(tech, value)
			}
		}
	}

	fmt.Println("Nama Project : " + r.PostForm.Get("projectName"))
	fmt.Println("Start : " + r.PostForm.Get("startDate"))
	fmt.Println("End : " + r.PostForm.Get("endDate"))
	fmt.Println("Desc : " + r.PostForm.Get("Description"))
	fmt.Println("Tech :", tech)

	http.Redirect(w, r, "/form-project", http.StatusMovedPermanently)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("view/contact.html")

	if err != nil {
		w.Write([]byte("messege: " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}
