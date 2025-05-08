package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(quote.Hello())

	names := []string{"Nicholas", "Andrew", "Boros", ""}

	msgs, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, msg := range msgs {
		fmt.Println(msg)
	}
}
