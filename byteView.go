package main

type ByteView struct {
	b []byte
}

func (b ByteView) Len() int {
	return len(b.b)
}

func (b ByteView) ByteSlice() []byte {
	return cloneBytes(b.b)
}

// turn type []byte to string
func (b ByteView) String() string {
	return string(b.b)
}

// make a new slice copy
// 防止被修改缓存值
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c

}
