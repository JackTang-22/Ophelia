package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1000)
	cfg := nsq.NewConfig()
	q, _ := nsq.NewConsumer("test", "ch", cfg)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("get a message: %s", message.Body)
		wg.Done()
		return nil
	}))

	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic(err)
	}
	wg.Wait()
}
