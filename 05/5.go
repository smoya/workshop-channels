package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	strChan := make(chan string)

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)
		go func() {
			strChan <- fmt.Sprintf("I'm the goroutine %v !", i)
		}()
	}

	for i := 0; i < 3; i++ {
		log.Println(<-strChan)
	}

	log.Println("done")
}
