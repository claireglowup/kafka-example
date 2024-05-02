package consumer

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer *kafka.Reader
}

func (c *KafkaConsumer) Consume(ctx context.Context) {

	for {

		message, err := c.Consumer.ReadMessage(ctx)
		if err != nil {
			logrus.Errorf("Unable to consume message. Error : %v", err)
			break
		}
		logrus.Infof("Message from kafka : %v", string(message.Value))

	}

}
