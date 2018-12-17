package main

import (
	"fmt"
	"time"
)

var req chan int

func main() {
	//limit is a channel that will hold the processed requests
	limit := make(chan int, 2)
	//ticker is a channel that will receive a value every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	req = make(chan int)
	var willBeEmpty = time.Now().Add(10 * time.Second)
	go func() {
		for {
			select {
			// When the time that we want to limit the resources has passed me empty the limit channel
			case <-ticker.C:
				willBeEmpty = time.Now().Add(10 * time.Second)
				limit = make(chan int, 2)
			case <-req:
				//Here we check if the number of requests we processed is the same as the limit in the specified time range
				if len(limit) != cap(limit) {
					fmt.Println("Limit not reached")
					limit <- 1
				} else {
					fmt.Printf("Limit reached wait %.2f seconds\n", willBeEmpty.Sub(time.Now()).Seconds())
				}
			}
		}
	}()
	// This is just to test if the limiter works
	i := 0
	for {
		req <- i
		i++
		time.Sleep(time.Second)
	}
}
