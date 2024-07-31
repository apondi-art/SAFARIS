package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	SAFARIS "SAFARIS/functions"
)

type Result struct {
	Drivers string
}

var (
	t = template.Must(template.ParseGlob("./templates/*.html"))

	// Create a genesis driver and initialize DriverBlock
	genesisDriver = &SAFARIS.Driver{
		Name:         "Genesis",
		ID:           "GEN!",
		VehicleReg:   "KAA 111A",
		PhoneNumber:  "0000000",
		TimeStamp:    time.Now(),
		PreviousHash: "", // Genesis block has no previous hash
	}

	// Calculate hash for the genesis driver

	// Initialize DriverBlock with the genesis driver
	Drivers = SAFARIS.DriverBlock{
		Drivers: []*SAFARIS.Driver{genesisDriver},
	}
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
	genesisDriver.Hash = SAFARIS.CalculateHash(genesisDriver.Name, genesisDriver.ID, genesisDriver.VehicleReg, genesisDriver.PhoneNumber, genesisDriver.PreviousHash)
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		}
		name := r.Form.Get("full-name")
		fmt.Println(name)
		number := r.Form.Get("phone")
		id := r.Form.Get("id-number")
		plate := r.Form.Get("vehicle-plate")

		// Add driver to the DriverBlock
		Drivers.AddDriver(name, id, plate, number)

		SaveDrivers()

		if err := t.ExecuteTemplate(w, "driver.html", nil); err != nil {
			fmt.Println(err)
		}
	}
	if err := t.ExecuteTemplate(w, "driver.html", nil); err != nil {
		fmt.Println(err)
	}
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if err := LoadDrivers(); err != nil {
		fmt.Println("Failed to load drivers:", err)
	}
	if r.Method == http.MethodGet {
		var name, phone, car string
		var drivers []string
		for i, block := range Drivers.Drivers {
			name += (fmt.Sprintf("Driver	%d:		", i+1))
			name += (fmt.Sprintf("Name:	%s		", block.Name))
			phone += (fmt.Sprintf("Phone Number:	%s		", block.PhoneNumber))
			car += fmt.Sprintf("Vehicle Registration:	%s		", block.VehicleReg)
			car += ("\n") // Adding extra newline for better readability
			drivers = append(drivers, name+" "+phone+" "+car+"\r\n")
			name, car, phone = "", "", ""
		}

		data := Result{strings.Join(drivers, "")} // Set the string builder content in Result

		// Debug print
		fmt.Println(data)

		if err := t.ExecuteTemplate(w, "list.html", data); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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

func SaveDrivers() {
	// Example: Save to a JSON file
	file, err := os.Create("drivers.json")
	if err != nil {
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Drivers); err != nil {
		return
	}
}

func LoadDrivers() error {
	file, err := os.Create("drivers.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Drivers); err != nil {
		return err
	}

	return nil
}
