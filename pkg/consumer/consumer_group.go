package consumer

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"sync"
)

type tp struct {
	topic     string
	partition int32
}
type ConsumerGroup struct {
	client    kgo.Client
	topics    []string
	workers   map[string]Worker
	consumers map[tp]*consumer
}

func NewConsumerGroup() ConsumerGroup {
	return ConsumerGroup{
		workers:   make(map[string]Worker),
		consumers: make(map[tp]*consumer),
	}
}

// Start starts the consumer group in a new go routine
func (g *ConsumerGroup) Start(brokers []string, group string) {

	opts := []kgo.Opt{
		kgo.SeedBrokers(brokers...), kgo.ConsumeTopics(g.topics...), kgo.ConsumerGroup(group),
		kgo.OnPartitionsAssigned(g.assigned),
		kgo.OnPartitionsRevoked(g.revoked),
		kgo.OnPartitionsLost(g.lost),
		kgo.AutoCommitMarks(),
		kgo.BlockRebalanceOnPoll(),
	}
	cl, err := kgo.NewClient(opts...)
	if err != nil {
		fmt.Println(err)
	}
	go g.poll(cl)

}

func (g *ConsumerGroup) assigned(_ context.Context, cl *kgo.Client, assigned map[string][]int32) {
	for topic, partitions := range assigned {
		for _, partition := range partitions {
			pc := &consumer{
				cl:         cl,
				topic:      topic,
				partition:  partition,
				consumeMsg: g.workers[topic].consumeFunc,
				recs:       make(chan kgo.FetchTopicPartition, 5),
			}
			g.consumers[tp{topic, partition}] = pc
			go pc.consume()
		}
	}
}
func (g *ConsumerGroup) revoked(ctx context.Context, cl *kgo.Client, revoked map[string][]int32) {
	g.killConsumers(revoked)
	if err := cl.CommitMarkedOffsets(ctx); err != nil {
		fmt.Printf("Revoke commit failed: %v\n", err)
	}
}

func (g *ConsumerGroup) lost(_ context.Context, cl *kgo.Client, lost map[string][]int32) {
	g.killConsumers(lost)
}

func (g *ConsumerGroup) killConsumers(lost map[string][]int32) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for topic, partitions := range lost {
		for _, partition := range partitions {
			tp := tp{topic: topic, partition: partition}
			pc := g.consumers[tp]
			delete(g.consumers, tp)
			close(pc.quit)
			fmt.Printf("waiting for work to finish  t %s p %d\n", topic, partition)
			wg.Add(1)
			go func() { <-pc.done; wg.Done() }()
		}
	}
}

func (g *ConsumerGroup) poll(cl *kgo.Client) {
	for {
		fetches := cl.PollRecords(context.Background(), 500)
		if fetches.IsClientClosed() {
			return
		}
		fetches.EachError(
			func(_ string, _ int32, err error) {
				// Note: you can delete this block, which will result
				// in these errors being sent to the partition
				// consumers, and then you can handle the errors there.
				panic(err)
			},
		)
		fetches.EachPartition(
			func(p kgo.FetchTopicPartition) {
				tp := tp{p.Topic, p.Partition}

				g.consumers[tp].recs <- p

			})
		cl.AllowRebalance()
	}

}

func NewWorker(topic string, consumerFunc ConsumeFunc) Worker {
	return Worker{topic: topic, consumeFunc: consumerFunc}
}

func (g *ConsumerGroup) RegisterWorkers(workers ...Worker) {

	for _, worker := range workers {
		g.topics = append(g.topics, worker.topic)
		g.workers[worker.topic] = worker
	}
}
