#### [714. 买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

```
示例 1:

输入: prices = [1, 3, 2, 8, 4, 9], fee = 2

输出: 8

解释: 能够达到的最大利润:  
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

注意:
0 < prices.length <= 50000.
0 < prices[i] < 50000.
0 <= fee < 50000.
```

方法一：动态规划
思路与算法

考虑到「不能同时参与多笔交易」，因此每天交易结束后只可能存在手里有一支股票或者没有股票的状态。

- 定义状态 `dp[i][0] `表示第 i 天交易完后手里没有股票的最大利润，`dp[i][1] `表示第 i 天交易完后手里持有一支股票的最大利润（i 从 0 开始）。
- 考虑 `dp[i][0] `的转移方程，如果这一天交易完后手里没有股票，那么可能的转移状态为前一天已经没有股票，即 `dp[i−1][0]`，或者前一天结束的时候手里持有一支股票，即 `dp[i−1][1]`，这时候我们要将其卖出，并获得`prices[i]` 的收益，但需要支付 `fee` 的手续费。因此为了收益最大化，我们列出如下的转移方程：
  `dp[i][0]=max{dp[i−1][0],dp[i−1][1]+prices[i]−fee}`

- 再来按照同样的方式考虑 `dp[i][1] `按状态转移，那么可能的转移状态为前一天已经持有一支股票，即 `dp[i−1][1]`，或者前一天结束时还没有股票，即 `dp[i−1][0]`，这时候我们要将其买入，并减少 `prices[i]` 的收益。可以列出如下的转移方程：
  `dp[i][1]=max{dp[i−1][1],dp[i−1][0]−prices[i]}`

- 对于初始状态，根据状态定义我们可以知道第 0 天交易结束的时候有 `dp[0][0]=0` 以及 `dp[0][1]=−prices[0]`。

因此，我们只要从前往后依次计算状态即可。由于全部交易结束后，持有股票的收益一定低于不持有股票的收益，因此这时候 `dp[n−1][0] `的收益必然是大于 `dp[n−1][1]` 的，最后的答案即为 `dp[n−1][0]`。

**代码演示（Golang）**

```go
func maxProfit(prices []int, fee int) int {
    n := len(prices)
    dp := make([][2]int, n)
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }
    return dp[n-1][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

> 时间复杂度：O(N)
>
> 空间复杂度：O(N)

###### 空间优化

- 注意到在状态转移方程中，`dp[i][0] `和 `dp[i][1]` 只会从 `dp[i−1][0]` 和 `dp[i−1][1]` 转移而来
- 因此我们不必使用数组存储所有的状态。使用两个变量 `sell` 以及 `buy` 分别表示 `dp[..][0] `和 `dp[..][1]` 直接进行状态转移即可。

```go
func maxProfit(prices []int, fee int) int {
    n := len(prices)
    sell, buy := 0, -prices[0]
    for i := 1; i < n; i++ {
        sell = max(sell, buy+prices[i]-fee)
        buy = max(buy, sell-prices[i])
    }
    return sell
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

> 时间复杂度：O(N)
>
> 空间复杂度：O(1)