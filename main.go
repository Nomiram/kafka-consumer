package main

import (
	"fmt"
	"kafka-consumer/consumer"
	"kafka-consumer/producer"
)

func main() {
	fmt.Println("test")
	go producer.Producer()
	<-producer.Sync
	go consumer.Consumer()
	<-consumer.Cons
	fmt.Println("done")
}
