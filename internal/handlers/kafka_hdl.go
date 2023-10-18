package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaHandler struct {
}

func NewKafkaHandler() *KafkaHandler {
	return &KafkaHandler{}
}

func (handler KafkaHandler) Listen() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092"},
		Topic:     "mytopic",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
		GroupID:   "consumer-group-id",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
