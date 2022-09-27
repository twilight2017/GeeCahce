package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0") //监听一个未被占用的端口，并返回Listener
	handlerError(t, err)                        //根据Listener启动http服务
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello") //发起一个Get请求，并检查返回值是否正确
	handlerError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handlerError(t, err)

	if string(body) != "Hello world" {
		t.Fatal("expected Hello world, but got", string(body))
	}
}
