## [LCP 19. 秋叶收藏集](https://leetcode-cn.com/problems/UlBDOe/)

小扣出去秋游，途中收集了一些红叶和黄叶，他利用这些叶子初步整理了一份秋叶收藏集 leaves， 字符串 leaves 仅包含小写字符 r 和 y， 其中字符 r 表示一片红叶，字符 y 表示一片黄叶。
出于美观整齐的考虑，小扣想要将收藏集中树叶的排列调整成「红、黄、红」三部分。每部分树叶数量可以不相等，但均需大于等于 1。每次调整操作，小扣可以将一片红叶替换成黄叶或者将一片黄叶替换成红叶。请问小扣最少需要多少次调整操作才能将秋叶收藏集调整完毕。

```
示例 1：

输入：leaves = "rrryyyrryyyrr"

输出：2

解释：调整两次，将中间的两片红叶替换成黄叶，得到 "rrryyyyyyyyrr"

示例 2：

输入：leaves = "ryr"

输出：0

解释：已符合要求，不需要额外操作

提示：

3 <= leaves.length <= 10<sup>5</sup>
leaves 中只包含字符 'r' 和字符 'y'
```



由于我们想要将收藏集中树叶的排列调整成「红、黄、红」三部分，因此我们可以用 3 个状态分别表示其中的每一部分，即状态 0 和状态 2 分别表示前面和后面的红色部分，状态 1 表示黄色部分。

状态转移方程如下图所示：

![images-20201004000328674](images/images-20201004000328674.png)

代码演示：

```java
class Solution {
    public int minimumOperations(String leaves) {
        int n = leaves.length();
        int[][] f = new int[n][3];
        f[0][0] = leaves.charAt(0) == 'y' ? 1 : 0;
        f[0][1] = f[0][2] = f[1][2] = Integer.MAX_VALUE;
        for (int i = 1; i < n; ++i) {
            int isRed = leaves.charAt(i) == 'r' ? 1 : 0;
            int isYellow = leaves.charAt(i) == 'y' ? 1 : 0;
            f[i][0] = f[i - 1][0] + isYellow;
            f[i][1] = Math.min(f[i - 1][0], f[i - 1][1]) + isRed;
            if (i >= 2) {
                f[i][2] = Math.min(f[i - 1][1], f[i - 1][2]) + isYellow;
            }
        }
        return f[n - 1][2];
    }
}
```

> 时间复杂度：O(n)，其中 n 是字符串leaves 的长度。
>
> 空间复杂度：O(n)。



```java
class Solution {
    public int minimumOperations(String leaves) {
        int[] f = new int[3];
        char[] chars = leaves.toCharArray();
        int len = leaves.length();
        f[0] = chars[0] == 'y' ? 1 : 0;
        //注意初始化给f[1]、f[2]一个很大的赋值（因为不存在第一片叶子属于第二三部分）
        //但是Integer.MAX_VALUE为最大值，遍历时候可能对其+1操作导致越界，故取值为Integer.MAX_VALUE - 1
        f[1] = f[2] = Integer.MAX_VALUE - 1;
        System.out.println(f[1]);
        for (int i = 1; i < len; i++) {
            int isRed = chars[i] == 'r' ? 1 : 0;
            int isYellow = chars[i] == 'y' ? 1 : 0;
            f[2] = Math.min(f[1], f[2]) + isYellow;
            f[1] = Math.min(f[0], f[1]) + isRed;
            f[0] = f[0] + isYellow;
        }
        return f[2];
    }
}
```

> 时间复杂度：O(n)，其中 n 是字符串leaves 的长度。
>
> 空间复杂度：O(1)。使用「降维」优化，用三个变量代替状态数组，即可将空间复杂度降低到 O(1)。

