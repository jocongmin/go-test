package main

import (
	"fmt"
)

type TestInter interface {
	say()
}

type Datas struct {
	one int
}

func (m Datas) say() {
	fmt.Println(m.one, "mmmm")
}

func main() {
	var inter TestInter
	var a Datas
	a.one = 23423423
	inter = a
	inter.say()
}
