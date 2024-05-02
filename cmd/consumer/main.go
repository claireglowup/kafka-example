package main

import (
	"bydoc/consumer"
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {

	topic := "my-topic"

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2023-01-02 14:24:04"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := kafka.ReaderConfig{
		Brokers:     []string{"kafka:9092"},
		Topic:       topic,
		Partition:   0,
		MaxBytes:    10e3,
		StartOffset: kafka.LastOffset,
	}

	reader := kafka.NewReader(kafkaConfig)
	defer reader.Close()

	kafkaConsumer := &consumer.KafkaConsumer{
		Consumer: reader,
	}

	kafkaConsumer.Consume(context.Background())

}
