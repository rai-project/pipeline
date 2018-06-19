package pipeline

import (
	"io"

	"context"
)

type Step interface {
	Info() string
	Run(ctx context.Context, in <-chan interface{}, out chan interface{}, opts ...Option)
	io.Closer
}

type StepFunction func(ctx context.Context, in interface{}, opts *Options) interface{}

func (p StepFunction) Info() string {
	return "StepFunction"
}

func (p StepFunction) Run(ctx context.Context, in <-chan interface{}, out chan interface{}, opts ...Option) {
	options := NewOptions(opts...)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				// if err := ctx.Err(); err != nil {
				// 	out <- err
				// }
				return
			case input, open := <-in:
				if !open {
					return
				}
				if err, ok := input.(error); ok {
					out <- err
					continue
				}
				out <- p(ctx, input, options)
			}
		}
	}()
}

func (p StepFunction) Close() error {
	return nil
}
