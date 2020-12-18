#### [389. 找不同](https://leetcode-cn.com/problems/find-the-difference/)

给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。

```
示例 1：

输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。
示例 2：

输入：s = "", t = "y"
输出："y"
示例 3：

输入：s = "a", t = "aa"
输出："a"
示例 4：

输入：s = "ae", t = "aea"
输出："a"

提示：

0 <= s.length <= 1000
t.length == s.length + 1
s 和 t 只包含小写字母
```

#### 解题思路（计数）

- 先遍历字符串 s，对其中的每个字符都将计数值加 1
- 然后遍历字符串 t，对其中的每个字符都将计数值减 1。
- 当发现某个字符计数值为负数时，说明该字符在字符串 t 中出现的次数大于在字符串 s 中出现的次数，因此该字符为被添加的字符。

**代码演示（Golang）**

```go
func findTheDifference(s string, t string) byte {
	var count [26]int
	for _, v := range s {
		count[v-'a']++
	}
	for _, v := range t {
		count[v-'a']--
		if count[v-'a'] < 0 {
			return byte(v)
		}
	}
    return ' '
}
```

> 时间复杂度：O(N)，其中 N 为字符串的长度。
>
> 空间复杂度：O(∣Σ∣)，其中 Σ 是字符集，这道题中字符串只包含小写字母，∣Σ∣=26。需要使用数组对每个字符计数。
>



#### 方法二（ASCII码求和）

- 将字符串 s 中每个字符的 ASCII 码的值求和，得到A<sub>s</sub> 
- 对字符串 t 同样的方法得到 A<sub>t</sub> 
- 两者的差值A<sub>t</sub> - A<sub>s</sub> 即代表了被添加的字符

**代码演示（Golang）**

```
func findTheDifference(s, t string) byte {
    sum := 0
    for _, v := range s {
        sum -= int(v)
    }
    for _, v := range t {
        sum += int(v)
    }
    return byte(sum)
}
```

> 时间复杂度：O(N)。
>
> 空间复杂度：O(1)。

### 方法三（位运算）

如果将两个字符串拼接成一个字符串，则问题转换成求字符串中出现奇数次的字符。类似于「[136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/)」，我们使用位运算的技巧解决本题。

```go
func findTheDifference(s, t string) (diff byte) {
    for i := range s {
    //出现偶数次的字符异或结果为0，所以最后将所有字符异或后，即最终答案
        diff ^= s[i] ^ t[i]
    }
    return diff ^ t[len(t)-1]
}
```

> 时间复杂度：O(N)。
>
> 空间复杂度：O(1)。