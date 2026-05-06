package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")

	}
	message := fmt.Sprintf("Hola, %v", name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %v. Bienvenido!",
		"Saludos, %v!",
		"Hola, %v! Que tengas un gran día!",
	}
	return formats[rand.Intn(len(formats))]
}
