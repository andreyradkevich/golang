package main

import "fmt"

func CreateSlice[T int | string](length int) []T {
	return make([]T, length)
}

func main() {
	intSlice := CreateSlice[int](10)
	stringSlice := CreateSlice[string](10)

	fmt.Println(intSlice)    // [0,0,0,0,0,0,0,0,0,0]
	fmt.Println(stringSlice) // [         ]
}

