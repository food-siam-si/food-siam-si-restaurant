package messagequeue

import (
	"food-siam-si-restaurant/config"

	"github.com/segmentio/kafka-go"
)

func NewKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.GetKafka().Broker},
		Topic:     config.GetKafka().Topic,
		Partition: 0,
		MaxBytes:  1e6,
		GroupID:   config.GetKafka().GroupId,
	})
}
