package main

import (
	"fmt"
	"io"
	"net/http"
	// "os"
)

func proxy(res http.ResponseWriter, req *http.Request) {
	acceptContentType := req.Header.Get("Accept")
	fmt.Println(req.RequestURI)
	res.Header().Set("Content-Type", acceptContentType)
	io.WriteString(res, "{'name':'value'}")
}

func main() {
	fmt.Println("start service")

	http.HandleFunc("/", proxy)
	http.ListenAndServe(":8080", nil)
}
