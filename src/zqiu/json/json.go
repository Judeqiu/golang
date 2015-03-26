package main

import (
	"encoding/json"
	"fmt"
)

type Baby struct {
	Method string "method"
}

func main() {
	b, _ := json.Marshal(Baby{"haha"})
	fmt.Println(string(b))
	var r interface{}
	json.Unmarshal(b, &r)
	fmt.Println(r.(map[string]interface{})["Method"])
}
