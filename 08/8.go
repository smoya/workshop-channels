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

	wg := sync.WaitGroup{}
	sendRandomAmount(&wg, strChan)
	wg.Wait()

	// Are we reaching this point?
	log.Println("All gouroutines are finished!")

	for m := range <-strChan {
		log.Println(m)
	}

	log.Println("done")
}

func sendRandomAmount(wg *sync.WaitGroup, strChan chan string) {
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(5)

	wg.Add(times)

	log.Printf("Messages to be sent: %v\n", times)
	for i := 0; i < times; i++ {
		go func(id int) {
			defer wg.Done()
			strChan <- fmt.Sprintf("I'm the goroutine %v !", id)
		}(i)
	}

}
