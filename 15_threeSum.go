func threeSum(nums []int) [][]int {
    sort.Ints(nums) //这个默认对整型数组升序排序
    n := len(nums)	
    result := [][]int{}
    
/*
这题思路就是先排序，然后选取一个固定点i，然后针对该固定点，有一个双指针，从i+1到n-1检查 看一下是否满足sum==0
如果sum<0那么说明需要left++ 如果sum > 0 那么就需要right--
*/
    for i := 0 ; i < n-2 ; i++{
        if nums[i] > 0 {
            break
        }
        
        if i>0 && nums[i] == nums[i-1]{
            continue
        }

        left := i+1
        right := n-1
        for left < right{
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{nums[i],nums[left],nums[right]})
                for left < right && nums[left] == nums[left+1]{
                    left++
                }
                for left < right && nums[right] == nums[right-1]{
                    right--
                }
                left++
                right--
            }else if sum < 0{
                left++
            }else{
                right--
            }
        }
    }
    return result
}


