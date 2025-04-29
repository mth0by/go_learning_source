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

	names := [3]string{"", "Vladus"} // ["", "Vladus", ""]
	for _, name := range names {
		massage, err := greetings.Hello(name)

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(massage)
	}
}
