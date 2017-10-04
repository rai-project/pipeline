package pipeline

import (
	"golang.org/x/net/context"
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

func ChannelBuffer(c int) Option {
	return func(o *Options) {
		o.channelBuffer = c
	}
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
