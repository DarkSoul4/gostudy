package main

import (
	"fmt"
)

type Sayer interface {
	say()
}

type dog struct {}

type cat struct {}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	var x Sayer
	a := cat{}
	x = a
	x.say()
}
