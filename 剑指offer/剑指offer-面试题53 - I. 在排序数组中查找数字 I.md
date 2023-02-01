

#### 原题链接：

https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/



#### 题目描述：

统计一个数字在排序数组中出现的次数。

**示例 :**

```
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
```



#### 解题思路：

利用折半查找到target，然后找到target的左右边界，则target数量为rightIndex - leftIndex + 1，如图：

![image-20220525172600095](images/image-20220525172600095.png)

![image-20220525172636103](images/image-20220525172636103.png)

**代码演示：**

```go
func search(nums []int, target int) int {
	leftIndex := findIndex(nums, target, true)
	rightIndex := findIndex(nums, target, false)
	if leftIndex < 0 {
		return 0
	}
	return rightIndex - leftIndex + 1
}

// 找左右index
func findIndex(nums []int, target int, isLeft bool) int {
	var (
		index = -1
		mid   = 0
		n     = len(nums)
		low   = 0
		high  = n - 1
	)
	for low <= high {
		mid = (low + high) / 2
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		}
		if target == nums[mid] {
			if isLeft {
				high = mid - 1
			} else {
				low = mid + 1
			}
			index = mid
		}
	}
	return index
}

```

> 执行用时：4 ms, 在所有 Go 提交中击败了95.73%的用户
>
> 内存消耗：3.8 MB, 在所有 Go 提交中击败了100.00%的用户

```go
func search(nums []int, target int) int {
    if len(nums)==0{
        return 0
    }
    low , high := 0 , len(nums) - 1
     mid := (low + high ) / 2
         count :=0
     for low <=  high {
         if nums[mid] < target {
             low = mid + 1
         }else if nums[mid] > target {
             high = mid  - 1
         }
         mid = (low + high ) / 2
        if nums[mid] == target {
            count++
            break
        }
     }

    for right := mid+1; right < len(nums) && nums[right] == target ; right++{
        count++
    }
    for left:= mid-1 ; left >= 0 && nums[left] == target ; left--{
        count++
    }
    return count
}
```

> 执行用时 :8 ms, 在所有 Go 提交中击败了93.10%的用户
>
> 内存消耗 :4.1 MB, 在所有 Go 提交中击败了100.00%的用户