package greetings

import (
	"reflect"
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
		if err != nil || got.Name != c.in || !c.want.MatchString(got.Greeting) || c.want.MatchString(got.greetingFormat) {
			t.Errorf("Greet(%q) = (%q, %q), want match to %q", c.in, got, err, c.want)
		}
	}
}

func TestGreetError(t *testing.T) {
	n := ""
	got, err := Greet(n)
	if err == nil || !reflect.ValueOf(got).IsZero() {
		t.Fatalf("Greet(%q) = (%q, %q), want error", n, got, err)
	}
}
