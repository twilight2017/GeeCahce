package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_geecache/"

//HTTPPool implements PeerPicker for a pool of HTTP peers
type HTTPPool struct {
	// this peer's bae URL, e.g. "https://example.net:8000"
	self     string
	basePath string
}

//NewHTTPPool initialize an HTTP pool of peers
func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}
