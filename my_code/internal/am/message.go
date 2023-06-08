package am

import (
	"context"
	"muazwzxv/Micheal-stack-book/internal/ddd"
)

type (
	Message interface {
		ddd.IDer
		MessageName() string
		Ack() error
		NAck() error
		Extend() error
		Kill() error
	}

	MessageHandler[K Message] interface {
		HandleMessage(ctx context.Context, msg K) error
	}

	MessageHandlerFunc[K Message] func(ctx context.Context, msg K) error

	MessagePublisher[V any] interface {
		Publish(ctx context.Context, topicName string, v V) error
	}

	MessageSubscriber[K Message] interface {
		Subscribe(topicName string, handler MessageHandler[K], options ...SubscriberOption) error
	}

	MessageStream[V any, K Message] interface {
		MessagePublisher[V]
		MessageSubscriber[K]
	}
)

func (f MessageHandlerFunc[K]) HandleMessage(ctx context.Context, msg K) error {
  return f(ctx, msg)
}
