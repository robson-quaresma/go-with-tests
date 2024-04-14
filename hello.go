package main

import "fmt"

const (
	spanish = "Spanish"
	frech   = "French"

	frechHelloPrefix   = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix // Default to english

	switch language {
	case frech:
		prefix = frechHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("word", "english"))
}
