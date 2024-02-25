package api

import "context"

type Client interface {
	Execute(ctx context.Context, execute string, arguments, result any) error
	RegisterEvent(name string, factory func() EventType)
}

type EventType interface {
	Event() string
}
