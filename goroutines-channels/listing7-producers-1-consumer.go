package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	payments := make(chan int)
	texts := make(chan string)
	quit := make(chan bool)
	go consume(payments, texts, quit)
	go producePayments(payments)
	go produceTexts(texts)
	time.Sleep(5 * time.Second)
	quit <- true
	time.Sleep(1 * time.Second)
}
func produceTexts(strings chan string) {
	for i := 1; i < 7; i++ {
		strings <- "message " + strconv.Itoa(i)
		time.Sleep(800 * time.Millisecond)
	}
}

func producePayments(ints chan int) {
	for i := 1; i < 9; i++ {
		ints <- i * 13
		time.Sleep(500 * time.Millisecond)
	}
}

func consume(payments chan int, texts chan string, quit chan bool) {
	for {
		select {
		case p := <-payments:
			fmt.Println("got a payment of ", p, "euro")
		case t := <-texts:
			fmt.Println("got a text: ", t)
		case <-quit:
			fmt.Println("quitting, bye")
			return
		}
	}
}
