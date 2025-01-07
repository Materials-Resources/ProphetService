package event

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

const (
	UpdateGroupTopic   = "catalog_group_update"
	CreateGroupTopic   = "catalog_group_create"
	UpdateProductTopic = "product_update"
	DeleteProductTopic = "product_delete"
)

// RegisterTopics registers the topics with the Kafka cluster
func (m *Manager) RegisterTopics() error {
	var adminClient *kadm.Client

	{
		client, err := kgo.NewClient(m.GetDefaultKgoOptions()...)
		if err != nil {
			return err
		}

		defer client.Close()

		adminClient = kadm.NewClient(client)

	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Check what topics are already registered
	registeredTopics, err := adminClient.ListTopics(ctx)
	if err != nil {
		return fmt.Errorf("could not retrieve topics: %v", err)
	}

	// Register the topics
	topics := []string{UpdateGroupTopic, CreateGroupTopic, UpdateProductTopic, DeleteProductTopic}
	for _, topic := range topics {
		if !registeredTopics.Has(topic) {
			_, err := adminClient.CreateTopics(ctx, -1, -1, nil, topic)
			if err != nil {
				return fmt.Errorf("could not create topic %s: %v", topic, err)
			}
		}
	}

	return nil

}
