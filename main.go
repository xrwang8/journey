package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type HelloServie interface {
	SayHello(name string) (string, error)
}

type Hello struct {
	endpoint  string
	FuncFiled func()
}

var _ HelloServie = &Hello{}

func (h *Hello) SayHello(name string) (string, error) {
	client := &http.Client{}
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

// try reflect set  field value
func PrintFuncName(val interface{}) {
	// reflect pointer
	v := reflect.ValueOf(val)
	// Gets the struct represented by a pointer
	elem := v.Elem()
	// struct information
	t := elem.Type()

	field := t.NumField()
	for i := 0; i < field; i++ {
		fieldName := t.Field(i)
		value := elem.Field(i)
		if value.CanSet() {
			fmt.Printf("%s Can be Set", fieldName.Name)
		}

	}

}

func main() {
	// 没有构造函数，直接使用结构体
	helloService := &Hello{endpoint: "127.0.0.1:8080"}
	//resp, err := helloService.SayHello("golang")
	//if err != nil {
	//	fmt.Printf("err:%+v", err)
	//	return
	//}
	//fmt.Printf("resp: %+v", resp)

	PrintFuncName(helloService)

}
