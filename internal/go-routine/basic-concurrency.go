package main

import "fmt"

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go generateNumbers(numbers)
	go calculateSquare(numbers, squares)

	printResponse(squares)

}

//The type chan<- int, a send-only channel of int, allows sends but not receives.
func generateNumbers(out chan<- int) {
	var i int
	for i = 0; i <= 200; i++ {
		fmt.Println("generate : ", i)
		out <- i
	}
	close(out)
}

// req <-chan int a receive-only channel of int, allow receives data only.
// out chan<- int a send-only channel of int, allow sends data only.
func calculateSquare(req <-chan int, out chan<- int) {
	for {
		num, ok := <-req
		if !ok {
			break
		}
		//fmt.Println("square : ", num)
		out <- num * num
	}
	close(out)
}

//The type <-chan int, a receive-only channel of int, allows receives but not sends.
func printResponse(req <-chan int) {
	for val := range req {
		fmt.Println("Square :", val)
	}
}
