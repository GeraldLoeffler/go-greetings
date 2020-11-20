package greetings

import (
	"regexp"
	"testing"
)

func TestGreetHappy(t *testing.T) {
	cases := []struct {
		in   string
		want *regexp.Regexp
	}{
		{"Gerald", regexp.MustCompile("\\bGerald\\b")},
		{"Dagmar", regexp.MustCompile("\\bDagmar\\b")},
	}
	for _, c := range cases {
		got, err := Greet(c.in)
		if err != nil || !c.want.MatchString(got) {
			t.Errorf("Greet(%q) = (%q, %q), want match to %q", c.in, got, err, c.want)
		}
	}
}

func TestGreetError(t *testing.T) {
	n := ""
	got, err := Greet(n)
	if err == nil || got != "" {
		t.Fatalf("Greet(%q) = (%q, %q), want error", n, got, err)
	}
}
