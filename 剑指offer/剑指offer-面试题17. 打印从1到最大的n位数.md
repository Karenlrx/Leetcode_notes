#### 原题链接：

https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/



#### 题目描述：

输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

```
示例 1:
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]

说明：

用返回一个整数列表来代替打印
n 为正整数
```



#### 解题思路：

解法一：不考虑溢出

**代码演示：**

```go
func printNumbers(n int) []int {
    num := 1
    for n !=0 {
        num = 10*num
        n--
    }
    var res []int
    for i := 0; i< num -1; i++ {
        res = append(res,i+1)
    }
    return res
}
```

> 执行用时 :8 ms, 在所有 Go 提交中击败了93.60%的用户
>
> 内存消耗 :6.9 MB, 在所有 Go 提交中击败了100.00%的用户



解法二：考虑溢出：

考虑结果溢出问题，故转化为大数问题，打印结果应为字符串。

