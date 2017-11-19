package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// We force, after 5 seconds, to cancel the context.
	// We could call this as side effect of a terminated connection to a queue system like RabbitMQ.
	go func() {
		<-time.After(5 * time.Second)
		cancel()
	}()

	messagesChan := listenQueue(ctx)

	readFromChan(messagesChan)

	log.Println("done")
}

func listenQueue(ctx context.Context) chan string {
	messagesChan := make(chan string)
	sendRandomAmount(ctx, messagesChan)

	return messagesChan
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

func sendRandomAmount(ctx context.Context, strChan chan string) {
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(5)

	go func() {
		<-ctx.Done()
		close(strChan)
	}()

	log.Printf("Messages to be sent: %v\n", times)
	for i := 0; i < times; i++ {
		go func(id int) {
			strChan <- fmt.Sprintf("I'm the goroutine %v !", id)
		}(i)
	}
}
