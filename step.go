package pipeline

import (
	"context"
	"io"
)

type Step interface {
	Info() string
	New(ctx context.Context) (Step, error)
	Run(ctx context.Context, in <-chan interface{}, out chan interface{})
	io.Closer
}

type StepFunction func(ctx context.Context, in interface{}) interface{}

func (p StepFunction) New(ctx context.Context) (Step, error) {
	return p, nil
}

func (p StepFunction) Info() string {
	return "StepFunction"
}

func (p StepFunction) Run(ctx context.Context, in <-chan interface{}, out chan interface{}) {
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
				out <- p(ctx, input)
			}
		}
	}()
}

func (p StepFunction) Close() error {
	return nil
}
