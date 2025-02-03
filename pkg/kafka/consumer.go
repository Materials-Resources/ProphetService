package kafka

import (
	"context"
	"github.com/twmb/franz-go/pkg/kgo"
	"sync"
)

type WorkerFunc func(rec *kgo.Record) error

type ConsumerGroup struct {
	topics []string

	Workers   map[string]WorkerFunc
	Consumers map[topicPartition]*consumer
}

func NewConsumerGroup() ConsumerGroup {
	return ConsumerGroup{
		Workers:   make(map[string]WorkerFunc),
		Consumers: make(map[topicPartition]*consumer),
	}
}

func (cg *ConsumerGroup) AddWorker(topic string, function WorkerFunc) {
	cg.topics = append(cg.topics, topic)
	cg.Workers[topic] = function
}

func (cg *ConsumerGroup) Consume(clientID string, groupID string, factory Factory) error {
	opts := []kgo.Opt{
		kgo.ConsumerGroup(groupID),
		kgo.ConsumeTopics(cg.topics...),
		kgo.OnPartitionsAssigned(cg.assigned),
		kgo.OnPartitionsRevoked(cg.lost),
		kgo.OnPartitionsLost(cg.lost),
		kgo.DisableAutoCommit(),
		kgo.BlockRebalanceOnPoll(),
	}

	cl, err := factory.NewKafkaClient(clientID, opts...)
	if err != nil {
		return err
	}

	go func() {
		cg.pollTopics(cl)
	}()

	return nil
}

// assigned is called when the consumer group is assigned partitions.
func (cg *ConsumerGroup) assigned(_ context.Context, _ *kgo.Client, assigned map[string][]int32) {
	for topic, partitions := range assigned {
		for _, partition := range partitions {
			topicPartition := topicPartition{topic: topic, partition: partition}
			consumer := &consumer{
				quit:       make(chan struct{}),
				done:       make(chan struct{}),
				recs:       make(chan []*kgo.Record, 5),
				workerFunc: cg.Workers[topic],
			}
			cg.Consumers[topicPartition] = consumer
			go consumer.Consume()
		}
	}
}

// revoked is called when the consumer group is revoked partitions.
func (cg *ConsumerGroup) lost(_ context.Context, _ *kgo.Client, revoked map[string][]int32) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for topic, partitions := range revoked {
		for _, partition := range partitions {
			topicPartition := topicPartition{topic: topic, partition: partition}
			consumer := cg.Consumers[topicPartition]
			delete(cg.Consumers, topicPartition)
			close(consumer.quit)
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-consumer.done
			}()
		}

	}
}

// pollTopics polls for records and sends them to the appropriate consumer.
func (cg *ConsumerGroup) pollTopics(cl *kgo.Client) {
	for {
		fetches := cl.PollRecords(context.Background(), 1000)
		if fetches.IsClientClosed() {
			return
		}

		fetches.EachError(func(s string, i int32, err error) {
			panic(err)
		})

		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			topicPartition := topicPartition{topic: p.Topic, partition: p.Partition}
			cg.Consumers[topicPartition].recs <- p.Records
		})
	}
}

type topicPartition struct {
	topic     string
	partition int32
}

type consumer struct {
	quit chan struct{}
	done chan struct{}
	recs chan []*kgo.Record

	workerFunc WorkerFunc
}

func (c *consumer) Consume() {
	defer close(c.done)
	for {
		select {
		case <-c.quit:
			return
		case recs := <-c.recs:
			for _, rec := range recs {
				err := c.workerFunc(rec)
				if err != nil {
					//pc.cl.CommitRecords(context.Background(), rec)
				}

			}
		}

	}
}
