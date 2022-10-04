package geecache

//A ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

//Len returns the view's length
func (v *ByteView) Len() int {
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

//选择byte是为了能够支持任意数据格式类型的存储
func cloneBytes(b []byte) []byte{
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
