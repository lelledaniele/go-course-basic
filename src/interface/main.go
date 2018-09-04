package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}

	printGreeting(eb)
}

func (englishBot) getGreeting() string {
	return "ciao"
}

func (englishBot) ciao() string {
	return "ciao"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
