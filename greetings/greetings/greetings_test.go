package greetings

import (
	"testing"
	"regexp"
)

func TestHello(t *testing.T) {
	name := "Petr"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Errorf(`Hello(%q) = %v, %v, want match for "", nil`, name, msg, err)
	}	
}
