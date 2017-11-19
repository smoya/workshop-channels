package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	strChan := make(chan string)

	sendRandomAmount(strChan)

	for m := range strChan {
		log.Println(m)
	}

	log.Println("done")
}

func sendRandomAmount(strChan chan string) {
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(5)

	log.Printf("Messages to be sent: %v\n", times)
	for i := 0; i < times; i++ {
		go func(id int) {
			strChan <- fmt.Sprintf("I'm the goroutine %v !", id)
		}(i)
	}
}
