package main

import (
	"fmt"
	"log"
	"github.com/felipeSMG/greetings"
)

func main() {
	message, err := greetings.Hello("Mundo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}