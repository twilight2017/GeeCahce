package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func handlerError(t *testing.T, err error) {
	//t.Helper()， 标注该函数是帮助函数，报错时将输出帮助函数调用者信息
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handlerError(t, err)
	defer ln.close()

	http.HandlerFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handlerError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handlerError(t, err)

	if string(Body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}
