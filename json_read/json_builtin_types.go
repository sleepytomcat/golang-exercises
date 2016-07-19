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

func main() {
	PrintJson(true)
	PrintJson(1)
	PrintJson(1.2345)
	PrintJson("test string")
	PrintJson([3]int{1, 2, 3})

	numbers := [4]int{2, 3, 5, 7}
	var slice []int = numbers[1:3]
	PrintJson(slice)

	var fruits map[string]int = map[string]int {
		"apples":   7,
		"peaches":  3,
		"aprictos": 12,
	}
	PrintJson(fruits)
}
