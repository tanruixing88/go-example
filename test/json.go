package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Attr struct {
	Interest string `json:"interest"`
}
type People struct {
	//Name 必须是大写解析出{Name:candy}，小写会导致解析出{name:}
	Name string `json: "name"`
	User string `json: "user"`
	Attr *Attr  `json: "attr"`
}

//json 解析字段结构要大写
func capitalStructVal() {
	js := `{
		"name": "candy",
		"user": "haha",
        "age": 20
	}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Printf("err:%s\r\n", err)
		return
	}
	fmt.Printf("people:%+v\r\n", p)
}

type J struct {
	a string //小写无tag
	b string `json:"B"` //小写加tag
	C string //大写无tag
	D string `json:"DD" other:"good"` //大写加tag
}

//展示结构体定义的tag
func showTag() {
	j := J{
		a: "1",
		b: "2",
		C: "3",
		D: "4",
	}
	t := reflect.TypeOf(&j).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%d th item %+v json tag:%+v other tag:%+v \r\n",
			i+1, t.Field(i).Name, t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("other"))
	}
}

func main() {
	capitalStructVal()
	showTag()
}
