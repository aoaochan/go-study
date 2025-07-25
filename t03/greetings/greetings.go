package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	var message string = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	var messages = make(map[string]string)

	for _, name := range names {
		var message, err = Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	var formats = []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well me!",
	}

	return formats[rand.Intn(len(formats))]
}
