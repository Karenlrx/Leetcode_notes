#### 原题链接：

https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/



#### 题目描述：

实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。 

```
示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25


```



#### 解题思路：

解法一：递归法

代码演示：

```go
func myPow(x float64, n int) float64 {
    if n == 0{
        return 1
    }
    if n == 1{
        return x
    }
    if n<0{
        x = 1/x
        n = -n
    }
    temp := myPow(x , n/2)
    
    //区分n为奇数还是偶数
    if n%2 == 0{
        return temp*temp
    }
    return x*temp*temp
}
```

> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户

解法二：迭代法

代码演示

```go
func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1.0/x
        n = -n
    } else if n == 0 {
        return 1
    }
    ans := 1.0
    for n > 1 {
        if n & 1 == 1 {
            ans *= x
            n--
        } else {
            x = x*x
            n /= 2
        }
    }
    ans *= x
    return ans
}

```

> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户



解法三：快速幂（用位运算来实现）

- n & 1 — 取n的二进制数最低位 , n & 1 =1 ，n为奇数，=0 为偶数 ,相当于 n % 2==0
- n >> 1 —右移1位, 去掉 n 的二进制数最低位 , 相当于 n / 2

- 当 n 为偶数时
  - x <sup> n </sup> = x <sup>( n / 2 )</sup> * x <sup>( n / 2 )</sup>
  - n >> 1 , n右移 1 位后，x 自己乘自己，因为二进制每位的差距是平方关系
- 当 n 为奇数时
  - 需要再乘以多出来的一次，即 x <sup> n</sup> = x * x <sup>( n - 1 )</sup>
  - n - 1 , x 不更新，将 x 累乘到 ret



代码演示：

```go
func myPow(x float64, n int) float64 {
    f:=false
    if n<0{
        f = true
        n = -n
    }
    if n == 0 {
        return 1.0
    }
    
    ans:=float64(1)
    for n>1 {
        //快速幂
        if n%2 == 1 {
            ans*=x
        }
        x = x*x
        n = n>>1
    }
    ans*=x
    if f {
        return 1/ans
    }
    return ans
}

```

> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户