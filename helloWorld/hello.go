package main

import (
	"hello/mock"
	"os"
	"time"
)

func main() {
	// fmt.Println(Hello(""))
	sleeper := &mock.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mock.Countdown(os.Stdout, sleeper)
}

const prefixHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefixHello + name
}
