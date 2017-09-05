package ops

import (
	"context"

	"github.com/rai-project/flow"
)

type urlReader struct {
}

func (p urlReader) New(ctx context.Context) (flow.Process, error) {
	return p, nil
}

func (p urlReader) Run(ctx context.Context) error {
	return nil
}

func (p urlReader) Close() error {
	return nil
}

func init() {
	flow.Register("URLReader", urlReader{})
}
