package main

import (
	"fmt"
	"net/http"

	"SAFARIS/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandle)
	fmt.Println("server successfully created")
	http.ListenAndServe(":8080", nil)
}
