package pipeline

import (
	"context"
)

type Options struct {
	channelBuffer int
	ctx           context.Context
}

type Option func(*Options)

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func (o *Options) Context() context.Context {
	return o.ctx
}

func ChannelBuffer(c int) Option {
	return func(o *Options) {
		o.channelBuffer = c
	}
}

func (o *Options) ChannelBuffer() int {
	return o.channelBuffer
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		channelBuffer: 1,
		ctx:           context.Background(),
	}
	for _, o := range opts {
		o(options)
	}
	return options
}
