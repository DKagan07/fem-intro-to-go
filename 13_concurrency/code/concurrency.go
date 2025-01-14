package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 300)
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC")
	}
}

func printStuff() {
	defer wg.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	// go say("Hello")
	// say("There")
	wg.Add(1)
	go printStuff()
	go say("hello")
	wg.Add(1)
	go printStuff()
	go say("world")
	wg.Wait()
}
