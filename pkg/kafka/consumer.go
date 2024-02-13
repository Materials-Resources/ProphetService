package kafka

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
)

type consumer struct {
	tracer     *kotel.Tracer
	consumeMsg func(ctx context.Context, rec *kgo.Record) error
	quit       chan struct{}
	recs       chan []*kgo.Record
}

func (c *consumer) consume(topic string, partition int32) {
	fmt.Printf("starting, t %s p %d\n", topic, partition)
	defer fmt.Printf("killing, t %s p %d\n", topic, partition)
	var (
		nrecs int
	)
	for {
		select {
		case <-c.quit:
			return
		case recs := <-c.recs:
			nrecs += len(recs)
			for _, rec := range recs {
				c.consumeRec(rec)

			}
		}
	}
}

func (c *consumer) consumeRec(rec *kgo.Record) {
	ctx, span := c.tracer.WithProcessSpan(rec)
	defer span.End()
	err := c.consumeMsg(ctx, rec)
	if err != nil {
		span.RecordError(err)
	}
}
