package flow

import (
	"context"
	"io"
)

type Process interface {
	New(ctx context.Context) (Process, error)
	Run(ctx context.Context, in <-chan interface{}) chan interface{}
	io.Closer
}

type ProcessFunction func(ctx context.Context, in interface{}) interface{}

func (p ProcessFunction) New(ctx context.Context) (Process, error) {
	return p, nil
}

func (p ProcessFunction) Run(ctx context.Context, in <-chan interface{}) chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case input := <-in:
				out <- p(ctx, input)
			}
		}
	}()
	return out
}

func (p ProcessFunction) Close() error {
	return nil
}
