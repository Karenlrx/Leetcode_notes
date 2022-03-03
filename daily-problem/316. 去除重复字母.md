#### [6. 去除重复字母](https://leetcode-cn.com/problems/remove-duplicate-letters/)

给你一个字符串 `s` ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 **返回结果的字典序最小**（要求不能打乱其他字符的相对位置）。

```
示例 1：
输入：s = "bcabc"
输出："abc"

示例 2：
输入：s = "cbacdcbc"
输出："acdb"

提示：
1 <= s.length <= 104
s 由小写英文字母组成

```

#### 解题思路

- 建立一个字典。其中 key 为 字符 c，value 为其出现的剩余次数。
- 从左往右遍历字符串，每次遍历到一个字符，其剩余出现次数 - 1。
- 对于每一个字符，如果其对应的剩余出现次数大于 1，我们可以选择丢弃（也可以选择不丢弃），否则不可以丢弃。
- 是否丢弃的标准和上面题目类似。如果栈中相邻的元素字典序更大，那么我们选择丢弃相邻的栈中的元素。

**注意：**

- 在考虑字符 s[i]时，如果它已经存在于栈中，则不能加入字符 s[i]。为此，需要记录每个字符是否出现在栈中。
- golang的for i, ch :=range s中，ch类型为int32类型（ASCII码），所以需要遍历时构建变量ch := s[i]存储每个字符。

**代码演示（Golang）**

```go
func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}
	//stack存每个ch
	stack := []byte{}
	//inStack判断每个ch在不在stack里面，在返回true
	inStack := [26]bool{}
	for i,_ := range s {
		ch := s[i]
		//如果没在stack里面
		if !inStack[ch-'a'] {
			//如果栈中有元素&&栈中相邻的元素字典序更大&&栈顶元素在接下来的字符串还有（不存在丢弃就没法输出）
			//那么我们选择丢弃栈顶的元素。
			for len(stack) > 0 && ch < stack[len(stack)-1] && count[stack[len(stack)-1] - 'a'] > 0{
				//注意需要先置栈顶元素对应的为false，再出栈
				//否则先出栈，栈顶元素会改变
				inStack[stack[len(stack)-1] - 'a'] = false
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		count[ch-'a']--
	}
	return string(stack)
}
```

> 时间复杂度：O(N)，其中 N 为字符串长度。代码中虽然有双重循环，但是每个字符至多只会入栈、出栈各一次。
>
> 空间复杂度：O(∣Σ∣)，其中 Σ 为字符集合，本题中字符均为小写字母，所以 ∣Σ∣=26。由于栈中的字符不能重复，因此栈中最多只能有 ∣Σ∣ 个字符，另外需要维护两个数组，分别记录每个字符是否出现在栈中以及每个字符的剩余数量。
>