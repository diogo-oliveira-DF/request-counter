package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diogo-oliveira-DF/request-counter/service"
)

func main() {
	service.LoadSavedData()

	http.HandleFunc("/", service.Handler)

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
