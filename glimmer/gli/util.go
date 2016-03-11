package gli

type Deleter interface {
	Delete()
}

func SafeDelete(deleters ...Deleter) {
	for _, deleter := range deleters {
		if deleter != nil {
			deleter.Delete()
		}
	}
}
