package generator

import (
	"github.com/ElfLabs/id-generator/pkg/format"
	"github.com/ElfLabs/id-generator/pkg/planner/snowflake"
	"github.com/ElfLabs/id-generator/pkg/sequencer"
)

type Options struct {
	Epoch  int64
	Count  int64
	Step   int64
	Size   uint
	Buffer bool

	StopCh <-chan struct{}
	ErrCh  chan<- error

	Storager
	Sequencer
	Formatter
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	var o Options
	o.Apply(opts)
	o.Init()
	return o
}

func (o *Options) Init() {
	switch {
	case o.Sequencer == nil && o.Formatter == nil:
		planner := snowflake.NewSnowflakePlanner()
		o.Epoch = planner.GetEpoch()
		o.Sequencer = planner
		o.Formatter = planner
	case o.Sequencer == nil:
		o.Epoch = sequencer.DefaultTimestampEpoch
		o.Sequencer = sequencer.NewTimestampSequencer()
	case o.Formatter == nil:
		o.Formatter = format.SnowflakeFormat{}
	}
}

func (o *Options) Apply(opts []Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o Options) GetEpoch() int64 {
	return o.Epoch
}

func (o Options) GetCount() int64 {
	return o.Count
}

func (o Options) GetStep() int64 {
	return o.Step
}

func WithEpoch(epoch int64) Option {
	return func(o *Options) {
		o.Epoch = epoch
	}
}

func WithCount(count int64) Option {
	return func(o *Options) {
		o.Count = count
	}
}

func WithStep(step int64) Option {
	return func(o *Options) {
		o.Step = step
	}
}

func WithRecover(count, step int64) Option {
	return func(o *Options) {
		o.Count = count
		o.Step = step
	}
}

func WithSequencer(sequencer Sequencer) Option {
	return func(o *Options) {
		o.Sequencer = sequencer
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *Options) {
		o.Formatter = formatter
	}
}

func WithPlanner(planner Planner) Option {
	return func(o *Options) {
		o.Epoch = planner.GetEpoch()
		o.Sequencer = planner
		o.Formatter = planner
	}
}

func WithBuffer(size uint) Option {
	return func(o *Options) {
		o.Buffer = true
		o.Size = size
	}
}

func WithStorager(storager Storager) Option {
	return func(o *Options) {
		o.Storager = storager
	}
}

func WithErrorChan(ch chan<- error) Option {
	return func(o *Options) {
		o.ErrCh = ch
	}
}

func WithStopChan(stopCh <-chan struct{}) Option {
	return func(o *Options) {
		o.StopCh = stopCh
	}
}
