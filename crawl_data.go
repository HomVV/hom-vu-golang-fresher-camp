package main

import (
	"fmt"
	"time"
)

func crawl(number int, nameOfRoutine int) {
	fmt.Printf("Routine %d Crawling %d \n", nameOfRoutine, number)
}
func main() {

	const numberOfRequests int = 1000
	data := make(chan int, numberOfRequests)
	defer close(data)
	for i := 1; i <= numberOfRequests; i++ {
		data <- i
	}
	for i := 1; i <= 5; i++ {
		go func(nameOfRoutine int) {
			for value := range data {
				crawl(value, nameOfRoutine)
			}
			fmt.Printf("Routine %d is done\n", nameOfRoutine)
		}(i)
	}
	time.Sleep(time.Second)
}
