package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var formats []string = []string{
	"Hello %v, my dear.",
	"Good morning %v!",
	"%v! Great to see you!",
}

var formatsLen int = len(formats)

func init() {
	log.Print("Initializing package greetings")
	rand.Seed(time.Now().UnixNano())
}

func Greet(name string) (string, error) {
	if name == "" {
		return *new(string), errors.New("Missing name")
	}
	return fmt.Sprintf(randomGreeting(), name), nil
}

func GreetAll(names []string) (map[string]string, error) {
	gs := make(map[string]string)
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
func randomGreeting() string {
	return formats[rand.Intn(formatsLen)]
}
