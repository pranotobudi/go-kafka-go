package main

// reference: https://towardsdev.com/go-kafka-go-7547c38b4f09

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "foo"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:19092", topic, partition)
	if err != nil {
		fmt.Println("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		fmt.Println("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		fmt.Println("failed to close writer:", err)
	}
	fmt.Println("Conn Close")
	fmt.Println(err)
}
