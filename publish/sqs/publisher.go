package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"

	"github.com/xabi93/messenger/publish"
	"github.com/xabi93/messenger/store"
)

//go:generate moq -pkg sqs_test -stub -out publisher_mock_test.go . SQSClient

// SQSClient defines the AWS SQS methods used by the Publisher. This is used for testing pourpouses.
type SQSClient interface {
	GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

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

// WithFifoQueue setups the flag to use fifo queue.
func WithFifoQueue(fifo bool) Option {
	return func(p *Publisher) {
		p.fifo = fifo
	}
}

// Open returns a new Publisher instance.
func Open(ctx context.Context, awsOpts sqs.Options, queue string, opts ...Option) (*Publisher, error) {
	return New(ctx, sqs.New(awsOpts), queue, opts...)
}

// New returns a new Publisher instance.
func New(ctx context.Context, svc SQSClient, queue string, opts ...Option) (*Publisher, error) {
	q, err := svc.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{QueueName: aws.String(queue)})
	if err != nil {
		return nil, fmt.Errorf("getting queue url: %w", err)
	}

	p := Publisher{
		svc:   svc,
		queue: aws.ToString(q.QueueUrl),
	}

	for _, opt := range opts {
		opt(&p)
	}

	return &p, nil
}

var _ publish.Queue = &Publisher{}

// Publisher handles the pubsub topic messages.
type Publisher struct {
	// sqs service instance where are going to publish messages
	svc SQSClient
	// queue url where are going to publish messages
	queue string
	// meta property of the message to use as ordering key
	metaOrdKey string
	// default ordering key in case not provided in message metadata
	defaultOrdKey string
	// flag to use fifo queue
	fifo bool
}

// Publish publishes the given message to the pubsub topic.
func (p Publisher) Publish(ctx context.Context, msg *store.Message) error {
	att := make(map[string]types.MessageAttributeValue, len(msg.Metadata))
	for k, v := range msg.Metadata {
		att[k] = types.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(v),
		}
	}

	_, err := p.svc.SendMessage(
		ctx,
		&sqs.SendMessageInput{
			MessageDeduplicationId: p.messageDeduplication(msg),
			MessageAttributes:      att,
			MessageBody:            aws.String(string(msg.Payload)),
			QueueUrl:               aws.String(p.queue),
			MessageGroupId:         p.orderingKey(msg),
		})
	if err != nil {
		return fmt.Errorf("publishing message: %w", err)
	}

	return nil
}

// messageDeduplication checks if the publisher is setup as fifo and returns the message deduplication id.
func (p Publisher) messageDeduplication(msg *store.Message) *string {
	if !p.fifo {
		return nil
	}

	return aws.String(msg.ID)
}

// orderingKey tries to get the ordering key from message metadata
// in case the message does not have the key it defaults to Publisher setup.
func (p Publisher) orderingKey(msg *store.Message) *string {
	if !p.fifo {
		return nil
	}

	key, ok := msg.Metadata[p.metaOrdKey]
	if ok {
		return aws.String(key)
	}

	if p.defaultOrdKey != "" {
		return aws.String(p.defaultOrdKey)
	}

	return nil
}
