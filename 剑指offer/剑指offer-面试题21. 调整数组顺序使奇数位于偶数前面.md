#### 原题链接：

https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/



#### 题目描述:

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。 

```
示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4] 
注：[3,1,2,4] 也是正确的答案之一。


提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10000

```



#### 解题思路：

双指针法（类似二路快排）

首尾双指针

- 定义头指针 left ，尾指针 right 。
- left 一直往右移，直到它指向的值为偶数。
- right 一直往左移， 直到它指向的值为奇数。
- 交换 nums[left] 和 nums[right] 。
- 重复上述操作，直到 left == right。

动画图解：

![img](image/f25bd8d3c3fd5d30969be2954685a21f67e254a6487c6d9d27edf6589a0fca55.gif)

代码演示：

```go
func exchange(nums []int) []int {
    if len(nums) == 0 {
        return nums
    }
    left , right := 0 , len(nums)-1
    for left !=right {
        if nums[left] % 2 == 0 {
            temp := nums[left]
            nums[left] = nums[right]
            nums[right] = temp
            right--
        }else {
            left++
        }
    }
    return nums
}
```

> 执行用时 :24 ms, 在所有 Go 提交中击败了91.98%的用户
>
> 内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户