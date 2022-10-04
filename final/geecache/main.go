package main

import (
	"fmt"
	"geecache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":"90"
	"Jack":"678"
	"Sam":"212"
}


func CreateGroup() *geecache.Group{
	return geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error){
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key];ok{
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		} 
	))
}

func startCacheServer(addr string, addrs []string, gee *geecache.Group){
	peers := geecache.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe[7:], peers)
}

func startAPIServer(apiAddr string, gee *geecache.Group){
	http.Handle("/api", httpHandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}
	))
	log.Println("fonted server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}
func main(){
	geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error){
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key];pk{
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}
	))
	addr := "localohost:9999"
	peers := geecache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
