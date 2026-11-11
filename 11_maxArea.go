/*给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。*/

//我先说明一下我的解法

//思路就是暴力穷举 这样会导致算法的计算复杂度是O(n²)，这就会导致一个算法超时的情况。

/*
func maxArea(height []int) int {
    var maxvolum int = 0
    var temp_maxvolum int
    for i := 0 ; i< len(height) - 1 ; i ++{
        for j := i + 1 ; j< len(height) ; j++{
            if (height[i]<=height[j]){
                temp_maxvolum = height[i] * (j-i)
                 if maxvolum < temp_maxvolum {
                    maxvolum = temp_maxvolum
                }
            }else{
                temp_maxvolum = height[j] * (j-i)
                 if maxvolum < temp_maxvolum {
                      maxvolum = temp_maxvolum
                 }
            }
        }
    }
    return maxvolum
}
*/

func maxArea(height []int) int {
    var maxvolum int = 0
	var left int = 0 
	var right int = len(height)-1

	for left < right {
		if height[left] < height[right] {
			if (height[left] * (right -left)) > maxvolum {
				maxvolum = (height[left] * (right -left))
			}
			left++
		}else{
			if (height[right] * (right -left)) > maxvolum {
				maxvolum = (height[right] * (right -left))
			}
			right--
		}
	}
    return maxvolum
}