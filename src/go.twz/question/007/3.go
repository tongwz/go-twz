package main

import (
	"encoding/json"
	"fmt"
)

type People1 struct {
	Name string `json:"name"`
}

func main() {
	js := `{
        "name":"11"
    }`
	var p People1
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("people: %#v", p)
}
