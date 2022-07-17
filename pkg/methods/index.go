package main

import "fmt"

// struct - similar to js object
type User struct {
	Age  string
	Name string
}

// methods on struct type
// method is a function on struct, in js method is a function on object

func (u User) GetAge() string {
	return u.Age
}

func (u User) GetName() string {
	return u.Name
}

func main() {

	strct := User{Name: "Andrew", Age: "28"}

	fmt.Println(strct.GetAge(), strct.GetName(), strct.Age, strct.Name) // 28 Andrew 28 Andrew
}
