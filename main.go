package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HelloServie interface {
	SayHello() (string, error)
}

type Hello struct {
}

func (h Hello) SayHello() (string, error) {
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:8080")
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go World!")
}

func main() {
	helloService := Hello{}
	resp, err := helloService.SayHello()
	if err != nil {
		fmt.Printf("err:%+v", err)
		return
	}
	fmt.Printf("resp: %+v", resp)
}
