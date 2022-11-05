package main

import (
	"fmt"
	"kafka-consumer/consumer"
	"kafka-consumer/producer"
	"sync"
	"time"
)

func main() {
	fmt.Println("test")
	conn := consumer.KafkaConsumer()
	defer conn.Close()
	w := producer.Producer()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 1000)
			fmt.Println("try")
			key, value := consumer.KafkaRead(conn)
			fmt.Println(i, " key: ", key, "value: ", value)
		}
		wg.Done()
	}()
	// time.Sleep(time.Second * 10)
	go func() {
		producer.Write(w, "Key-A", "Hello World11")
		producer.Write(w, "Key-B", "Hello World22")
		producer.Write(w, "Key-C", "Hello World33")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}
