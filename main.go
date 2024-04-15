package main

import (
	"os"
	"time"

	mocking "github.com/xergon85/learn-go-with-tests/mocking"
)

func main() {
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency.MyGreetHandler)))
	sleeper := mocking.NewConfigurableSleeper(3*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
