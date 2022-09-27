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
