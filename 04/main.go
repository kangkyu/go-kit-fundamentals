package main

import (
	"fmt"
)

func main() {
	fmt.Println(getGreeting())
}

func getGreeting() string {
	return fmt.Sprintf("Welcome to Go kit 0.12 Fundamentals!")
}
