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

	readFromChan(strChan)

	log.Println("done")
}

func readFromChan(channel chan string) {
	for {
		select {
		case m, ok := <-channel:
			if !ok {
				// channel closed
				return
			}
			log.Println(m)
		}
	}
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
