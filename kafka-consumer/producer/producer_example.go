package producer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var Sync = make(chan struct{})

func Producer() {
	fmt.Println("msg from Producer")
	/*
		resp, err := http.Get("http://kafka:9092")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(body)
	*/
	// to produce messages
	topic := "my-topic-1"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	// conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:29092", topic, partition)
	fmt.Println(conn)
	if err != nil {

		log.Fatal("Producer: failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	Sync <- struct{}{}
}
