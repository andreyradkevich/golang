package main

import (
	"fmt"
	"golang/pkg/closure"
	"golang/pkg/player"
	"golang/pkg/slices"
	"golang/pkg/user"
)

func main() {

	userObj := user.User{Age: "10", Name: "andrew"}
	playerObj := player.Player{Name: "Pl"}
	closureFunc := closure.Closure()
	slice := slices.CreateSlice[int](10)

	fmt.Println(user.GetAgeNameString(userObj), playerObj.GetName(), closureFunc(), slice)
}
