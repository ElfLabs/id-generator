package format

type Scope struct {
	Bits  uint8 `json:"bits" yaml:"bits"`
	Shift uint8 `json:"shift" yaml:"shift"`
}

type Format struct {
	Region Scope `json:"region" yaml:"region"`
	Node   Scope `json:"node" yaml:"node"`
	Count  Scope `json:"count" yaml:"count"`
	Step   Scope `json:"step" yaml:"step"`
}

func (f Format) RegionBits() uint8 {
	return f.Region.Bits
}

func (f Format) NodeBits() uint8 {
	return f.Node.Bits
}

func (f Format) CountBits() uint8 {
	return f.Count.Bits
}

func (f Format) StepBits() uint8 {
	return f.Step.Bits
}

func (f Format) RegionShift() uint8 {
	return f.Region.Shift
}

func (f Format) NodeShift() uint8 {
	return f.Node.Shift
}

func (f Format) CountShift() uint8 {
	return f.Count.Shift
}

func (f Format) StepShift() uint8 {
	return f.Step.Shift
}

func (f Format) RegionMax() int64 {
	return -1 ^ (-1 << f.Region.Bits)
}

func (f Format) NodeMax() int64 {
	return -1 ^ (-1 << f.Node.Bits)
}

func (f Format) CountMax() int64 {
	return -1 ^ (-1 << f.Count.Bits)
}

func (f Format) StepMax() int64 {
	return -1 ^ (-1 << f.Step.Bits)
}

func (f Format) RegionMask() int64 {
	return (-1 ^ (-1 << f.Region.Bits)) << f.Region.Shift
}

func (f Format) NodeMask() int64 {
	return (-1 ^ (-1 << f.Node.Bits)) << f.Node.Shift
}

func (f Format) CountMask() int64 {
	return (-1 ^ (-1 << f.Count.Bits)) << f.Count.Shift
}

func (f Format) StepMask() int64 {
	return (-1 ^ (-1 << f.Step.Bits)) << f.Step.Shift
}
