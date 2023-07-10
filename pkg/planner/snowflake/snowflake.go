package snowflake

import (
	"github.com/ElfLabs/id-generator/pkg/format"
	"github.com/ElfLabs/id-generator/pkg/sequencer"
)

/*
+-------------------------------------------------+
|   41 Bit Count   | 10 Bit NodeID  | 12 Bit Step |
+-------------------------------------------------+
*/

const (
	DefaultEpoch = 1234567890123
)

type Snowflake struct {
	Epoch int64
	format.SnowflakeFormat
	*sequencer.Timestamp
}

type Option func(s *Snowflake)

func NewSnowflakePlanner(opts ...Option) *Snowflake {
	s := &Snowflake{
		Epoch:           DefaultEpoch,
		SnowflakeFormat: format.NewSnowflakeFormat(),
		Timestamp:       sequencer.NewTimestampSequencer(),
	}
	s.apply(opts)
	return s
}

func (s *Snowflake) apply(opts []Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s Snowflake) GetEpoch() int64 {
	return s.Epoch
}

func WithEpoch(epoch int64) Option {
	return func(s *Snowflake) {
		s.Epoch = epoch
	}
}
