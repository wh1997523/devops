package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (u User) NewUser() string {
	return u.Name + strconv.Itoa(u.Age)
}

func main() {
	n := "whh"
	a := 17
	nu := User{Name: n, Age: a}
	res := nu
	fmt.Println(res)

}
