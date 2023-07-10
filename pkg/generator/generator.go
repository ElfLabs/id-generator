package generator

import (
	"errors"
	"strconv"
	"sync"
)

var ErrCountOverflow = errors.New("count overflow")

type generator struct {
	mu sync.Mutex

	region int64
	node   int64
	count  int64
	step   int64
	epoch  int64

	countMax int64
	stepMax  int64

	sequencer Sequencer
	formatter Formatter
}

func NewGenerator(region, node int64, opts ...Option) (Generator, error) {
	return NewGeneratorWithOptions(region, node, NewOptions(opts...))
}

func NewGeneratorWithOptions(region, node int64, options Options) (Generator, error) {
	if err := tryRecoverByStorager(&options); err != nil {
		return nil, err
	}

	g := &generator{
		region:    region,
		node:      node,
		count:     options.GetCount(),
		step:      options.GetStep(),
		epoch:     options.GetEpoch(),
		sequencer: options.Sequencer,
		formatter: options.Formatter,
	}

	if options.Storager != nil {
		g.sequencer.Init(g.count + g.epoch)
	}

	regionMax := GetRegionMax(g.formatter)
	if g.region < 0 || g.region > regionMax {
		return nil, errors.New("Region number must be between 0 and " + strconv.FormatInt(regionMax, 10))
	}

	nodeMax := GetNodeMax(g.formatter)
	if g.node < 0 || g.node > nodeMax {
		return nil, errors.New("generator number must be between 0 and " + strconv.FormatInt(nodeMax, 10))
	}

	g.countMax = GetCountMax(g.formatter)
	g.stepMax = GetStepMax(g.formatter)

	if options.Buffer {
		return NewBufferGeneratorWithOptions(g, options), nil
	}

	return g, nil
}

func (g generator) GetEpoch() int64 {
	return g.epoch
}

func (g generator) GetSequencer() Sequencer {
	return g.sequencer
}

func (g generator) GetFormatter() Formatter {
	return g.formatter
}

func (g generator) Shift() ID {
	id := (g.region << g.formatter.RegionShift()) |
		(g.node << g.formatter.NodeShift()) |
		(g.count << g.formatter.CountShift()) |
		(g.step << g.formatter.StepShift())
	return ID(id)
}

func (g *generator) Generate() ID {
	g.mu.Lock()
	defer g.mu.Unlock()

	current := g.sequencer.Current() - g.epoch

	if current == g.count {
		g.step = (g.step + 1) & g.stepMax
		if g.step == 0 {
			for current <= g.count {
				current = g.sequencer.Next() - g.epoch
			}
		}
	} else {
		g.step = 0
	}

	if current > g.countMax {
		panic(ErrCountOverflow)
	}

	g.count = current

	return g.Shift()
}
