package gdhelper

func IIF[T any](v bool, trueVal T, falseVal T) T {
	if v {
		return trueVal
	}

	return falseVal
}
