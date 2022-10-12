package client

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	w *kafka.Writer
}

func NewKafkaClient(brokers []string, topic string, groupID string) *KafkaClient {
	client := &KafkaClient{}
	client.w = kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	return client
}

func (k *KafkaClient) SendMessage(ctx context.Context, key string, msg interface{}) (err error) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}
	return k.w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: []byte(jsonData),
	})
}

func (k *KafkaClient) Close() (err error) {
	err = k.w.Close()
	if err != nil {
		return
	}
	return
}
