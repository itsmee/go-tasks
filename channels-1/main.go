package main

import (
	"fmt"
)

func main() {
	in := make(chan string)
	out := make(chan string)

	items := []string{"a", "b", "c", "d", "e"}

	fmt.Println(in, out, items)
	// TODO: process all the items in a separate goroutine,
	// this goroutine accepts the messages from the in channel,
	// capitalizes each item, and sends the processed items to the out channel

	// TODO: read the processed items from the out channel
	// and print the processed items
	// (hint: what happens if we read them in the main thread? what if we read them in goroutine?)

	// TODO: program should exit once all the items are processed and pritned
}

func process(in chan string, out chan string) {
	// implement me
}

func output(out chan string) {
	// implement me
}
