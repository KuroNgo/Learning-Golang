package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello2(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// một slice của message định dạng.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	//a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	// a map to asociate names with mesages
	messages := make(map[string]string)
	//Loop through the received slice of names, calling
	// The Heello function to get a mesage for each name
	for _, name := range names {
		message, err := Hello2(name)

		// Nil được dùng để kiểm tra xem một con trỏ có trỏ đến đối tượng hay không
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}
