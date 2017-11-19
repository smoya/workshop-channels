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

	messagesChan := listenQueue(ctx)

	readFromChan(messagesChan)

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
