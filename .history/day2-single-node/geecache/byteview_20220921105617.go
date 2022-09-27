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

//String returns the data as a string, making a copy if necessary
func (v ByteView) String() string {
	return string(v.b)
}

//returns the copy of []byte
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
