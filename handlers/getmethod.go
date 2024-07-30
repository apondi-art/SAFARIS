package handlers

import (
	"net/http"
	"text/template"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method Not Allowed", http.StatusMethodNotAllowed)
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error while parsing The file", http.StatusBadRequest)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "internalServer Error", http.StatusInternalServerError)
	}
}
