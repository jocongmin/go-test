package main

import "encoding/json"
import "fmt"

type TotalTableItem struct {
	Year    int
	Sum     DetailStruct
	Average DetailStruct
}

type DetailStruct struct {
	One   int
	Two   int
	Three int
}

func main() {
	var totalData []TotalTableItem
	var sum DetailStruct
	var average DetailStruct
	for i := 1; i < 8; i++ {
		sum = DetailStruct{
			One:   i,
			Two:   i,
			Three: i,
		}
		average = DetailStruct{
			One:   i * 100,
			Two:   i * 100,
			Three: i * 110,
		}
		totalData = append(totalData, TotalTableItem{Year: i * 2000, Sum: sum, Average: average})
	}
	bytes, _ := json.Marshal(totalData)
	fmt.Println(string(bytes))
}
