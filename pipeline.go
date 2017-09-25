package pipeline

import "golang.org/x/net/context"

type pipeline struct {
	steps   []Step
	options *Options
}

func New(opts ...Option) *pipeline {
	options := NewOptions()
	for _, o := range opts {
		o(options)
	}
	return &pipeline{
		options: options,
	}
}

func (p *pipeline) Then(step Step) *pipeline {
	p.steps = append(p.steps, step)
	return p
}

func (p *pipeline) Step(ctx context.Context, s Step, in <-chan interface{}, out chan interface{}) {
	s.Run(ctx, in, out)
}

func (p *pipeline) Run(in <-chan interface{}, opts ...Option) <-chan interface{} {
	var out chan interface{}

	ctx := p.options.ctx
	channelBuffer := p.options.channelBuffer
	for _, step := range p.steps {
		out = make(chan interface{}, channelBuffer)
		step.Run(ctx, in, out, opts...)
		in = out
	}
	return out
}

func (p *pipeline) Close() error {
	for _, step := range p.steps {
		step.Close()
	}
	return nil
}
