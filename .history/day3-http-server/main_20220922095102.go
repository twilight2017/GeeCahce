package main

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("Hello world!"))
}

func main() {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
