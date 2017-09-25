package pipeline

import (
	"github.com/rai-project/tracer"
	"golang.org/x/net/context"
)

type Options struct {
	channelBuffer int
	Tracer        tracer.Tracer
	ctx           context.Context
}

type Option func(*Options)

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func ChannelBuffer(c int) Option {
	return func(o *Options) {
		o.channelBuffer = c
	}
}

func Tracer(tr tracer.Tracer) Option {
	return func(o *Options) {
		o.Tracer = tr
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		channelBuffer: 1,
		ctx:           context.Background(),
		Tracer:        tracer.Std(),
	}
	for _, o := range opts {
		o(options)
	}
	return options
}
