package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

//Hash maps bytes to unit32
type Hash func(data []byte) uint32

//Map constains all hash keys
type Map struct {
	hash     Hash
	replicas int   //虚拟节点倍数
	keys     []int //哈希环
	//键是虚拟节点的哈希值，值是真实节点的名称
	hashMap map[int]string //虚拟节点与真实节点映射表
}

func New(replicas int, hash Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     hash,
		hashMap:  make(map[int]string),
	}
	if hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

//Add方法用于添加真实节点或者机器， Add adds some keys to the hash
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			//1.计算hash值，通过添加编号的方式区分不同的虚拟节点
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			//2.把新增的节点放置在hash环上
			m.keys = append(m.keys, hash)
			//3.增加虚拟节点和真实节点的映射关系
			m.hashMap[hash] = key
		}
	}
	//给节点环进行排序
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	//1.根据节点的key计算得到节点的hash值
	hash := int(m.hash([]byte(key)))
	//2.查找节点,通过二分法找到满足条件的最小索引
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
