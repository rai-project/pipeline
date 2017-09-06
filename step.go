package pipeline

import (
	"context"
	"io"
)

type Step interface {
	New(ctx context.Context) (Step, error)
	Run(ctx context.Context, in <-chan interface{}) chan interface{}
	io.Closer
}

type StepFunction func(ctx context.Context, in interface{}) interface{}

func (p StepFunction) New(ctx context.Context) (Step, error) {
	return p, nil
}

func (p StepFunction) Run(ctx context.Context, in <-chan interface{}) chan interface{} {
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

func (p StepFunction) Close() error {
	return nil
}
