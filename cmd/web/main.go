package main

import (
	"fmt"
	"net/http"

	"github.com/TartuDen/HelloWorldWebApp/pkg/handler"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handler.MainHandler)
	http.HandleFunc("/about", handler.AboutHandler)

	fmt.Println("Running on port: ", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error listening port", port, err)
	}
}
