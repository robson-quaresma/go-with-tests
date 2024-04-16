package main

import (
	"log"
	"net/http"

	dependencyinjection "github.com/robson-quaresma/go-with-tests/dependency_injection"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
