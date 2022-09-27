package geecache

// A ByteView holds an immutable view of bytes
//一个只读数据结构来表示缓存值
type ByteView struct {
	b []byte
}

//Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

//ByteSlice returns a copy of the data as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}
