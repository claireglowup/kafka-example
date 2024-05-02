package main

import (
	"bydoc/producer"
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {

	topic := "my-topic"

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2023-01-02 14:24:04"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Balancer: &kafka.LeastBytes{},
	}

	writer := kafka.NewWriter(kafkaConfig)
	defer writer.Close()

	kafka := &producer.KafkaProducer{
		Producer: writer,
	}

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("Message number %v", i)
		err := kafka.SendMessage(context.Background(), topic, msg)
		if err != nil {
			panic(err)
		}
	}

}
