package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	log.Println("serving on localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}
