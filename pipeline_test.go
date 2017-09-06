package pipeline

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	ctx := context.Background()

	pipe := New(ctx).Then(
		StepFunction(func(ctx context.Context, in0 interface{}) interface{} {
			in, ok := in0.(int)
			if !ok {
				return errors.Errorf("invalid value %v", in0)
			}
			return in + 1
		}),
	).Then(
		StepFunction(func(ctx context.Context, in0 interface{}) interface{} {
			in, ok := in0.(int)
			if !ok {
				return errors.Errorf("invalid value %v", in0)
			}
			return in * 4
		}),
	)

	input := make(chan interface{})
	go func() {
		for ii := 0; ii < 3; ii++ {
			input <- ii
		}
		close(input)
	}()

	outputs := pipe.Run(ctx, input)

	ii := 0
	for output := range outputs {
		assert.NotEmpty(t, output)
		assert.Equal(t, output, (ii+1)*4)
		ii++
	}
}
