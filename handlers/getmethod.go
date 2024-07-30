package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	SAFARIS "SAFARIS/functions"
)

type Result struct {
	Driver string
}

var (
	t       = template.Must(template.ParseGlob("./templates/*.html"))
	Drivers SAFARIS.DriverBlock
)

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
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
	name := r.Form.Get("full-name")
	number := r.Form.Get("phone")
	id := r.Form.Get("id-number")
	plate := r.Form.Get("vehicle-plate")
	Drivers.AddDriver(name, id, plate, number)

	var driver string
	for i, block := range Drivers.Drivers {
		driver = fmt.Sprintf("Block %d:\n", i+1)
		driver += fmt.Sprintf("  Hash: %s\n", block.Hash)
		driver += fmt.Sprintf("  Previous Hash: %s\n", block.Name)
		driver += fmt.Sprintf("    Driver ID: %s\n", block.ID)
		driver += fmt.Sprintf("    User: %+v\n", block.VehicleReg)
		driver += fmt.Sprintf("    User: %+v\n", block.TimeStamp)
	}
	data := Result{driver}
	err := t.ExecuteTemplate(w, "driver.html", data)
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
