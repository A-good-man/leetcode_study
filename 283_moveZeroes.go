func moveZeroes(nums []int)  {
    if nums == nil { //对空数组的操作
        return
    }
    var count =0 //记录一下需要已经是第几个非0变量了。需要放在nums的第几个位置。
    for i, value := range nums { 
        if value != 0{
            temp := nums[count]
            nums[count] = value
            nums[i] = temp
            count ++
        }
    }
}
/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

 

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]
 

提示:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
 

进阶：你能尽量减少完成的操作次数吗？


*/