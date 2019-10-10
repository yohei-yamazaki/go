package main

import "fmt"

func main() {
	// dumplingList:= 0
	fmt.Println("Start")

	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	go func() {
		push()
		ch1 <- true
	}()
	<-ch1

	go func() {
		pop()
		ch2 <- true
	}()
	<-ch2

	go func() {
		pop()
		ch3 <- true
	}()
	<-ch3

	fmt.Println("Finish!")
}

func push() {
	fmt.Println("push")
}

func pop() {
	fmt.Println("pup")
}
