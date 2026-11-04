func twoSum(nums []int, target int) []int {
	var map1 = make(map[int]int)
	for index1 ,value := range nums {//这里的思想就是边查 如果没查到就把这个值存入到map中。如果查到了就return
		if index2 , ok := map1[target - value] ; ok {
			return []int{index1 , index2}
		}
		map1[value]=index1
	}
	return nil//go语言函数如果有放回值就必须return。
}

//这是我刷的第一个leetcode，说实话虽然很简单但是还是很巧妙