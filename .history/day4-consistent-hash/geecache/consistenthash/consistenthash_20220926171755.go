package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

//Hash map bytes to uint32
func Hash func(data []byte) uint32

//Map constains all hashed keys
type Map struct{
	hash Hash  //hash函数
	replicas int //虚拟节点倍数
	keys []int //哈希环
	//该映射表：键是虚拟节点的哈希值，值是真实节点的名称
	hashMap map[int]string  //虚拟节点与真实节点的映射表
}

//New creates a Map instance
//构造函数New允许自定义虚拟节点倍数和hash函数
func New(reolicas int, fn Hash) *Map{
	m := &Map{
		replicas: replicas,
		hash: fn,
		hashMap: make(map[int]string)
	}
	if m.hash == nil{
		m.hash = crc32.ChecksumIEEE //默认使用的Hash算法
	}
	return m
}

//该方法添加真实的节点/机器
// Add adds some keys to the hash
//允许传入0个或多个真实节点的名称
func (m *Map) Add(key ...string){
	for _, key := range keys{
		//对每个真实节点，创建replicas个数个虚拟节点，虚拟节点的名称通过strconv.Itoa(i)+key做区分
		for i := 0;i<m.replicas;i++{
			hash := int(m.hash([]byte(strconv.Itoa(i)+key)))
			m.keys = append(m.keys, hash)//把hash值添加到环上
			m.hashMap[hash] = key //添加虚拟节点和真实节点的映射关系
		}
	}
	sort.Ints(m.keys) //环上的hash值做排序
}

//Get gets the closest item in the hash to the provided key
func(m *Map) Get(key string) string{
	//hash环上没有节点的情况
	if len(m.keys==0){
		return ""
	}

	//1.计算key的hash值
	hash := int(m.hash([]byte(key)))
	//Binary search for appropriate replica
	//2.顺时针找到第一个匹配的虚拟节点的下标，用取余数的方式来选择顺时针下距离最近的虚拟节点
	idx := sort.Search(len(m.keys), func(i int) bool{
		return m.keys[i] >= hash
	})
    //3.通过hashMap映射得到真实的节点
	return m.hashMap[m.keys[idx%len(m.keys)]]
}