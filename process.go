package flow

import (
	"context"
	"io"
)

type Process interface {
	New(ctx context.Context) (Process, error)
	Run(ctx context.Context) error
	io.Closer
}

type ProcessFunction func(ctx context.Context) error

func (p ProcessFunction) New(ctx context.Context) (Process, error) {
	return p, nil
}

func (p ProcessFunction) Run(ctx context.Context) error {
	return p(ctx)
}

func (p ProcessFunction) Close() error {
	return nil
}
