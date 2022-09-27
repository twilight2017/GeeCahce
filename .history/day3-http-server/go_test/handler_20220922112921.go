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
