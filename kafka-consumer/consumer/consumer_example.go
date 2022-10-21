package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var Cons = make(chan struct{})

func Consumer() {
	fmt.Println("msg from Consumer")

	// to consume messages
	topic := "my-topic-1"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		// Cons <- struct{}{}
		log.Fatal("Comsumer: failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
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