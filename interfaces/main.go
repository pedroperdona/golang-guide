package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func printGreeting(bot bot) {
	fmt.Println(bot.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
