package broker

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Broker struct {
	config *Config
	topic  string

	Msg chan []byte
}

func New(config *Config) *Broker {
	return &Broker{
		config: config,
	}
}
func (b *Broker) Consumer() error {
	consumer, err := sarama.NewConsumer([]string{b.config.Url}, nil)

	if err != nil {
		return err
	}
	defer consumer.Close()
	b.topic = "test"
	partConsumer, err := consumer.ConsumePartition("ps", 1, sarama.OffsetNewest)
	for {
		select {
		case m, ok := <-partConsumer.Messages():
			if !ok {
				return err
			}

			fmt.Println(m.Value)
		}
	}
}

func (b *Broker) Producer(a string) error {
	producer, err := sarama.NewSyncProducer([]string{b.config.Url}, nil)
	if err != nil {
		return nil
	}
	b.topic = "test"
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: b.topic,
		Key:   sarama.StringEncoder(1),
		Value: sarama.StringEncoder(a),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
