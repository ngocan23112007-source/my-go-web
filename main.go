package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Name string
}

func main() {
	// Phục vụ file tĩnh (CSS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route chính
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{}

		if r.Method == http.MethodPost {
			r.ParseForm()
			data.Name = r.FormValue("name")
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)
	})

	log.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
