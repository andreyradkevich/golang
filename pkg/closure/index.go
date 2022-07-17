package main

import "fmt"

func Increment(num *int) func() {
	return func() {
		*num += 1
		fmt.Println(*num, ": +")
	}
}

func Decrement(num *int) func() {
	return func() {
		*num -= 1
		fmt.Println(*num, ": -")
	}
}

func main() {
	num := 10
	increment := Increment(&num)
	decrement := Decrement(&num)

	increment() // 11: +
	increment() // 12: +
	decrement() // 11: -
	decrement() // 10: -
	decrement() // 9: -
}
