package main

import (
	"fmt"
)

type Gotodo interface {
	sayOne()
	sayTwo()
	sayThree()
}

func (m BaseStr) sayOne() {
	fmt.Println(m.one)
}

func (m BaseStr) sayTwo() {
	fmt.Println(m.two)
}

func (m BaseStr) sayThree() {
	fmt.Println(m.three)
}

type BaseStr struct {
	one   string
	two   string
	three string
}

func main() {
	var abc BaseStr
	abc.one = "skldfjsdjkf"
	abc.two = "sdklfjsdkjf"
	abc.three = "sdklfjsdkfj"
	var fn Gotodo
	fn = abc
	fn.sayOne()
	fn.sayTwo()
	fn.sayThree()

}
