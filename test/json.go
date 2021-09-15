package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	//Name 必须是大写解析出{Name:candy}，小写会导致解析出{name:}
	Name string `json: "name"`
}

func main() {
	js := `{
		"name": "candy"
	}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Printf("err:%s\r\n", err)
		return
	}
	fmt.Printf("people:%+v\r\n", p)
}
