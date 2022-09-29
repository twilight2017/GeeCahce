package geecache

//PeerPicky is the interface that must be implemented to locate
//the peer that owns a special key
//根据传入的key选择相应节点的PeerGetter
type PeerPicker interface{
	PickPeer (key string) (peer PeerGetter, ok bool)
	Get(group string, key string)([]byte, error)
}