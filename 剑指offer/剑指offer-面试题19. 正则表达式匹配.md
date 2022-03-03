#### 原题链接：

https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/



#### 题目描述：

实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

```
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。


```



#### 解题思路：

解法一：

1. 当p为空时，此时若s为空，返回true，若s不为空，说明不匹配，返回false
2. 当p只有1个字符时：
   2.1 s为空，不匹配返回false
   2.2 s不为空，满足s的第一个字符和p的第一个字符匹配，或者p的第一个字符为'.'。则s和p都往后移动一个字符，看看后面是否匹配。
   2.3 如果s不为空，且不满足2.2的条件，说明不匹配，返回false
   当p的字符数大于两个的时候
   3.1 p的第二个字符为*,此时有两种情况 eg：s="ab..." p="a*..."
   3.1.1 s的第一个字符和p的第一个字符相等或p的第一个字符为'.'
   此时有3种继续向后的方案
   1. s向后移一个字符，p向后移动两个字符 此时*的作用的是前面的字符a出现1次 即p变为"a..."
   2. s向后移一个字符，p不变 此时*的作用的是前面的字符a出现2次(当下匹配出现两次，*还保留) 即p变为"aa*...."
   3. s移动，p移动两个字符,此时*的作用的是前面的字符a出现0次。即跳过*和前面的一个字符
     3.1.2 不满足3.1.1的情况
     说明s和p的第一个字符不匹配，但是由于p的第二个字符*,可以把p的第一个字符去掉。
     此时s不变，p向后移动两个字符
     3.2 p的第二个字符不为*
     比较s和p的第一个字符是否匹配，是的话，就s和p都向后移动一个字符，比较后面的，反之不匹配。



```go
func isMatch(s string, p string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(p)+1)
	}
	return isMatchCore(s, p, 0, 0, dp)
}
func isMatchCore(s, p string, i, j int, dp [][]int) bool {
	if dp[i][j] == 1 {
		return true
	} else if dp[i][j] == -1 {
		return false
	}
	flag := false
    // p为空时
	if j == len(p) {
        // s，p均为空时，返回true
		if i == len(s) {
			flag = true
		} else {
            //当p为空,s不为空是，返回false
			flag = false
		}

	}
    // p为一个字符时
	if j == len(p)-1 {
        // s为空，返回false
		if i == len(s) {
			flag = false
		} else if s[i] == p[j] || p[j] == '.' {
			flag = isMatchCore(s, p, i+1, j+1, dp)
		} else {
			flag = false
		}
	}

    //一般情况
	if j < len(p)-1 {
		if p[j+1] == '*' {
            //s与p的前一个字符匹配时 || p的前一个字符为万能字符.
			if i <= len(s)-1 && (s[i] == p[j] || p[j] == '.') {
                // 以s="ab(...)"    p="a*(...)"为例
                //isMatchCore(s, p, i+1, j+2, dp)表示p="a(...)"，即*表示1个a
                //isMatchCore(s, p, i+1, j, dp)表示p="aa(...)"，即*表示2个a
                //isMatchCore(s, p, i, j+2, dp)表示p="(...)"，即*表示0个a
				flag = isMatchCore(s, p, i+1, j+2, dp) || isMatchCore(s, p, i+1, j, dp) || isMatchCore(s, p, i, j+2, dp)
			} else {
                // s与p的前一个字符不匹配时
				flag = isMatchCore(s, p, i, j+2, dp)
			}

        //一般情况，对应字符匹配或者p[j]为万能字符时
		} else if i <= len(s)-1 && (s[i] == p[j] || p[j] == '.') {
			flag = isMatchCore(s, p, i+1, j+1, dp)
		}

	}
	if flag == true {
		dp[i][j] = 1
	} else {
		dp[i][j] = -1

	}
	return flag

}


```

> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.9 MB, 在所有 Go 提交中击败了25.00%的用户



解法二：动态规划

思路：

- 如果 p.charAt(j) == s.charAt(i) : dp[i][j] = dp[i-1][j-1]；
- 如果 p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1]；
- 如果 p.charAt(j) == '*'：*
  - *如果 p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2] //in this case, a* only counts as empty
  - 如果 p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.'：
    - dp[i][j] = dp[i-1][j] //in this case, a* counts as multiple a
    - or dp[i][j] = dp[i][j-1] // in this case, a* counts as single a
    - or dp[i][j] = dp[i][j-2] // in this case, a* counts as empty



代码演示：

```go
func isMatch(s string, p string) bool {
    m, n := len(s),  len(p)
    //定义二维切片
   dp := make([][]bool, m+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, n+1)
	}
    dp[0][0]=true
    //判断当s为空的时候的情况
    for k :=1 ; k < n+1; k++{
        //p当前指向的字符如果为*且之前的为真，如s: ""  p: "a*"
        if p[k-1] =='*' && dp[0][k-2] {
            dp[0][k] = true
        }
    }
    for i :=1;i<=m;i++ {
            for j :=1; j<=n; j++ {
                //如果当前字符相等 || p当前的字符为万能字符
                if s[i-1]==p[j-1] || p[j-1]=='.' {
                 //dp[i][j当前的匹配情况与dp[i-1][j-1]情况相同
                    dp[i][j]=dp[i-1][j-1]
                }else if p[j-1]=='*' {
                    if s[i-1]!=p[j-2] && p[j-2]!='.' {
                        dp[i][j]=dp[i][j-2]
                    }else {
                        //dp[i][j] = dp[i-1][j] // 2个字符匹配的情况	
						// dp[i][j] = dp[i][j-1] // 单个字符匹配的情况
						// dp[i][j] = dp[i][j-2] // 没有匹配的情况
                        dp[i][j]=dp[i][j-1] || dp[i][j-2] || dp[i-1][j]
                    }
                }
            }
        }
    return dp[m][n]
}
    
```

> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.4 MB, 在所有 Go 提交中击败了100.00%的用户