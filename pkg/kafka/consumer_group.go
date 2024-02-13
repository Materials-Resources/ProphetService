package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
)

type ConsumeMsg func(ctx context.Context, rec *kgo.Record) error

func NewConsumerGroup(consumeMsg ConsumeMsg, tracer *kotel.Tracer) *ConsumerGroup {
	return &ConsumerGroup{consumers: make(map[string]map[int32]consumer), consumeMsg: consumeMsg, tracer: tracer}
}

type ConsumerGroup struct {
	tracer     *kotel.Tracer
	consumeMsg ConsumeMsg
	mu         sync.Mutex // gaurds assigning / losing vs. polling
	consumers  map[string]map[int32]consumer
}

// Start starts the consumer group in a new go routine
func (g *ConsumerGroup) Start(brokers []string, topic, group string, kotel *kotel.Kotel) {

	cl, err := kgo.NewClient(kgo.SeedBrokers(brokers...), kgo.ConsumeTopics(topic), kgo.ConsumerGroup(group),
		kgo.OnPartitionsAssigned(g.assigned), kgo.OnPartitionsRevoked(g.lost), kgo.OnPartitionsRevoked(g.lost),
		kgo.WithHooks(kotel.Hooks()...),
	)
	if err != nil {
		fmt.Println(err)
	}
	go g.poll(cl)

}

func (g *ConsumerGroup) assigned(_ context.Context, cl *kgo.Client, assigned map[string][]int32) {
	g.mu.Lock()
	defer g.mu.Unlock()
	for topic, partitions := range assigned {
		if g.consumers[topic] == nil {
			g.consumers[topic] = make(map[int32]consumer)
		}
		for _, partition := range partitions {
			pc := consumer{
				tracer:     g.tracer,
				consumeMsg: g.consumeMsg,
				quit:       make(chan struct{}),
				recs:       make(chan []*kgo.Record, 10),
			}
			g.consumers[topic][partition] = pc
			go pc.consume(topic, partition)
		}
	}
}

func (g *ConsumerGroup) lost(_ context.Context, cl *kgo.Client, lost map[string][]int32) {
	g.mu.Lock()
	defer g.mu.Unlock()
	for topic, partitions := range lost {
		ptopics := g.consumers[topic]
		for _, partition := range partitions {
			pc := ptopics[partition]
			delete(ptopics, partition)
			if len(ptopics) == 0 {
				delete(g.consumers, topic)
			}
			close(pc.quit)
		}
	}
}

func (g *ConsumerGroup) poll(cl *kgo.Client) {
	for {
		fetches := cl.PollFetches(context.Background())
		if fetches.IsClientClosed() {
			return
		}
		fetches.EachError(func(_ string, _ int32, err error) {
			// Note: you can delete this block, which will result
			// in these errors being sent to the partition
			// consumers, and then you can handle the errors there.
			panic(err)
		},
		)
		fetches.EachTopic(func(t kgo.FetchTopic) {
			g.mu.Lock()
			tconsumers := g.consumers[t.Topic]
			g.mu.Unlock()
			if tconsumers == nil {
				return
			}
			t.EachPartition(func(p kgo.FetchPartition) {
				pc, ok := tconsumers[p.Partition]
				if !ok {
					return
				}
				select {
				case pc.recs <- p.Records:
				case <-pc.quit:
				}
			},
			)
		},
		)
	}
}
