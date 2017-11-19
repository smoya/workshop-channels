package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	strChan := make(chan string)

	sendRandomAmount(strChan)

loop:
	for {
		select {
		case m, ok := <-strChan:
			if !ok {
				// channel closed
				break loop
			}
			log.Println(m)
		}
	}

	log.Println("done")
}

func sendRandomAmount(strChan chan string) {
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(5)

	wg := sync.WaitGroup{}
	wg.Add(times)

	// once all the goroutines are done, we close the channel.
	go func() {
		wg.Wait()
		close(strChan)
	}()

	log.Printf("Messages to be sent: %v\n", times)
	for i := 0; i < times; i++ {
		go func(id int) {
			defer wg.Done()
			strChan <- fmt.Sprintf("I'm the goroutine %v !", id)
		}(i)
	}
}
