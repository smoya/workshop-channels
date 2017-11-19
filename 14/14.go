package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ensureInterruptionsStopApplication(cancel)

	messagesChan := listenQueue()

	readFromChan(ctx, messagesChan)

	log.Println("done")
}

func ensureInterruptionsStopApplication(cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Printf(fmt.Sprintf("Got signal %s. Canceling the context...\n", s))
		cancelFunc()
	}()
}

func listenQueue() chan string {
	messagesChan := make(chan string)
	sendRandomAmount(messagesChan)

	return messagesChan
}

func readFromChan(ctx context.Context, channel chan string) {
	for {
		select {
		case <-ctx.Done():
			log.Println("context was canceled")
			return
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

	// We don't care here at this example about closing the channel.

	log.Printf("Messages to be sent: %v\n", times)
	for i := 0; i < times; i++ {
		go func(id int) {
			strChan <- fmt.Sprintf("I'm the goroutine %v !", id)
		}(i)
	}
}
