package format

/*
+-------------------------------------------------+
|   41 Bit Count   | 10 Bit NodeID  | 12 Bit Step |
+-------------------------------------------------+
*/

const (
	CountBits = 41
	NodeBits  = 10
	StepBits  = 12

	CountShift = NodeBits + StepBits
	NodeShift  = StepBits
	StepShift  = 0
)

type SnowflakeFormat struct{}

func NewSnowflakeFormat() SnowflakeFormat {
	return SnowflakeFormat{}
}

func (f SnowflakeFormat) RegionBits() uint8 {
	return 0
}

func (f SnowflakeFormat) NodeBits() uint8 {
	return NodeBits
}

func (f SnowflakeFormat) CountBits() uint8 {
	return CountBits
}

func (f SnowflakeFormat) StepBits() uint8 {
	return StepBits
}

func (f SnowflakeFormat) RegionShift() uint8 {
	return 0
}

func (f SnowflakeFormat) NodeShift() uint8 {
	return NodeShift
}

func (f SnowflakeFormat) CountShift() uint8 {
	return CountShift
}

func (f SnowflakeFormat) StepShift() uint8 {
	return StepShift
}
