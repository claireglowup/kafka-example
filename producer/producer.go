package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	Producer *kafka.Writer
}

func (p *KafkaProducer) SendMessage(ctx context.Context, topic, msg string) error {

	kafkaMsg := kafka.Message{
		Topic: topic,
		Value: []byte(msg),
	}

	if err := p.Producer.WriteMessages(ctx, kafkaMsg); err != nil {
		logrus.Errorf("Send message error: %v", err)
		return err
	}

	logrus.Infof("Send Message succes, Topic %v", topic)
	return nil

}
