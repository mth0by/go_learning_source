package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hellos(names []string) ([]string, error) {
	if len(names) == 0 {
		return []string{}, errors.New("names list is empty")
	}

	var greetings []string
	for _, name := range names {
		greet, err := Hello(name)
		if err != nil {
			continue
		}

		greetings = append(greetings, greet)
	}

	return greetings, nil

}

func Hello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf(getRandomGreeting(), name)

	return message, nil 
}

func getRandomGreeting() string {
	choices := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v",
		"Hail, %v! Well met!",
		"Ricorda, %v, dove vedrai",
	}

	return choices[rand.Intn(len(choices))]
}
