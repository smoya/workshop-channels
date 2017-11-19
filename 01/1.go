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

	log.Println(<-strChan)
	log.Println(<-strChan)
	log.Println(<-strChan)

	log.Println("done")
}
