#### [剑指 Offer 58 - II. 左旋转字符串](https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)

字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

```
示例 1：

输入: s = "abcdefg", k = 2
输出: "cdefgab"
示例 2：

输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"


限制：

1 <= k < s.length <= 10000
```

#### 解题思路

##### 方法一：暴力解

1. 记录前面需要翻转的，2. 记录后面不需要翻转的，3. 拼接。over~

##### 方法二：原地置换

曾经的考研真题，做过好多遍T_T

- 反转区间为前n的子串；
- 反转区间为n到末尾的子串；
- 反转整个字符串；

（图源：https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/solution/yuan-di-ju-bu-fan-zhuan-zheng-ti-fan-zhuan-xiang-j/）

![剑指Offer58-II.左旋转字符串.png](images/1599203229-TUcYHl-剑指Offer58-II.左旋转字符串.png)

#### 代码演示

```go
func reverseLeftWords(s string, n int) string {
	var sb,left strings.Builder
	for i := 0; i < n; i++ {
		left.WriteByte(s[i])
	}
	for j := n; j < len(s); j++ {
		sb.WriteByte(s[j])
	}
	sb.WriteString(left.String())
	return sb.String()
}
```



```go
func reverseLeftWords(s string, n int) string {
    b := []byte(s)
    // 1. 反转前n个字符
    // 2. 反转第n到end字符
    // 3. 反转整个字符
    reverse(b, 0, n-1)
    reverse(b, n, len(b)-1)
    reverse(b, 0, len(b)-1)
    return string(b)
}
// 切片是引用传递
func reverse(b []byte, left, right int){
    for left < right{
        b[left], b[right] = b[right],b[left]
        left++
        right--
    }
}

```

