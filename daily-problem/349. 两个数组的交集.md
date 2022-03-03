#### [349. 两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays/)

给定两个数组，编写一个函数来计算它们的交集。

```
示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]


说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
```

#### 解题思路：

1. 将`nums1`放入`set1`中（排除重复元素）。
2. 判断`nums2`中有没有set1中的元素，有则放入`set2`中。
3. 将HashSet转换为int[]数组
   - set遍历（`for (int num : set2) res[index++] = num;`）
   - 使用`set.toArray()`方法转换为Integer[]数组，再转换为int[]数组

代码演示：

```java
import java.util.HashSet;
class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        HashSet<Integer> set1 = new HashSet<>();
        HashSet<Integer> set2 = new HashSet<>();

        for (int i = 0; i < nums1.length; i++) {
            set1.add(nums1[i]);
        }
        for (int j =0; j < nums2.length;j++) {
            if (set1.contains(nums2[j])) set2.add(nums2[j]);
        }
        //1. 利用set的遍历
        int[] res = new int[set2.size()];
        int index = 0;
        for (int num : set2) res[index++] = num;
        //2.使用toArray()方法； 缺点：只能转换成Integer[]数组
//        Integer[] tmp = set2.toArray(new Integer[set2.size()]);
//        int[] res = new int[tmp.length];
//        for (int k = 0; k < tmp.length; k++) res[k] = tmp[k].intValue();
        return res;
    }
}
```

