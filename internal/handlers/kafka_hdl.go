package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"food-siam-si-restaurant/internal/core/ports"

	"github.com/segmentio/kafka-go"
)

type KafkaHandler struct {
	r *kafka.Reader
	s ports.RestaurantService
}

type KafkaMessage struct {
	RestaurantId uint32  `json:"restaurantId"`
	AverageScore float32 `json:"averageScore"`
}

func NewKafkaHandler(r *kafka.Reader, s ports.RestaurantService) *KafkaHandler {
	return &KafkaHandler{
		r: r,
		s: s,
	}
}

func (handler *KafkaHandler) Listen() {
	for {
		m, err := handler.r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var message KafkaMessage

		if err := json.Unmarshal([]byte(m.Value), &message); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
		}

		handler.s.UpdateAverageScore(message.RestaurantId, message.AverageScore)
	}
}
