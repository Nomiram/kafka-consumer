package main

import (
	"fmt"
	"kafka-consumer/consumer"
	"kafka-consumer/producer"
)

func main() {
	fmt.Println("test")
	go consumer.Consumer()
	go producer.Producer()
	<-consumer.Cons
	fmt.Println("done")
}
