package geecache

/*GeeCache不实现直接从数据源获取数据，原因有：
1.数据源的种类太多，无法一一实现
2.扩展性不好
如何从源头获取数据，应该是用户考虑完成实现的部分
*/

// A Getter loads data for a key
type Getter interface{
	//用[]byte存数据，是为了让它能支持任意数据格式
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function
type GetterFunc func(key string) ([]byte, error)

//Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error){
	rteurn f(key)
}