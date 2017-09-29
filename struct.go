package main

import "encoding/json"
import "fmt"

type AlipayRemoteReqStruct struct {
	Ono         string
	OrderItem   []Item
	OrderRefund []Refund
}
type Item struct {
	Ono string
	Oid int
}
type Refund struct {
	Ono     string
	Item    int
	Content string
	Imgs    string
	Status  string
}

func main() {
	var m AlipayRemoteReqStruct
	m.Ono = "12345"
	m.OrderItem = append(m.OrderItem, Item{Ono: "Shanghai_VPN", Oid: 1})
	m.OrderItem = append(m.OrderItem, Item{Ono: "Beijing_VPN", Oid: 2})
	bytes, _ := json.Marshal(m)
	fmt.Printf("json:m,%s\n", bytes)
	var js AlipayRemoteReqStruct
	err := json.Unmarshal(bytes, &js)
	if err != nil {
		fmt.Printf("format err:%s\n", err.Error())
		return
	}
}
