package main

import (
	"log"
)

func main() {
	strChan := make(chan string)

	for i := 0; i < 3; i++ {
		go func() {
			strChan <- "I'm a goroutine!"
		}()
	}

	for m := range strChan {
		log.Println(m)
	}

	log.Println("done")
}
