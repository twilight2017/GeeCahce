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
}
