package closure

func Closure() func() int {
	x := 10

	return func() int {
		return x + 1
	}
}
