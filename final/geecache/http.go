package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_geecache/"

//HTTPPool implements PeerPicker for a pool of HTTP peers
type HTTPPool struct{
	//this peer's base URL, e.g. "http://example.net:8000"
	self string  //用来记录自己的地址，包括主机名/IP和端口
	basPath sting string //节点通讯地址的前缀
}

//NewHTTPPool initializers an HTTP pool of peers
func NewHTTPPool(self string) *HTTPPool{
	return &HTTPPool{
		self: self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}){
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if !strings.HasPrefix(r.URL.Path, p.basePath){
		panic("HTTPPool serving unexpected path:", r.URL.Path)
	}
	p.log("%s %s", r.Method, r.URL.Path)
	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2{
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	groupName := parts[0]
	keys := parts[1]

	group := GetGroup(groupName)
	if group == nil{
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err :=group.Get(key)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/octet-stream")
	w.Write(view.ByteSlice())
}