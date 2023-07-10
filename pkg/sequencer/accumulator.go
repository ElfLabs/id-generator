package sequencer

type Accumulator struct {
	count int64
}

func (t *Accumulator) Init(count int64) {
	t.count = count
}

func (a Accumulator) Current() int64 {
	return a.count
}

func (a *Accumulator) Next() int64 {
	a.count += 1
	return a.count
}
