package main

import (
	"fmt"
)

func main() {
	goFor()
}
func test() {
	tb := make(map[string]interface{})
	arr := [3]int{3423, 324, 23432}
	tb["one"] = "sdlkfjsdlkfj"
	tb["two"] = "sdlkfjsdlkfj"
	tb["three"] = arr
	fmt.Println(tb["three"], "tb")
}

func appd() {
	arr := []int{23423, 32432432, 324234}
	b := append(arr, 23423, 3423, 4, 32, 432, 432, 4, 324, 32, 4, 32432)
	fmt.Println(b)
}

func goFor() {
	i := 6
	var t int
	for i < 50 {
		i++
		t += i
	}
	fmt.Println(t)
}
