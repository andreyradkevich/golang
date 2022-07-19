package main

import "fmt"

func Increment(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func Decrement(i int) func() int {
	return func() int {
		i--
		return i
	}
}

func main() {
	num := 10
	increment := Increment(num)
	decrement := Decrement(num)

	fmt.Println(increment(), decrement(), decrement()) // 11 9 8
}
