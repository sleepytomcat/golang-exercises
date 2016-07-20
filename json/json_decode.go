package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var animalsJson = []byte(`[{"Name": "Platypus", "Order": "Monotremata"}, {"Name": "Quoll",    "Order": "Dasyuromorphia"}]`)

	var decoded interface{}
	err := json.Unmarshal(animalsJson, &decoded)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(decoded)
}
