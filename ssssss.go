package main

import "encoding/json"
import "fmt"
import "strconv"
import (
	"github.com/drone/routes"
	"net/http"
)

type ResInfo struct {
	Data YearDataStruct
	Msg  string
}

type YearDataStruct struct {
	MouthAll []MouthStruct
	Sum      DetailStruct
	Average  DetailStruct
	Quarter  []QuarterStruct
}
type DetailStruct struct {
	One   int
	Two   int
	Three int
}
type QuarterStruct struct {
	DetailStruct
	QuarterNum int
}

type MouthStruct struct {
	Mouth        int
	PartmentItem []ItemArrStruct
}
type ItemArrStruct struct {
	PartMent string
	Detail   DetailStruct
}

func cross(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
}
func main() {
	fmt.Println("正在启动WEB服务...")
	var mux *routes.RouteMux = routes.New()
	mux.Get("/test", DataSend)
	//http.Handle("/", mux)
	http.ListenAndServe(":8088", mux)
	fmt.Println("服务已停止")
}

func DataSend(w http.ResponseWriter, r *http.Request) {
	var data YearDataStruct
	for i := 1; i < 13; i++ {
		var partMent []ItemArrStruct
		for t := 1; t < 4; t++ {
			partMent = append(partMent, ItemArrStruct{
				PartMent: "" + strconv.Itoa(t) + "abc",
				Detail: DetailStruct{
					One:   45,
					Two:   23,
					Three: 3432,
				},
			})
		}
		data.MouthAll = append(data.MouthAll, MouthStruct{
			Mouth:        i,
			PartmentItem: partMent,
		})
	}
	for i := 1; i <= 4; i++ {
		data.Quarter = append(data.Quarter, QuarterStruct{
			DetailStruct: DetailStruct{
				One:   4324,
				Two:   23423,
				Three: 4324234,
			},
			QuarterNum: i,
		})
	}
	data.Sum = DetailStruct{
		One:   4546,
		Two:   454,
		Three: 454,
	}
	data.Average = DetailStruct{
		One:   4546,
		Two:   454,
		Three: 465465,
	}
	var res ResInfo
	res.Data = data
	res.Msg = "right"
	resJson, _ := json.Marshal(res)
	fmt.Fprintln(w, string(resJson))
}
