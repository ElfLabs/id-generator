package generator

import (
	"fmt"
)

const (
	DefaultIDChanSize = 1000
)

type buffer struct {
	generator Generator
	storager  Storager
	idCh      chan ID
	errCh     chan<- error
}

func NewBufferGenerator(generator Generator, opts ...Option) Generator {
	return NewBufferGeneratorWithOptions(generator, NewOptions(opts...))
}

func NewBufferGeneratorWithOptions(generator Generator, options Options) Generator {
	b := &buffer{
		generator: generator,
	}

	b.apply(options)

	return b
}

func (b *buffer) apply(options Options) {
	if options.Size == 0 {
		options.Size = DefaultIDChanSize
	}
	b.storager = options.Storager
	b.idCh = make(chan ID, options.Size)
	b.errCh = options.ErrCh
	if options.StopCh != nil {
		go b.Run(options.StopCh)
	}
}

func (b *buffer) run() {
	defer func() {
		if e := recover(); e != nil {
			switch err := e.(type) {
			case error:
				b.reportError(err)
			default:
				b.reportError(fmt.Errorf("%v", err))
			}
		}
	}()
	b.idCh <- b.generator.Generate()
}

func (b *buffer) Run(done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			b.run()
		}
	}
}

func (b *buffer) reportError(err error) {
	if err == nil {
		return
	}
	if b.errCh == nil {
		return
	}
	b.errCh <- err
}

func (b *buffer) store(id ID) {
	var formatter Formatter
	if i, ok := b.generator.(interface {
		GetFormatter() Formatter
	}); ok {
		formatter = i.GetFormatter()
	}
	if formatter == nil {
		return
	}

	var epoch int64 = 0

	if i, ok := b.generator.(interface {
		GetEpoch() int64
	}); ok {
		epoch = i.GetEpoch()
	}
	backup := NewBackup(id, formatter, epoch)

	err := b.storager.Set(backup)
	if err != nil {
		b.reportError(err)
		return
	}
}

func (b *buffer) Generate() ID {
	id := <-b.idCh
	if b.storager != nil {
		b.store(id)
	}
	return id
}
