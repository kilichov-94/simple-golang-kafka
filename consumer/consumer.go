package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic-test", 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 3))

	// message, _ := conn.ReadMessage(1e6) its read only one message
	// fmt.Println(string(message.Value))

	batch := conn.ReadBatch(1e3, 1e9) // its read all messages
	bytes := make([]byte, 1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
