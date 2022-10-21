package producer

import (
	"context"
	"fmt"
	"log"

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
	// topic := "my-topic-1"
	topic := "test"
	// partition := 0
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr: kafka.TCP("kafka:9092"),
		// Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	/*
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
	*/
	Sync <- struct{}{}
}
