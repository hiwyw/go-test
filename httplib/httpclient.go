package main

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/buger/jsonparser"
)

func main() {
	url := "https://192.168.134.231:20120/views"
	user := "admin"
	password := "admin"

	req := httplib.Get(url).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetTimeout(100*time.Second, 30*time.Second).SetBasicAuth(user, password).Header("Content-type", "application/json")

	response, err := req.Bytes()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))
	v, t, _, err := jsonparser.Get(response, "resources")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v), t)
}
