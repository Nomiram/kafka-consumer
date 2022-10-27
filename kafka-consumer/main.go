package main

import (
	"fmt"
	"kafka-consumer/consumer"
	"kafka-consumer/producer"
	"sync"
)

func main() {
	fmt.Println("test")
	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	wg.Add(2)
	go func() {
		consumer.Consumer(ch)
		wg.Done()
	}()
	go func() {
		producer.Producer(ch)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}
