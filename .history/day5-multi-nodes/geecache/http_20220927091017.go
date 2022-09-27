//为HTTPPool实现客户端的功能
//1.创建具体的http客户端httpGetter
package geecache

import (
	"url"
	"net/http"
)

type httpGetter struct{
	baseURL string
}

//2.实现PeerGetter接口
func (h *httpGetter) Get(group string, key string) ([]byte, error){
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		//QueryEscape方法对传入的string安全地进行解码，使之可以安全地用在URL查询中
		url.QueryEscape(group),
		url.QueryEscape(key),
	)
	res, err := http.Get(u)
	if err != nil{
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK
}