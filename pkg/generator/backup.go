package generator

type Backup struct {
	ID     ID    `json:"id"`
	Region int64 `json:"region"`
	Node   int64 `json:"node"`
	Count  int64 `json:"count"`
	Step   int64 `json:"step"`
	Epoch  int64 `json:"epoch"`
}

func NewBackup(id ID, formatter Formatter, epoch int64) Backup {
	var backup = Backup{
		ID:    id,
		Epoch: epoch,
	}

	backup.Region, backup.Node, backup.Count, backup.Step = ParseID(id, formatter)

	return backup
}
