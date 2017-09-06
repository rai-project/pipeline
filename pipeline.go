package pipeline

import "golang.org/x/net/context"

type pipeline struct {
	steps []Step
	ctx   context.Context
}

func New(ctx context.Context) *pipeline {
	return &pipeline{
		ctx: ctx,
	}
}

func (p *pipeline) Then(step Step) *pipeline {
	p.steps = append(p.steps, step)
	return p
}

func (p *pipeline) Step(ctx context.Context, s Step, in <-chan interface{}, out chan interface{}) {
	s.Run(ctx, in, out)
}

func (p *pipeline) Run(ctx context.Context, in <-chan interface{}) chan interface{} {
	var out chan interface{}
	for _, step := range p.steps {
		out = make(chan interface{})
		step.Run(ctx, in, out)
		in = out
	}
	return out
}
