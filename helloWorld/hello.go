package main

import "fmt"

func main() {
	fmt.Println(Hello(""))
}

const prefixHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefixHello + name
}
