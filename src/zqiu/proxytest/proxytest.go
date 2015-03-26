package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start request")
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(0)
	}
	defer resp.Body.Close()
}
