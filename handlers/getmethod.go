package handlers

import (
	"net/http"
	"text/template"
)
var t = template.Must(template.ParseGlob("./templates/*.html"))
	
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method Not Allowed", http.StatusMethodNotAllowed)
	}
	
	err := t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "internalServer Error", http.StatusInternalServerError)
	}
}

func JoinusHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method Not Allowed", http.StatusMethodNotAllowed)
	}
	
	err := t.ExecuteTemplate(w, "joinus.html", nil)
	if err != nil {
		http.Error(w, "internalServer Error", http.StatusInternalServerError)
	}
}

func DriverHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method Not Allowed", http.StatusMethodNotAllowed)
	}
	
	err := t.ExecuteTemplate(w, "driver.html", nil)
	if err != nil {
		http.Error(w, "internalServer Error", http.StatusInternalServerError)
	}
}

func UserHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method Not Allowed", http.StatusMethodNotAllowed)
	}
	
	err := t.ExecuteTemplate(w, "user.html", nil)
	if err != nil {
		http.Error(w, "internalServer Error", http.StatusInternalServerError)
	}
}