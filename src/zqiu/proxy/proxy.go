package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	// "os"
)

func proxy(res http.ResponseWriter, req *http.Request) {

	acceptContentType := req.Header.Get("Accept")
	requestURI := req.RequestURI
	path := strings.Split(requestURI, "/")
	protocol := path[1]
	domain := path[2]
	realUrl := protocol + "://" + domain

	for i, v := range path {
		if i > 2 {
			realUrl += "/" + v
		}
	}

	fmt.Println("try to ", req.Method, " to ", realUrl)
	var resp *http.Response
	var err error

	client := http.Client{}

	switch req.Method {
	default:
		newReq, newErr := http.NewRequest(req.Method, realUrl, req.Body)

		for name, values := range req.Header {
			for _, value := range values {
				newReq.Header.Add(name, value)
			}

		}

		err = newErr
		if err == nil {

			resp, err = client.Do(newReq)

		}
	}
	if err != nil {
		fmt.Println(err)
		res.Header().Set("Content-Type", acceptContentType)
		io.WriteString(res, "{'error':"+err.Error()+"}")
	} else {
		defer resp.Body.Close()
		res.Header().Set("Content-Type", acceptContentType)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			io.WriteString(res, "{'error':"+err.Error()+"}")
		} else {
			io.WriteString(res, string(body))
		}

	}
}

func main() {
	fmt.Println("start service")

	http.HandleFunc("/", proxy)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
