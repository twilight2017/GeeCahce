package main

import "log"
import "net/http"

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("Hello world!"))
}

//任何实现了ServerHTTP方法的对象都可以作为HTTP的handler
func main() {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
