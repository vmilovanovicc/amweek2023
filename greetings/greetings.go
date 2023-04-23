package greetings

import "fmt"

// SayHello prints greetings
func SayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
