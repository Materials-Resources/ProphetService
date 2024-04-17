package consumer

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

type Worker interface {
	GetTopic() string
	ConsumeRecord(rec *kgo.Record) error
	SendToDLQ(cl *kgo.Client, rec *kgo.Record, err error) error
}

type ConsumeFunc func(rec *kgo.Record) error

type consumer struct {
	cl        *kgo.Client
	topic     string
	partition int32

	done       chan struct{}
	quit       chan struct{}
	consumeMsg ConsumeFunc
	sendToDLQ  func(cl *kgo.Client, rec *kgo.Record, err error) error
	recs       chan kgo.FetchTopicPartition
}

func (c *consumer) consume() {
	for {
		select {
		case <-c.quit:
			return
		case p := <-c.recs:
			for _, rec := range p.Records {
				err := c.consumeMsg(rec)
				if err != nil {
					c.sendToDLQ(c.cl, rec, err)
				}
				c.cl.MarkCommitRecords(rec)
			}
		}
	}
}
