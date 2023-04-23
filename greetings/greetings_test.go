package greetings

import (
	"testing"
)

func Test_sayHello(t *testing.T) {
	name := "Bob"
	want := "Hi Bob"

	if got := SayHello(name); got != want {
		t.Errorf("sayHello(%q) = %q, want %q", name, got, want)
	}
}
