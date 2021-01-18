package inputs

import (
	"context"
	"log"

	"github.com/karimra/gnmic/outputs"
)

type Input interface {
	Start(context.Context, map[string]interface{}, ...Option) error
	Close() error
	SetLogger(*log.Logger)
	SetOutputs(map[string]outputs.Output)
}

type Initializer func() Input

var InputTypes = []string{
	"nats",
	"stan",
}

var Inputs = map[string]Initializer{}

func Register(name string, initFn Initializer) {
	Inputs[name] = initFn
}

type Option func(Input)

func WithLogger(logger *log.Logger) Option {
	return func(i Input) {
		i.SetLogger(logger)
	}
}

func WithOutputs(outs map[string]outputs.Output) Option {
	return func(i Input) {
		i.SetOutputs(outs)
	}
}
