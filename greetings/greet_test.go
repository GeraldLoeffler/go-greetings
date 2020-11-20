package greetings

import (
	"regexp"
	"testing"
)

func TestGreet(t *testing.T) {
	n := "Gerald"
	want := regexp.MustCompile("\\b" + n + "\\b")
	g, err := Greet(n)
	if err != nil || !want.MatchString(g) {
		t.Fatalf("Greet(%v)=%q,%v want match to regex", n, g, err)
	}
}

func TestGreetEmpty(t *testing.T) {
	n := ""
	g, err := Greet(n)
	if err == nil || g != "" {
		t.Fatalf("Greet(%v)=%q,%v want error", n, g, err)
	}
}
