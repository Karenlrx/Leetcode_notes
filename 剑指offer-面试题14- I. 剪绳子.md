#### github无法显示latex公式解决办法：

Open the URL in Browser：[https://karenlrx.github.io/Leetcode_notes/%E5%89%91%E6%8C%87offer-%E9%9D%A2%E8%AF%95%E9%A2%9814-%20I.%20%E5%89%AA%E7%BB%B3%E5%AD%90.html](https://karenlrx.github.io/Leetcode_notes/剑指offer-面试题14- I. 剪绳子.html)



#### 原题链接：

https://leetcode-cn.com/problems/jian-sheng-zi-lcof/



#### 题目描述：

给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m] 。请问 k[0]*k[1]*...*k[m] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

```
示例 1：

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36

提示：2 <= n <= 58

```



#### 解题思路：

- 设将长度为 n 的绳子切为 a 段：
  $$
  n = n_1 + n_2 + ... + n_a
  $$

本题等价于求解：
$$
\max(n_1 \times n_2 \times ... \times n_a)
max(n 
1
	
 ×n 
2
	
 ×...×n 
a
	
 )
$$
以下数学推导总体分为两步：① 当所有绳段长度相等时，乘积最大。② 最优的绳段长度为 33 。

数学推导：
以下公式为“算术几何均值不等式” ，等号当且仅当 n~1~ = n~2~ = ... = n~an~ 时成立。
$$
\frac{n_1 + n_2 + ... + n_a}{a} \geq \sqrt[a]{n_1 n_2 ... n_a}
a
n 
1
	
 +n 
2
	
 +...+n 
a
	


 ≥ 
a

n 
1
	
 n 
2
	
 ...n 
a
$$
​	

==推论一： 将绳子以相等的长度等分为多段 ，得到的乘积最大。==

设将绳子按照 x长度等分为 a 段，即 n=ax ，则乘积为 x^ax^ 。观察以下公式，由于 n 为常数，因此当$x^{\frac{1}{x}}$取最大值时， 乘积达到最大值。

$$
x^a = x^{\frac{n}{x}} = (x^{\frac{1}{x}})^n
$$
根据分析，可将问题转化为求$ y = x^{\frac{1}{x}}$ 的极大值，因此对 x求导数。
$$
\begin{aligned} \ln y & = \frac{1}{x} \ln x & \text{取对数} \\ \frac{1}{y} \dot {y} & = \frac{1}{x^2} - \frac{1}{x^2} \ln x & \text{对 $x$ 求导} \\ & = \frac{1 - \ln x}{x^2} \\ \dot {y} & = \frac{1 - \ln x}{x^2} x^{\frac{1}{x}} & \text{整理得} \end{aligned}
$$
​	

令 $\dot {y} = 0 $，则$ 1 - \ln x = 0$ ，易得驻点为$x_0 = e \approx 2.7 $；根据以下公式，可知$ x_0$为极大值点。
$$
\dot {y} \begin{cases} > 0 & , x \in [- \infty, e) \\ < 0 & , x \in (e, \infty] \\ \end{cases}
$$
由于切分长度 x 必须为整数，最接近 e 的整数为 2 或 3 。如下式所示，代入 x = 2 和 x = 3 ，得出 x = 3 时，乘积达到最大。
$$
y(3) = 3^{1/3} \approx 1.44 \\ y(2) = 2^{1/2} \approx 1.41
$$


==推论二： 尽可能将绳子以长度 3 等分为多段时，乘积最大。==

切分规则：

- 最优： 3 。把绳子尽可能切为多个长度为 3 的片段，留下的最后一段绳子的长度可能为 0,1,2 三种情况。
- 次优： 2 。若最后一段绳子长度为 2 ；则保留，不再拆为 1+1。
- 最差： 1 。若最后一段绳子长度为 1 ；则应把一份 3 + 1 替换为 2 + 2，因为$ 2 \times 2 > 3 \times 1$。



算法流程：

