#### [888. 公平的糖果棒交换](https://leetcode-cn.com/problems/fair-candy-swap/)

爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 根糖果棒的大小，B[j] 是鲍勃拥有的第 j 根糖果棒的大小。

因为他们是朋友，所以他们想交换一根糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）

返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。

如果有多个答案，你可以返回其中任何一个。保证答案存在。

```
示例 1：

输入：A = [1,1], B = [2,2]
输出：[1,2]
示例 2：

输入：A = [1,2], B = [2,3]
输出：[1,2]
示例 3：

输入：A = [2], B = [1,3]
输出：[2,3]
示例 4：

输入：A = [1,2,5], B = [2,4]
输出：[5,4]


提示：

1 <= A.length <= 10000
1 <= B.length <= 10000
1 <= A[i] <= 100000
1 <= B[i] <= 100000
保证爱丽丝与鲍勃的糖果总量不同。
答案肯定存在。
```

#### 解题思路

记爱丽丝的糖果棒的总大小为 `sumA`，鲍勃的糖果棒的总大小为 `sumB`。设答案为 `{x,y}`，即爱丽丝的大小为 x 的糖果棒与鲍勃的大小为 y 的糖果棒交换，则有如下等式：

```
sumA − x + y = sumB + x − y
```

化简，得：

```
x = y + (sumA−sumB)/2	
```

即对于 B 中的任意一个数 y'，只要 A 中存在一个数 x'，满足 

```
x' = y' + (sumA−sumB)/2	
```

那么 `{x',y'}` 即为一组可行解。

为了快速查询 A 中是否存在某个数，我们可以先将 A 中的数字存入哈希表(set)中。然后遍历 B 序列中的数 y'，在哈希表中查询是否有对应的 x'。

#### Notice

**1. 关于golang struct的使用技巧：声明为声明为map[string]struct{}**

- 由于struct{}是空，不关心内容，这样map便改造为set。

- map可以通过“comma ok”机制来获取该key是否存在,例如`_, ok := map["key"]`,如果没有对应的值,ok为false。

- 可以通过定义成`map[string]struct{}`的形式,值不再占用内存。其值仅有两种状态，有或无。

- 如果定义的是map[string]bool，则结果有true、false或没有。

**2. 关于golang死循环和条件循环的return语句**

- 条件循环里需在跳出循环后声明默认返回值（因为可能没有符合条件的情况）
- 死循环在循环内部定义返回值即可，不需要在跳出循环后声明默认返回值（因为没有符合条件的情况不会跳出循环）

**代码演示（Golang）**

```go
func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	//hashA := make(map[int]struct{}, len(A))
    hashA := map[int]struct{}{}
	for _, v := range A {
		sumA += v
        hashA[v] = struct{}{}
	}
	for i := range B {
		sumB += B[i]
	}
	tmp := (sumA - sumB) / 2
	for i := range B {
		swapB := B[i]
		swapA := swapB + tmp
	    if  _,ok := hashA[swapA]; ok  {
			return []int{swapA,swapB}
		}
	}
    return nil
}
```

> 时间复杂度：O(n+m)，其中 n 是序列 A 的长度，m 是序列 B 的长度。
>
> 空间复杂度：O(n)，其中 n 是序列 A 的长度。我们需要建立一个和序列 A 等大的哈希表。
>