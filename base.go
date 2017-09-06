package pipeline

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
)

type BaseStep struct{}

func (p BaseStep) do(ctx context.Context, in0 interface{}) interface{} {
	return errors.New("the base step is not implemented")
}

func (p BaseStep) Run(ctx context.Context, in <-chan interface{}, out chan interface{}) {
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
				out <- p.do(ctx, input)
			}
		}
	}()
}
