#### [387. 字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

```
 

示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2
```

#### 解题思路（哈希表存储频数）

对字符串进行两次遍历。

- 第一次遍历，使用哈希映射统计出字符串中每个字符出现的次数。
- 第二次遍历，只要遍历到了一个只出现一次的字符，那么就返回它的索引，否则在遍历结束后返回 -1。

**代码演示（Golang）**

```go
func firstUniqChar(s string) int {
    count := [26]int{}
    for _, ch := range s {
        count[ch-'a']++
    }
    for i, ch := range s {
        if count[ch-'a'] == 1 {
            return i
        }
    }
    return -1
}

```

