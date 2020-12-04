// Package greetings provides simple functions to greet named individuals.
package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Greeting struct {
	Name           string // the name of the individual being greeted
	Greeting       string // the complete greeting to the named individual
	greetingFormat string // the format string on which the greeting is based (unexported)
}

var formats []string = []string{
	"Hello %v, my dear.",
	"Good morning %v!",
	"%v! Great to see you!",
}

var formatsLen int = len(formats)

func init() {
	log.Print("Initializing file greet.go")
	rand.Seed(time.Now().UnixNano())
}

// Greet returns a random greeting to the named individual, or an error if no name is given.
func Greet(name string) (Greeting, error) {
	if name == "" {
		return Greeting{}, errors.New("Missing name")
	}
	gf := randomGreetingFormat()
	g := Greeting{name, fmt.Sprintf(gf, name), gf}
	return g, nil
}

// GreetAll returns random greetings to the named individuals, as a map from name to greeting,
// or an error if any name is missing.
func GreetAll(names []string) (map[string]Greeting, error) {
	gs := make(map[string]Greeting)
	for _, n := range names {
		g, err := Greet(n)
		if err != nil {
			return nil, err
		}
		gs[n] = g
	}
	return gs, nil
}

// A random greeting as a format suitable for the *Printf functions with one placeholder for the name.
func randomGreetingFormat() string {
	return formats[rand.Intn(formatsLen)]
}
