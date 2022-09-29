package singleflight

import "sync"

//正在进行中的，或已经结束的请求，使用sync.WaitGroup避免重入
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

//singleflight的主数据结构，管理不同key的请求(call)
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

//Do方法，接收两个参数（第一个参数是key，第二个参数是fn）
//Do的作用就是，针对相同的key，无论Do被调用多少次，函数fn只会被调用一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock() //保护Group
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Lock()
		c.wg.Wait()         //如果请求正在进行中，则等待
		return c.val, c.err //请求结束，返回结果
	}
	c := new(call)
	c.wg.Add(1)  //发起请求前加锁
	g.m[key] = c //添加到g.m，表明key已经有对应的请求正在处理
	g.mu.Unlock()

	c.val, c.err = fn() //调用fn，发起请求
	c.wg.Done()         //请求结束

	g.mu.Lock()
	delete(g.m, key) //更新g.m
	g.mu.Unlock()

	return c.val, c.err //返回结果
}
