package infra

import (
	"github.com/segmentio/kafka-go"
	"nightingale/config"
)

type KafkaClient struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
}

func NewKafkaClient(cfg config.KafkaConfig) (*KafkaClient, error) {
	client := &KafkaClient{}
	client.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Brokers,
		GroupID:  cfg.GroupID,
		Topic:    cfg.Topic,
		MaxBytes: 10 << 20,
	})

	client.Writer = &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Brokers...),
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	return client, nil
}
