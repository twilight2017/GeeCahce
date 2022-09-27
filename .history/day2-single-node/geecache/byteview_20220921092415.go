package geecache

//使用集合set记录已经打印过的数字
var set = make(map[int]bool, 0)

func printOnce(num int){
	if, _, exit := set[num];!exit{
		fmt.Printf(num)
	}
	set[num=True]
}