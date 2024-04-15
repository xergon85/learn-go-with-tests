package main

import (
	"log"
	"net/http"

	dependency "github.com/xergon85/learn-go-with-tests/dependency"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency.MyGreetHandler)))
}
