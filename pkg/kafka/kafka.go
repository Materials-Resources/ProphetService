package kafka

import (
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
)

func NewProducerClient(brokers []string, kotel *kotel.Kotel) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(brokers...),
		kgo.WithHooks(kotel.Hooks()...),
	}
	return kgo.NewClient(opts...)
}

func NewConsumerClient(brokers []string, topic string) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(brokers...),
		kgo.ConsumeTopics(topic),
	}
	return kgo.NewClient(opts...)
}
