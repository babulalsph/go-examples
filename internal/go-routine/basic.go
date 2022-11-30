package main

import (
	"fmt"
	"time"
)

// Concurrency : Concurrency is about dealing with lots of things at once.
// Parallelism : Parallelism is about doing lots of things at once.
// Concurrency is about Structure But Parallelism is about Execution

func HelloFunc(msg string, delay time.Duration) {
	for {
		fmt.Println(msg)
		time.Sleep(delay)
	}
}

func main() {

	go HelloFunc("A--", 200*time.Millisecond)
	go HelloFunc("--B--", 500*time.Millisecond)
	go HelloFunc("--C", 900*time.Millisecond)

	time.Sleep(20 * time.Second)
}
