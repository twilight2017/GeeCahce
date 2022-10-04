package geecache

//PeerPicker is the interface that must be implemented to locate
//the peers that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerPicker, ok bool)
}
