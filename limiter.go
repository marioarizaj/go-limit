package main

import (
	"fmt"
	"time"
)

var req chan int

func main() {
	limit := make(chan int, 2)
	ticker := time.NewTicker(10 * time.Second)
	req = make(chan int)
	var willBeEmpty = time.Now().Add(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				willBeEmpty = time.Now().Add(10 * time.Second)
				limit = make(chan int, 2)
			case <-req:
				if len(limit) != cap(limit) {
					fmt.Println("Limit not reached")
					limit <- 1
				} else {
					fmt.Printf("Limit reached wait %.2f seconds\n", willBeEmpty.Sub(time.Now()).Seconds())
				}
			}
		}
	}()
	i := 0
	for {
		req <- i
		i++
		time.Sleep(time.Second)
	}
}
