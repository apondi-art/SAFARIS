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
	http.HandleFunc("/list", handlers.ListHandler)
	http.HandleFunc("/bid", handlers.BidHandle)
	fmt.Println("server successfully created")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
