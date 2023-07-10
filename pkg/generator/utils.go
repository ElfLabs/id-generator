package generator

import "os"

func Abs(n int64) int64 {
	y := n >> 63       // y ← x >> 63
	return (n ^ y) - y // (x ⨁ y) - y
}

func GetMax(bits uint8) int64 {
	return -1 ^ (-1 << bits)
}

func GetMask(bits, shift uint8) int64 {
	return GetMax(bits) << shift
}

func GetOriginal(id ID, bits, shift uint8) int64 {
	return (int64(id) & GetMask(bits, shift)) >> shift
}

func GetRegionMax(formatter Formatter) int64 {
	return -1 ^ (-1 << formatter.RegionBits())
}

func GetNodeMax(formatter Formatter) int64 {
	return -1 ^ (-1 << formatter.NodeBits())
}

func GetCountMax(formatter Formatter) int64 {
	return -1 ^ (-1 << formatter.CountBits())
}

func GetStepMax(formatter Formatter) int64 {
	return -1 ^ (-1 << formatter.StepBits())
}

func ParseID(id ID, formatter Formatter) (region, node, count, step int64) {
	region = GetOriginal(id, formatter.RegionBits(), formatter.RegionShift())
	node = GetOriginal(id, formatter.NodeBits(), formatter.NodeShift())
	count = GetOriginal(id, formatter.CountBits(), formatter.CountShift())
	step = GetOriginal(id, formatter.StepBits(), formatter.StepShift())
	return
}

func tryRecoverByStorager(options *Options) error {
	if options.Storager == nil {
		return nil
	}

	var backup Backup
	err := options.Storager.Get(&backup)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	if options.Epoch == 0 {
		options.Epoch = backup.Epoch
	}
	if options.Count == 0 {
		options.Count = backup.Count
	}
	if options.Step == 0 {
		options.Step = backup.Step
	}

	return nil
}
