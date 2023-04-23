package main

import (
	"fmt"
	"random/greetings"
)

func main() {
	msg := greetings.SayHello("Alice")
	fmt.Println(msg)
}
