package mid

/*
+-----------------------------------------------------------------+
|   24 Bit Count   | 2 Bit NodeID |  4 Bit Step  | 2 Bit RegionID |
+-----------------------------------------------------------------+
*/

const (
	CountBits  = 24
	NodeBits   = 2
	StepBits   = 4
	RegionBits = 2

	CountShift  = RegionBits + StepBits + NodeBits
	NodeShift   = RegionBits + StepBits
	StepShift   = RegionBits
	RegionShift = 0
)

const (
	DefaultStartMid = 100000000
	MaxMid          = 4294967295
)

type Mid struct {
	count int64
}

func NewMidPlanner(startId int64) *Mid {
	return &Mid{
		count: startId >> CountShift,
	}
}

func (m *Mid) Init(count int64) {
	m.count = count
}

func (m Mid) Current() int64 {
	return m.count
}

func (m *Mid) Next() int64 {
	m.count += 1
	return m.count
}

func (m Mid) GetEpoch() int64 {
	return 0
}

func (m Mid) RegionBits() uint8 {
	return RegionBits
}

func (m Mid) NodeBits() uint8 {
	return NodeBits
}

func (m Mid) CountBits() uint8 {
	return CountBits
}

func (m Mid) StepBits() uint8 {
	return StepBits
}

func (m Mid) RegionShift() uint8 {
	return RegionShift
}

func (m Mid) NodeShift() uint8 {
	return NodeShift
}

func (m Mid) CountShift() uint8 {
	return CountShift
}

func (m Mid) StepShift() uint8 {
	return StepShift
}
