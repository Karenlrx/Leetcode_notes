#### [1207. 独一无二的出现次数](https://leetcode-cn.com/problems/unique-number-of-occurrences/)

给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。

如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。

```
示例 1：
输入：arr = [1,2,2,1,1,3]
输出：true
解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。

示例 2：
输入：arr = [1,2]
输出：false

示例 3：
输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
输出：true

提示：
1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
```



#### 解题思路：

- 要先计算每个数出现的次数。后面的只需要判断这个出现次数的数组中元素是否有重复的即可。

- 把出现次数的数组放到集合set中，如果有重复的就会被替换掉，**在set集合中如果有相同的元素，就会存储失败，**返回false，每次存储的时候我们只要判断是否存储成功即可。



**代码演示：**

```java
import java.util.HashMap;
import java.util.HashSet;

class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < arr.length; i++) {
            map.put(arr[i],map.getOrDefault(arr[i], 0)+1);
        }
        HashSet<Integer> set = new HashSet<>();
        for (int val:map.values()) {
            if (!set.add(val)) return false;
        }
        return true;
    }
}
```

> 时间复杂度：O(N)，其中 N 为数组的长度。遍历原始数组需要 O(N)时间，而遍历中间过程产生的哈希表又需要 O(N) 的时间。
>
> 空间复杂度：O(N)。

