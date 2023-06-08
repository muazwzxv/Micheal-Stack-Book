package ddd

import (
	"time"

	"github.com/google/uuid"
)

type (
	EventPayload interface{}

  EventOption interface {
  configureEvent(*event)
}

	Event interface {
		IDer
		EventName() string
		Payload() EventPayload
		Metadata() Metadata
		OccuredAt() time.Time
	}

  event struct {
    Entity
    paylaod EventPayload
    metadata Metadata
    occuredAt time.Time
  }
)

var _ Event = (*event)(nil)

func NewEvent(
  name string,
  payload EventPayload,
  options ...EventOption,
) event {
  return newEvent(name, payload, options...)
}

func newEvent(
  name string,
  payload EventPayload,
  options ...EventOption,
) event {
  evt := event{
    Entity: NewEntity(uuid.NewString(), name),
    paylaod: payload,
    metadata: make(Metadata),
    occuredAt: time.Now(),
  }

  for _, option := range options {
    option.configureEvent(&evt)
  }

  return evt
}

func (e event) EventName() string {
  return e.name
}

func (e event) Payload() EventPayload {
  return e.paylaod
}

func (e event) Metadata() Metadata {
  return e.metadata
}

func (e event) OccuredAt() time.Time {
  return e.occuredAt
}

