package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func Hello(name, lang string) string {
	if name == "" {
		name = "World!"
	}
	prefix := mapPrefix(lang)

	return prefix + name
}

func mapPrefix(lang string) (prefix string) { // private functions are in lowercase named return will create prefix variable with "" and be returning prefix
	switch lang {

	case "Spanish":
		prefix = spanishHelloPrefix

	case "French":
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World!", ""))
}
