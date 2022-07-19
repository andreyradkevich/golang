package main

import "fmt"

func Increment(initial int) func() int {
	i := initial

	return func() int {
		i++
		return i
	}
}

func Decrement(initial int) func() int {
	i := initial

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
