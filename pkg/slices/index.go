package slices

func CreateSlice[T int | string](length int) []T {
	return make([]T, length)
}
