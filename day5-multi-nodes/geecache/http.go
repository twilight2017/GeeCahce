//为HTTPPool实现客户端的功能
//1.创建具体的http客户端httpGetter
package geecache

import (
	"net/http"
	"url"
)

type httpGetter struct {
	baseURL string //将要访问的远程节点的地址
}

//2.实现PeerGetter接口
func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		//QueryEscape方法对传入的string安全地进行解码，使之可以安全地用在URL查询中
		url.QueryEscape(group),
		url.QueryEscape(key),
	)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	return bytes, nil
}

const(
	 defaultBasePath = "/_geecache/"
	 defaultReplicas = 50
)

//HTTPPool implements PeerPicker for a pool of HTTP peers
type HTTPPool struct {
	// this peer's bae URL, e.g. "https://example.net:8000"
	self     string //用来记录自己的地址，包括主机名/IP和端口
	basePath string //节点通讯地址前缀
	mu sync.Mutex //guards peers and httpGetters
	peers *consistenthash.Map  //类型是一致性Hash的Map.用来根据key选择节点
	httpGetters map[string]*httpGetter  //keyed by e.g."http://10.0.0.2:8008"，用于映射远程节点和对应的httpGetter,呈现一一对应关系
}

//NewHTTPPool initialize an HTTP pool of peers
func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

//log info with server name
func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

//ServeHTTP handle all http requests
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}
	p.Log("%s, %s", r.Method, r.URL.Path)
	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group"+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key) //获得缓存数据
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/octet-stream") //将缓存值作为httpResponse的body返回
	w.Write(view.ByteSlice())

}

//Set updates the pool's list of peers
func (p *HTTPPool) Set (peers ...strings){
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers := consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	//申请httpGetter空间
	p.httpGetter = make(map[string] *httpGetter, len(peers))
	for _, peer := range peers{
		p.httpGetter[peer] = &httpGetter{
			baseURL: peer+p.basePath
		}
	}
}

//PickPeer picks a peer according to a key
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool){
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer !=p.self{
		p.log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, true
}

var _ PeerPicker = (*HTTPPool)(nil)