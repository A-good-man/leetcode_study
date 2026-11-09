func longestConsecutive(nums []int) int {
    if len(nums) == 0 {  //对空数组的一个判断
    return 0
    }
    sort.Slice(nums, func(i,j int) bool{//升序排序
        return nums[i] < nums[j]
    })
    var nums1 = make([]int , 0 , len(nums))
    nums1 = append(nums1 , nums[0])//去重操作，注意这里要处理空数组。
    for i := 1 ; i < len(nums) ;i++{
        if nums[i] != nums[i-1] {
            nums1 = append(nums1 , nums[i])
        }
    }
    //处理去重后的数组
    var count int = 1
    var max  int = 1
    for i := 0 ; i < len(nums1)-1 ; i++{
        if nums1[i] + 1 == nums1[i+1]{
            count++     
        }else{
            if max < count{
                max = count 
            }
            count =1
        }
        if count > max {//如果中途没断 ，就应该赋值给max就好了
            max = count
        }
    }
    return max
}


//还有方法二：


/*
核心思想：利用哈希表快速判断数字是否存在，并且只从连续序列的起点开始遍历，避免重复计算。

具体步骤如下：
1. 构建哈希表：将数组中的所有数字存入 map，实现去重，并支持平均 \(O(1)\) 时间复杂度的元素查找。
2. 判断序列起点：遍历 map 中的数字 num，如果 num-1 不存在，说明 num 是某个连续序列的起点；否则跳过，避免重复遍历。
3. 寻找连续序列：从起点 num 开始，不断检查 num+1、num+2 等数字是否存在，并统计当前连续序列的长度。
4. 更新最大长度：每找到一个连续序列，就更新最长连续序列的长度，最终返回最大值。

func longestConsecutive(nums []int) int{
	var map1 = make(map[int]bool)
		for _ , value :=range nums{
			map1[value] = true
	    }

	var max int = 0
	
	for  value := range map1 {
		if !map1[value-1] {
			current := value 
			count := 1
		    for map1[current+1]{//这个就和while(1)有点像啊
			    current ++
			    count ++
		    }
		    if count > max{
			max = count
		    }
	    }
    }
	return max	
}


*/

