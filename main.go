package main

import (
	"fmt"
	"net/http"

	"SAFARIS/handlers"
)

func main() {
	PORT := ":8080"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", handlers.HomeHandle)
	http.HandleFunc("/joinus", handlers.JoinusHandle)
	http.HandleFunc("/driver", handlers.DriverHandle)
	http.HandleFunc("/user", handlers.UserHandle)
	fmt.Printf("server running in http://localhost%s/", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}
