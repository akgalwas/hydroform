package k8s

const (
	OnConflictLabel   string = "on-conflict"
	ReplaceOnConflict string = "replace"
)

type Func func() error

func update(merge Func, update Func, labels map[string]string) error {

	isMerge := func(labels map[string]string) bool {
		val, ok := labels[OnConflictLabel]
		return ok && val == ReplaceOnConflict
	}

	if isMerge(labels) {
		err := merge()
		if err != nil {
			return err
		}
	}

	return update()
}
