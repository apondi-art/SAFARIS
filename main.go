package main

import (
	"fmt"
	"net/http"

	"SAFARIS/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandle)
	http.HandleFunc("/joinus", handlers.JoinusHandle)
	http.HandleFunc("/driver", handlers.DriverHandle)
	http.HandleFunc("/user", handlers.UserHandle)
	fmt.Println("server successfully created")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
