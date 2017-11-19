package main

import (
	"fmt"
	"log"
)

func main() {
	strChan := make(chan string)

	for i := 0; i < 3; i++ {
		go func(id int) {
			strChan <- fmt.Sprintf("I'm the goroutine %v !", id)
		}(i)
	}

	for i := 0; i < 3; i++ {
		log.Println(<-strChan)
	}

	log.Println("done")
}
