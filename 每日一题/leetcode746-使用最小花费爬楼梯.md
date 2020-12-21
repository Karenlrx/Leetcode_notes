#### [746. 使用最小花费爬楼梯](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)

数组的每个索引作为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

```
示例 1:

输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。
 示例 2:

输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。
注意：

cost 的长度将会在 [2, 1000]。
每一个 cost[i] 将会是一个Integer类型，范围为 [0, 999]。
```

#### 解题思路（动态规划）

到达第i级台阶的阶梯顶部的最小花费，有两个选择：

- 先付出最小总花费`dp[i-1]`到达第i级台阶（即第`i-1`级台阶的阶梯顶部），踏上第`i`级台阶需要再花费`cost[i]`，再迈一步到达第i级台阶的阶梯顶部，最小总花费为dp[i-1] + cost[i])`；

- 先付出最小总花费`dp[i-2]`到达第`i-1`级台阶（即第`i-2`级台阶的阶梯顶部），踏上第`i-1`级台阶需要再花费`cost[i-1]`，再迈两步跨过第`i`级台阶直接到达第`i`级台阶的阶梯顶部，最小总花费为`dp[i-2] + cost[i-1])`；

则`dp[i]`是上面这两个最小总花费中的最小值：`dp[i] = min(dp[i-1] + cost[i], dp[i-2] + cost[i-1])`。

最小总花费的初始值为：`dp[0]=dp[1]=0`

**代码演示（Golang）**

```go
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int,n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

```

> 时间复杂度：O(n)，其中 n 是数组 cost 的长度。需要依次计算每个dp 值，每个值的计算需要常数时间，因此总时间复杂度是 O(n)。
>
> 空间复杂度：O(n)，dp数组大小。



上述代码的时间复杂度和空间复杂度都是 O(n)。注意到当 `i≥2` 时，`dp[i]` 只和 `dp[i−1]` 与 `dp[i−2]` 有关，因此可以使用滚动数组的思想，将空间复杂度优化到 O(1)。

**代码演示（空间优化）**

```go
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    pre, cur := 0, 0
    for i := 2; i <= n; i++ {
        pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
    }
    return cur
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

```

> 时间复杂度：O(n)，其中 n 是数组 cost 的长度。需要依次计算每个dp 值，每个值的计算需要常数时间，因此总时间复杂度是 O(n)。
>
> 空间复杂度：O(1)。使用滚动数组的思想，只需要使用有限的额外空间。
>