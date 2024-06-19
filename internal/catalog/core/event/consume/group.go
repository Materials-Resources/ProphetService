package consume

import (
	"context"
	"github.com/twmb/franz-go/pkg/kgo"
	"sync"
)

type topicPartition struct {
	topic     string
	partition int32
}

type ConsumerGroup struct {
	workerFunctions map[string]WorkerFunc
	consumers       map[topicPartition]*PartitionConsumer
}

func NewConsumerGroup() *ConsumerGroup {
	return &ConsumerGroup{
		workerFunctions: make(map[string]WorkerFunc),
		consumers:       make(map[topicPartition]*PartitionConsumer),
	}
}

func (cg *ConsumerGroup) AddWorker(topic string, consumeFunc WorkerFunc) {
	cg.workerFunctions[topic] = consumeFunc
}

// assigned is called when the consumer group is assigned partitions.
func (cg *ConsumerGroup) assigned(_ context.Context, cl *kgo.Client, assigned map[string][]int32) {
	for topic, partitions := range assigned {
		for _, partition := range partitions {
			topicPartition := topicPartition{topic: topic, partition: partition}
			consumer := &PartitionConsumer{
				cl:        cl,
				topic:     topic,
				partition: partition,
				quit:      make(chan struct{}),
				done:      make(chan struct{}),
				recs:      make(chan []*kgo.Record, 5),
				wokerFunc: cg.workerFunctions[topic],
			}
			cg.consumers[topicPartition] = consumer
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
			consumer := cg.consumers[topicPartition]
			delete(cg.consumers, topicPartition)
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
			cg.consumers[topicPartition].recs <- p.Records
		})
	}
}

type PartitionConsumer struct {
	cl        *kgo.Client
	topic     string
	partition int32

	quit chan struct{}
	done chan struct{}
	recs chan []*kgo.Record

	wokerFunc WorkerFunc
}

func (pc *PartitionConsumer) Consume() {
	defer close(pc.done)
	for {
		select {
		case <-pc.quit:
			return
		case recs := <-pc.recs:
			for _, rec := range recs {
				err := pc.wokerFunc(rec)
				if err != nil {
					pc.cl.CommitRecords(context.Background(), rec)
				}
			}
		}

	}
}
