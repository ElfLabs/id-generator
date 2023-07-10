package sequencer

import (
	"errors"
	"time"
)

const (
	DefaultTimestampEpoch = 1234567890123
	DefaultMaxBackwards   = 10
)

var ErrTimeRollback = errors.New("time rollback")

type TimeBackwardHandler func(offset int64)

type Timestamp struct {
	MaxBackwards    int64
	backwardHandler TimeBackwardHandler
	latest          int64
}

type TimestampOption func(timestamp *Timestamp)

func NewTimestampSequencer(opts ...TimestampOption) *Timestamp {
	t := Timestamp{
		MaxBackwards: DefaultMaxBackwards,
	}

	for _, opt := range opts {
		opt(&t)
	}
	if t.backwardHandler == nil {
		t.backwardHandler = t.defaultBackwardHandler
	}

	return &t
}

func (t *Timestamp) defaultBackwardHandler(offset int64) {
	panic(ErrTimeRollback)
}

func (t *Timestamp) Init(count int64) {
	t.latest = count
}

func (t Timestamp) Current() int64 {
	return time.Now().UnixNano() / 1e6
}

func (t *Timestamp) Next() int64 {
	now := time.Now().UnixNano() / 1e6

	for now <= t.latest {
		offset := t.latest - now
		if offset > t.MaxBackwards {
			t.backwardHandler(offset)
		} else {
			time.Sleep(time.Millisecond)
		}
		now = time.Now().UnixNano() / 1e6 // waiting next millisecond
	}
	t.latest = now

	return now
}

func WithLatestTimestamp(latest int64) TimestampOption {
	return func(timestamp *Timestamp) {
		timestamp.latest = latest
	}
}

func WithMaxBackwards(n int64) TimestampOption {
	return func(timestamp *Timestamp) {
		timestamp.MaxBackwards = n
	}
}

func WithTimeBackwardHandler(handle TimeBackwardHandler) TimestampOption {
	return func(timestamp *Timestamp) {
		if handle == nil {
			return
		}
		timestamp.backwardHandler = handle
	}
}
