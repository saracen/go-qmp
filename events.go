package qmp

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/saracen/go-qmp/api"
)

var ErrListenerClosed = errors.New("listener closed")

type Event struct {
	Name      string
	Event     api.EventType
	Timestamp time.Time
}

type event struct {
	Event     string          `json:"event"`
	Data      json.RawMessage `json:"data"`
	Timestamp struct {
		Second       int64 `json:"seconds"`
		Microseconds int64 `json:"microseconds"`
	} `json:"timestamp"`
}

type eventHandler struct {
	logger *slog.Logger

	mu          sync.Mutex
	registered  map[string]func() api.EventType
	subscribers map[*subscriber]func(Event) bool
}

func newEventHandler(logger *slog.Logger) *eventHandler {
	return &eventHandler{
		logger:      logger,
		registered:  make(map[string]func() api.EventType),
		subscribers: make(map[*subscriber]func(Event) bool),
	}
}

type subscriber struct {
	items []Event
	mu    sync.Mutex
	cond  *sync.Cond
	close func()
}

func (h *eventHandler) register(name string, factory func() api.EventType) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.registered[name] = factory
}

func (h *eventHandler) new(ctx context.Context, filter func(Event) bool) <-chan Event {
	ch := make(chan Event)

	ctx, cancel := context.WithCancel(ctx)

	sub := &subscriber{close: cancel}
	sub.cond = sync.NewCond(&sub.mu)

	h.mu.Lock()
	h.subscribers[sub] = filter
	h.mu.Unlock()

	go func() {
		<-ctx.Done()
		sub.cond.Signal()
	}()

	go func() {
		defer func() {
			close(ch)

			h.mu.Lock()
			delete(h.subscribers, sub)
			h.mu.Unlock()
		}()

		defer func() {
			sub.mu.Lock()
			sub.items = nil
			sub.mu.Unlock()
		}()

		for {
			sub.mu.Lock()

			if len(sub.items) == 0 {
				sub.cond.Wait()
				if ctx.Err() != nil {
					sub.mu.Unlock()
					return
				}
			}

			select {
			case <-ctx.Done():
				sub.mu.Unlock()
				return

			case ch <- sub.items[0]:
				sub.items = sub.items[1:]
			default:
			}

			sub.mu.Unlock()
		}
	}()

	return ch
}

func (h *eventHandler) handle(e event) {
	h.mu.Lock()
	defer h.mu.Unlock()

	factory := h.registered[e.Event]
	if factory == nil {
		h.logger.Warn("no registered event factor", "event", e.Event)
		return
	}

	data := factory()
	if err := json.Unmarshal(e.Data, &data); err != nil {
		h.logger.Warn("unmarshaling event", "event", e.Event)
		return
	}

	event := Event{
		Name:  e.Event,
		Event: data,
		Timestamp: time.Unix(0,
			(e.Timestamp.Second*int64(time.Second))+
				(e.Timestamp.Microseconds*int64(time.Microsecond)),
		),
	}

	for sub, filter := range h.subscribers {
		if filter(event) {
			sub.mu.Lock()
			sub.items = append(sub.items, event)
			sub.cond.Signal()
			sub.mu.Unlock()
		}
	}
}

func (h *eventHandler) clear() {
	h.mu.Lock()
	subs := make([]*subscriber, 0, len(h.subscribers))
	for sub := range h.subscribers {
		subs = append(subs, sub)
	}
	h.mu.Unlock()

	for _, sub := range subs {
		sub.close()
	}
}