- 当 $n \leq 3 $时，按照规则应不切分，但由于题目要求必须剪成 m>1 段，因此必须剪出一段长度为 1 的绳子，即返回 n - 1 。
- 当 n>3时，求 n 除以 3 的 整数部分 a 和余数部分 b （即 n = 3a + b ），并分为以下三种情况：
  - 当 b = 0 时，直接返回 $3^a$ ；
  - 当 b = 1 时，要将一个 1 + 3 转换为 2+2，因此返回 $3^{a-1} \times 4$；
  - 当 b = 2 时，返回$ 3^a \times 2$。



解法一：数学方法

代码演示：

```go
func cuttingRope(n int) int {
    //举例：
    //7   3 3 1  => 3 2 2
    // 6   3 3
    // 5   3 2 
    // 4   3 1  => 2 2  
    if n <= 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    // 能分成几个3
    parts := n / 3
    another := n % 3
    var result float64

    switch another {
        case 2:
        //使用Go内置math，func Pow(x, y float64) float64
            result = math.Pow(3, float64(parts))
            result *= 2
        case 1:
            result = math.Pow(3, float64(parts-1))
            result *= 4
        default:
            result = math.Pow(3, float64(parts))
    }
    return int(result)

}
```

> 时间复杂度 **O(1)** ： 仅有求整、求余、次方运算。
> 求整和求余运算：资料提到不超过机器数的整数可以看作是 O(1)；
> 幂运算：查阅资料，提到浮点取幂为 O(1) 。
> 空间复杂度 **O(1)** ： 变量 a 和 b 使用常数大小额外空间。
>
> 执行用时：**0 ms**
>
> 内存消耗：1.9 MB



解法二：动态规划

我们也可以使用动态规划，从已知值 F(2)F(2) 逐步迭代到目标值 F(n)F(n)，它是一种自底向上的方法。
建立一维动态数组 dp：

- 边界条件：dp[1] = dp[2] = 1，表示长度为 2 的绳子最大乘积为 1；
- 状态转移方程：dp[i] = max(dp[i], max((i - j) * j, j * dp[i - j]))，可以这样理解：

![14.jpg](image/82b25ac6bcb742f31e5202e4af993d98abfea6a0c385379b214440bbb84b9bb4-14-1586708357251.jpg)

代码演示：

```go
func cuttingRope(n int) int {
	dp := make([]int,n+1)
	dp[1] = 1 
    dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j :=1 ;  j <= i ;j++{
			dp[i] = max(dp[i], max((i - j) * j, j * dp[i - j]))// 递推公式
		}
	}
	return dp[n]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

```

> 时间复杂度：**O(n^2^**)    空间复杂度：**O(n)**
>
> 执行用时：0 ms
>
> 内存消耗：2.1 MB





解法三：优化后的动态规划

我们发现任何大于 33 的数都可以拆分为数字 1，2，3 的和，且它们对 3 的余数总是 0，1，2，因此我们可以仅用 dp[0]，dp[1]，dp[2] 表示所有大于 3 的值，这样空间复杂度可降到 O(1)。

![14.gif](image/3be12f435b2a0668eecd747c5d08188128fde7764b99116123b86880280f62ca-14.gif)

这样重复使用 dp 数组，只须一趟遍历即可完成，可使时间复杂度降到 O(N)。



**代码演示：**

```go
func cuttingRope(n int) int {
	var dp = [3]int{0,1,1} 
	for i := 3; i < n+1; i++ {
			dp[i%3] = Max(max(dp[(i - 1) % 3], i - 1),
                                            2 * max(dp[(i - 2) % 3], i - 2),
                                            3 * max(dp[(i - 3) % 3], i - 3))
		}
    return dp[n%3]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

//因为max只能接收2个参数，所以注意重新定义一个Max
func Max(i int, j int, k int) int {
	if i > j {
        if i >k{
            return i
        }
	}else{
        if j > k {
            return j
        }
    }
	return k
}
```

> 时间复杂度：**O(n**)    空间复杂度：**O(1)**
>
> 执行用时：0 ms
>
> 内存消耗：1.9 MB

