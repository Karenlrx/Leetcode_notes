#### [剑指 Offer 50. 第一个只出现一次的字符](https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/)

在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

```
示例 1:

输入：s = "abaccdeff"
输出：'b'
示例 2:

输入：s = "" 
输出：' '


限制：

0 <= s 的长度 <= 50000
```



### 解题思路

遍历两次字符串：

- 第一次，记录所有ch的次数（可以用哈希表，也可以用定长数组）；
- 第二次，找到第一个ch次数为1的值返回。

**代码演示：**

```go
func firstUniqChar(s string) byte {
    res := make([]int, 26)
    for _, ch := range s {
        res[ch-'a']++ 
    }
    for _, ch := range s {
        if res[ch-'a']==1 { 
            return byte(ch)
        }
    }
    return ' '
} 

// 似乎hashmap要慢很多
func firstUniqChar(s string) byte {
    hashMap := make(map[rune]int, 0)
    for _, ch := range s {
        val, := hashMap[ch]
        if ok {
            hashMap[ch] = val+1
        } else {
            hashMap[ch] = 1
        }
        
    }
    for _, ch := range s {
        if hashMap[ch] == 1 { 
            return byte(ch)
        }
    }
    return ' '
} 

```



TODO：golang Map与slice性能问题：