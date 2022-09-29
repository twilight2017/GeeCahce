package main

import (
	"net/http"
	"flag"
)

var db := map[string]string{
	"TOM":"3232"
	"JACK":"13"
	"Cassie":"212"
}

func CreateGroup *geecache.Group{
	return geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error){
			log.Println("[SlowDB] search DB", key)
			if v, ok := db[key];ok{
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}
	))
}

//启动缓存服务器，创建HTTPPool，添加节点信息，注册到Gee中，启动3个http服务
func startCacheServer(addr string, addrs []string, gee *geecache.Group){
	peers := geecache.NewHTTPPool(addr)
	peers.Set(addrs)
	gee.RegisterPeers(peers)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServer(addr[7:], peers))
}

//启动一个API服务（端口9999），与用户交互
func startAPIServer(apiAddr string, gee *geecache.Group){
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			key:= r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.write(view.ByteSlice())
		}
	))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServer(apiAddr[7:], nil))
}

//在指定端口启动HTTP服务
func main(){
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int][string]{
		8001: "http://localhost:8001",
		8002: "http://localohost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap{
		addrs = append(addrs, v)
	}

	gee := CreateGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], []strings(addrs), gee)
}