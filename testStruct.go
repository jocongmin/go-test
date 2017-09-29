package main

import "encoding/json"
import "fmt"

type YearTableData struct {
	MouthArr []MouthItemStruct
}

type MouthItemStruct struct {
	Mouth      int
	DetailItem []DetailItemStruct
}
type DetailItemStruct struct {
	Key int
	One string
	Two string
}

func main() {
	var data YearTableData
	var detailData []DetailItemStruct
	var mouthItemData []MouthItemStruct
	for t := 1; t < 6; t++ {
		for i := 1; i < 6; i++ {
			detailData = append(detailData, DetailItemStruct{Key: i, One: "sdlkfjsdjf", Two: "sdlkfjsdjkf"})
		}
		mouthItemData = append(mouthItemData, MouthItemStruct{Mouth: t, DetailItem: detailData})
	}

	data.MouthArr = mouthItemData
	bytes, _ := json.Marshal(data)
	fmt.Println(string(bytes))
}
