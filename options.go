package qmp

import (
	"context"
	"log/slog"
	"net"
)

type Dialer interface {
	DialContext(ctx context.Context, network string, addr string) (net.Conn, error)
}

type Options struct {
	Dialer Dialer
	Logger *slog.Logger
}

type Option func(*Options) error

func WithDialer(dialer Dialer) Option {
	return func(o *Options) error {
		o.Dialer = dialer
		return nil
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(o *Options) error {
		o.Logger = logger
		return nil
	}
}

func ToPtr[T any](v T) *T {
	return &v
}

func FromPtr[T any](v *T) T {
	if v == nil {
		var def T
		return def
	}
	return *v
}
