package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func PrintJson(value interface{}) {
	var jsonBuffer []byte
	jsonBuffer, _ = json.Marshal(value)
	fmt.Println(value, "("+reflect.TypeOf(value).String()+") as JSON:", string(jsonBuffer))
}

type TestData struct {
	Name1 int
	name2 int // lowercase name, ignored by encoding/json
	Name3 int `json:"name3"` // custom JSON name for this one
	Name4 int `json:"-"`     // field marked to be ignored
}

func main() {
	data := TestData{1, 2, 3, 4}
	PrintJson(data)
}
