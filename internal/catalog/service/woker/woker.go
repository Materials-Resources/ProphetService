package woker

import "github.com/twmb/franz-go/pkg/kgo"

type Worker interface {
	ConsumeUpdateProduct(rec *kgo.Record) error
}
