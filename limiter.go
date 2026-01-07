package main

import (
	"fmt"
	"time"
)

var lastRequest time.Time

func allowRequest() bool {
	if time.Since(lastRequest).Seconds() > 2 {
		lastRequest = time.Now()
		return true
	}
	return false
}

func main() {
	for i := 0; i < 5; i++ {
		if allowRequest() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Too many requests")
		}
		time.Sleep(1 * time.Second)
	}
}
