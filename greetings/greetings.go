package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	// := declare and initialize a value
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}