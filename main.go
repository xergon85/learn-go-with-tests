package main

import (
	"log"
	"net/http"

	dependacy "github.com/xergon85/learn-go-with-tests/dependecy"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependacy.MyGreetHandler)))
}
