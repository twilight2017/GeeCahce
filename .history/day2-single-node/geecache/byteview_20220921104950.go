package main

import (
	"fmt"
	"sync"
	"time"
)

//使用集合set记录已经打印过的数字
var set = make(map[int]bool, 0)
var m sync.Mutex //申请互斥锁

func printOnce(num int) {
	m.Lock()
	defer m.Unlock()
	if _, exit := set[num]; !exit {
		fmt.Println(num)
	}
	set[num] = true
}

func main() {
	//打开10个并发进程
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
