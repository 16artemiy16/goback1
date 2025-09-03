package main

import "fmt"

func printHi(ch chan int) {
	fmt.Println("HI!")
	ch <- 1
}

func main() {
	ch := make(chan int)
	go printHi(ch)
	go func() {
		fmt.Println("INNER HI!!!!!!!")
	}()
	//<-ch
	//<-ch
	fmt.Println("FINISH!")
}
