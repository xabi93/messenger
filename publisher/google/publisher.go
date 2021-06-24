package google

import (
	"context"

	"cloud.google.com/go/pubsub"

	"github.com/xabi93/messenger"
	"github.com/xabi93/messenger/publisher"
)

// Option is a function to set options to Publisher.
type Option func(*Publisher)

// WithMetaOrderingKey setups the metadata key to get the ordering key.
func WithMetaOrderingKey(key string) Option {
	return func(p *Publisher) {
		p.metaOrdKey = key
	}
}

// WithDefaultOrderingKey setups the default ordering key.
func WithDefaultOrderingKey(key string) Option {
	return func(p *Publisher) {
		p.defaultOrdKey = key
	}
}

// New returns a new Publisher instance.
func New(topic *pubsub.Topic, opts ...Option) *Publisher {
	p := Publisher{topic: topic}

	for _, opt := range opts {
		opt(&p)
	}

	return &p
}

var _ publisher.Publisher = &Publisher{}

// Publisher handles the pubsub topic messages.
type Publisher struct {
	// pubsub topic instance where are going to publish messages
	topic *pubsub.Topic
	// meta property of the message to use as ordering key
	metaOrdKey string
	// default ordering key in case not provided in message metadata
	defaultOrdKey string
}

// Publish publishes the given message to the pubsub topic.
func (p Publisher) Publish(ctx context.Context, msg *messenger.Message) error {
	_, err := p.topic.Publish(ctx, &pubsub.Message{
		Attributes:  msg.Metadata,
		Data:        msg.Payload,
		OrderingKey: p.orderingKey(msg),
	}).Get(ctx)

	return err
}

// orderingKey tries to get the ordering key from message metadata
// in case the message does not have the key it defaults to Publisher setup.
func (p Publisher) orderingKey(msg *messenger.Message) string {
	if key := msg.Metadata.Get(p.metaOrdKey); key != "" {
		return key
	}

	return p.defaultOrdKey
}