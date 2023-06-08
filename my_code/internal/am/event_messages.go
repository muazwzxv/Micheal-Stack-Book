package am

import (
	"context"
	"muazwzxv/Micheal-stack-book/internal/ddd"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	EventMessage interface {
		Message
		ddd.Event
	}

	EventPublisher  = MessagePublisher[ddd.Event]
	EventSubScriber = MessageSubscriber[EventMessage]
	EventStream     = MessageStream[ddd.Event, EventMessage]

	eventStream struct {
		// registry
		stream MessageStream[RawMessage, RawMessage]
	}

	eventMessage struct {
		id        string
		name      string
		payload   ddd.EventPayload
		metadata  ddd.Metadata
		occuredAt time.Time
		msg       RawMessage
	}
)

var (
	_ EventMessage = (*eventMessage)(nil)
	_ EventStream  = (*eventStream)(nil)
)

func NewEventStream(
	// registry
	stream MessageStream[RawMessage, RawMessage],
) EventStream {
	return &eventStream{
		stream: stream,
	}
}

func (s eventStream) Publish(ctx context.Context, topicName string, event ddd.Event) error {
	metadata, err := structpb.NewStruct(event.Metadata())
	if err != nil {
		return err
	}

	// TODO: serdes for json and protobuf

	data, err := proto.Marshal(&EventMessageData{
		Payload:   nil, // TODO: payload from serdes
		OccuredAt: timestamppb.New(event.OccuredAt()),
		Metadata:  metadata,
	})
  if err != nil {
    return err
  }

	return s.stream.Publish(ctx, topicName, rawMessage{
		id:   event.ID(),
		name: event.EventName(),
		data: data,
	})
}

func (s eventStream) Subscribe(
	topicName string,
	handler MessageHandler[EventMessage],
	options ...SubscriberOption,
) error {
	cfg := NewSubscriberConfig(options)

	var filters map[string]struct{}
	if len(cfg.MessageFilter()) > 0 {
		filters = make(map[string]struct{})
		for _, key := range cfg.MessageFilter() {
			filters[key] = struct{}{}
		}
	}

	fn := MessageHandlerFunc[RawMessage](func(ctx context.Context, msg RawMessage) error {
		var eventData EventMessageData // protobuf type

		if filters != nil {
			if _, exists := filters[msg.MessageName()]; !exists {
				return nil
			}
		}

		err := proto.Unmarshal(msg.Data(), &eventData)
		if err != nil {
			return err
		}

		eventName := msg.MessageName()

		// TODO: serdes for JSON and Protobuf

		eventMsg := eventMessage{
			id:        msg.ID(),
			name:      eventName,
			payload:   nil, // TODO: payload from serdes
			metadata:  eventData.GetMetadata().AsMap(),
			occuredAt: eventData.GetOccuredAt().AsTime(),
			msg:       msg,
		}

		return handler.HandleMessage(ctx, eventMsg)
	})

	return s.stream.Subscribe(topicName, fn, options...)
}

func (e eventMessage) ID() string {
	return e.id
}

func (e eventMessage) EventName() string {
	return e.name
}

func (e eventMessage) Payload() ddd.EventPayload {
	return e.payload
}

func (e eventMessage) Metadata() ddd.Metadata {
	return e.metadata
}

func (e eventMessage) OccuredAt() time.Time {
	return e.occuredAt
}

func (e eventMessage) MessageName() string {
	return e.msg.MessageName()
}

func (e eventMessage) Ack() error {
	return e.msg.Ack()
}

func (e eventMessage) NAck() error {
	return e.msg.NAck()
}

func (e eventMessage) Extend() error {
	return e.msg.Extend()
}

func (e eventMessage) Kill() error {
	return e.msg.Kill()
}
