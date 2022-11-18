package main

import "fmt"

const englishHelloPrefix = "Hello, "

const japanese = "Japanese"
const japaneseHelloPrefix = "こんにちは、"

const french = "French"
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("World!", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	}

	return prefix
}
