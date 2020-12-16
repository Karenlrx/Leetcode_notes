#### [290. 单词规律](https://leetcode-cn.com/problems/word-pattern/)

给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

```
示例1:
输入: pattern = "abba", str = "dog cat cat dog"
输出: true

示例 2:
输入:pattern = "abba", str = "dog cat cat fish"
输出: false

示例 3:
输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false

示例 4:
输入: pattern = "abba", str = "dog dog dog dog"
输出: false

说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
```

####     解题思路

  在本题中，我们需要判断字符与字符串之间是否恰好一一对应。即任意一个字符都对应着唯一的字符串，任意一个字符串也只被唯一的一个字符对应。在集合论中，这种关系被称为「双向连接」。

- 如果是从pattern到s的单向连接的话：那么意味着"a"只对应着"dog",但"dog"都对应了谁无所谓，它可能除了"a"之外也对应着别人。因此，`abba`与"dog dog dog dog"符合单向连接的对应规律。在这里a只对应"dog"，但"dog"同时对应着"a"和"b"

- 而双向连接意味着，"a"只对应着"dog"，且"dog"也只对应着"a"，因此在双向连接的对应规律下，上述的`abba`与"dog dog dog dog"是匹配失败的，只有`abba`与"dog cat cat dog"这种可以匹配成功。

  想要解决本题，我们可以利用哈希表记录每一个字符对应的字符串，以及每一个字符串对应的字符。然后我们枚举每一对字符与字符串的配对过程，不断更新哈希表，如果发生了冲突，则说明给定的输入不满足双向连接关系。

1. 先将给定的s以空格为界分割成各个字符串，存放在数组words中。如：将"dog cat cat dog"分隔成["dog", "cat", "cat", "dog"]

2. 准备一个哈希表`ch2word`，哈希表的key记录着pattern中的字符，哈希表的value记录着arr数组中的字符串
3. 准备一个哈希表`word2ch`，哈希表的key记录着words数组中的字符串，哈希表的value记录着pattern中的字符

4. 之后遍历words数组，
   - `word2ch[word] > 0 && word2ch[word] != ch`表示同一个单词已有标识是否和当前标识匹配，如`pattern = "abaa", str = "dog cat cat fish" `对于s的第三个单词cat赋予的标识为b，但当前ch为'a'，return false。
   - `ch2word[ch] != "" && ch2word[ch] != word`表示标识ch已经赋予了对应单词，如`pattern = "abba", str = "dog cat cat fish"`，对于第四个单词fish，当前的标识a已经赋予给dog，与当前的单词不符，return false。
   - 否则，把pattern[i]和words[i]这二者的对应关系分别加入哈希表中。
     

**代码演示（Golang）**

```go
func wordPattern(pattern string, s string) bool {
	word2ch := map[string]byte{}
	ch2word := map[byte]string{}
	words := strings.Split(s, " ")
	//先判断单词数量与匹配的个数是否一致
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		ch := pattern[i]
		//word2ch[word] > 0 && word2ch[word] != ch表示同一个单词已有标识是否和当前标识匹配
		//如pattern = "abaa", str = "dog cat cat fish"  对于s的第三个单词cat赋予的标识为b，但当前ch为'a'，return false
		//ch2word[ch] != "" && ch2word[ch] != word表示标识ch已经赋予了对应单词
		//如pattern = "abba", str = "dog cat cat fish"，对于第四个单词fish，当前的标识a已经赋予给dog，与当前的单词不符，return false
		if word2ch[word] > 0 && word2ch[word] != ch || ch2word[ch] != "" && ch2word[ch] != word {
			return false
		}
		//为每个word标识ch
		//如：key:“dog”  value:"a"
		word2ch[word] = ch
		//为每个标识ch匹配word
		//如：key:“a”  value:"dog"
		ch2word[ch] = word
	}
	return true
}
```

> 时间复杂度：O(n+m)，其中 n 为 pattern 的长度，m 为 s 的长度。插入和查询哈希表的均摊时间复杂度均为 O(n+m)。每一个字符至多只被遍历一次。
>
> 空间复杂度：O(n+m)，其中 n 为 pattern 的长度，m 为 s 的长度。最坏情况下，我们需要存储 pattern 中的每一个字符和 str 中的每一个字符串。
>