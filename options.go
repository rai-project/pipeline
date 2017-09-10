package pipeline

import "golang.org/x/net/context"

type Options struct {
	channelBuffer int
	ctx           context.Context
}

type Option func(*Options) *Options

func Context(ctx context.Context) Option {
	return func(o *Options) *Options {
		o.ctx = ctx
		return o
	}
}

func ChannelBuffer(c int) Option {
	return func(o *Options) *Options {
		o.channelBuffer = c
		return o
	}
}

func NewOptions() *Options {
	return &Options{
		channelBuffer: 1,
		ctx:           context.Background(),
	}
}
