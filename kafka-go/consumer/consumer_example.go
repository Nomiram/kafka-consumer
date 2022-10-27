package consumer

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/segmentio/kafka-go"
)

func Consumer(s chan struct{}) {
	fmt.Println("msg from Consumer")

	// to consume messages
	topic := "my-topic-1"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		//
		log.Fatal("Comsumer: failed to dial leader:", err)
	}

	s <- struct{}{}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	// i:=0
	for {
		n, _ := batch.Read(b)
		//if err != nil {
		//	break
		//}
		if string(b[:n]) != "" {
			fmt.Println(string(b[:n]))
			break
			// if i >= 3 {break} 
			// else {i=i+3}
		}
	}

	if err := batch.Close(); err != nil {
		// Cons <- struct{}{}
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		// Cons <- struct{}{}
		log.Fatal("failed to close connection:", err)
	}

}
