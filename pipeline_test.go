package pipeline

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {

	pipe := New().Then(
		StepFunction(func(ctx context.Context, in0 interface{}, _ *Options) interface{} {
			in, ok := in0.(int)
			if !ok {
				return errors.Errorf("invalid value %v", in0)
			}
			return in + 1
		}),
	).Then(
		StepFunction(func(ctx context.Context, in0 interface{}, _ *Options) interface{} {
			in, ok := in0.(int)
			if !ok {
				return errors.Errorf("invalid value %v", in0)
			}
			return in * 4
		}),
	)

	input := make(chan interface{})
	go func() {
		defer close(input)
		for ii := 0; ii < 10; ii++ {
			input <- ii
		}
	}()

	outputs := pipe.Run(input)

	ii := 0
	for output := range outputs {
		assert.NotEmpty(t, output)
		assert.Equal(t, output, (ii+1)*4)
		ii++
	}
}
