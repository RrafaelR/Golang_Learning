package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote/v4"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println(quote.Go())
	// Get a greeting message and print it.
	message, err := greetings.Hello("Rafael")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
