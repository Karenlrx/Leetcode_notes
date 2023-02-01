#### [剑指 Offer 57 - II. 和为s的连续正数序列](https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/)

```
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

 

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5
```



#### 解题思路

设一个连续的正整数序列。构建滑动窗口（双指针），因为是从小到大排列，滑动窗口只会向右滑动的特性完美符合。

（注意：本题并未说不同序列之间不可使用相同的数字）。

**算法流程：**

1. 初始化左边界`i=1`，右边界`j=2`，每次判断元素和sum与target关系：

    - 如果`sum < target`：向右移动右边界`j=j+1`，并更新sum；
    - 如果`sum > target`：说明以i为基数的数组不存在，更新`sum=sum-i`，向右移动左边界`i=i+1`；
    - 如果`sum=target`：说明找到数组，记录该数组，并更新`sum=sum-i`，向右移动左边界`i=i+1`。

    > 一个target=9的求解示例：

    ![Picture2.png](images/1611495306-LsrxgS-Picture2-20230201094856391.png)



**代码演示：**

观察本文的算法流程发现，当 `s=target 和 s>target` 的移动边界操作相同，因此可以合并，代码如下所示。
```go
func findContinuousSequence(target int) [][]int {
	i := 1
	j := 2
	sum := 3
	results := make([][]int, 0)
	for i < j {
        // 如果相等，得到答案，将其添加到二维数组中
		if sum == target {
			res := make([]int, 0, j-i+1)
			for k := i; k <= j; k++ {
				res = append(res, k)
			}
			results = append(results, res)
		}
        // 移动右边界，增加sum
		if sum < target {
			j++
			sum += j
		} else {
            // 移动做编辑啊，减少sum
			sum -= i
			i++
		}
	}
	return results
}
```

> 时间复杂度 O(N) ： 其中 N=target ；连续整数序列至少有两个数字，而 i<j 恒成立，因此至多循环 target 次（ i , j 都移动到 target/2），使用 O(N) 时间；当 i=1 时，达到最大序列长度-1+根号下(1+8*sum)/2 ，考虑到解的稀疏性，将列表构建时间简化考虑为 O(1) ；
>
> 空间复杂度 O(1) ： 变量 i, j , s 使用常数大小的额外空间。