package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crosscutgymnast/personal-portfolio/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	log.Fatal(http.ListenAndServe(portNumber, nil))

}
