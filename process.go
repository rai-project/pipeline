package flow

import (
	"context"
	"io"
)

type Process interface {
	Run(ctx context.Context) error
	io.Closer
}
