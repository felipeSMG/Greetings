package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")

	}
	message := fmt.Sprintf("Hola, %v", name)
	return message, nil
}
