package main

import "fmt"

type bot interface { // all things that have this property are bot interface
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func runIntroduceInterface() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { // it work's for eb and sb because they'd have the same code into the function,
	// so use interface for don't repeat the func
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello !!"
}

func (spanishBot) getGreeting() string {
	return "Hola!!"
}
