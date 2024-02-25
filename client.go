package qmp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"

	"github.com/saracen/go-qmp/api"
)

var ErrConnectionClosed = errors.New("connection closed")

type Client struct {
	logger *slog.Logger

	conn   net.Conn
	writer *json.Encoder
	reader *json.Decoder

	greeting Greeting

	execId atomic.Int64

	mu  sync.Mutex
	err error

	messages *messageHandler
	events   *eventHandler
}

type request struct {
	ID        int64  `json:"id"`
	Execute   string `json:"execute"`
	Arguments any    `json:"arguments,omitempty"`
}

type response struct {
	ID      int64 `json:"id"`
	Message `json:",inline"`
	event   `json:",inline"`
}

func Dial(ctx context.Context, network, address string, opts ...Option) (*Client, error) {
	var options Options

	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}

	if options.Dialer == nil {
		options.Dialer = &net.Dialer{}
	}
	if options.Logger == nil {
		options.Logger = slog.Default()
	}

	conn, err := options.Dialer.DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}

	c := &Client{
		logger:   options.Logger,
		conn:     conn,
		writer:   json.NewEncoder(conn),
		reader:   json.NewDecoder(conn),
		messages: newMessageHandler(),
		events:   newEventHandler(options.Logger),
	}
	c.reader.DisallowUnknownFields()

	if err := c.negotiate(); err != nil {
		c.conn.Close()
		return nil, fmt.Errorf("negotiate: %w", err)
	}

	go c.process()

	return c, nil
}

func (c *Client) Close() {
	c.conn.Close()

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err != nil {
		c.err = ErrConnectionClosed
	}
}

func (c *Client) Listen(ctx context.Context, filter func(Event) bool) <-chan Event {
	return c.events.new(ctx, filter)
}

func (c *Client) RegisterEvent(name string, factory func() api.EventType) {
	c.events.register(name, factory)
}

func (c *Client) Execute(ctx context.Context, execute string, arguments, result any) error {
	req := request{
		ID:        c.execId.Add(1),
		Execute:   execute,
		Arguments: arguments,
	}

	reply, err := c.issue(req)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()

	case msg, ok := <-reply:
		if !ok {
			// we only ever expect 1 message, so if this is closed,
			// it's because we're shutting down.
			return ErrConnectionClosed
		}

		if msg.Error != nil {
			return msg.Error
		}

		if result == nil {
			return nil
		}

		err := json.Unmarshal(msg.Return, result)
		if err != nil {
			return fmt.Errorf("unmarshaling return: %w", err)
		}
	}

	return nil
}

func (c *Client) issue(req request) (chan Message, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err != nil {
		return nil, c.err
	}

	err := c.writer.Encode(req)
	if err != nil {
		return nil, err
	}

	return c.messages.register(req.ID), c.err
}

func (c *Client) process() {
	defer func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.messages.clear()
		c.events.clear()
	}()

	for {
		var resp response

		if c.shouldStop(c.reader.Decode(&resp)) {
			return
		}

		// handle events
		if resp.event.Event != "" {
			c.events.handle(resp.event)
			continue
		}

		if c.messages.handle(resp) {
			continue
		}

		if resp.ID == 0 && resp.Error != nil {
			msg := "server sent an error response without an ID, expected server parser failure: %v"
			if c.shouldStop(fmt.Errorf(msg, resp.Message)) {
				return
			}
		}

		if resp.Error != nil {
			c.logger.Error("message dropped", "id", resp.ID, "err", resp.Error)
		} else {
			c.logger.Warn("message dropped", "id", resp.ID)
		}

		encoded, _ := json.Marshal(resp)
		c.logger.Debug("unroutable message", "msg", string(encoded))
	}
}

func (c *Client) shouldStop(err error) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err != nil {
		return true
	}

	if err == nil {
		return false
	}
	c.err = err

	return true
}
