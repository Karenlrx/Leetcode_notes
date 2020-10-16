#### [977. 有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)

给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

```
示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]

示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]

提示：

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A 已按非递减顺序排序。
```



#### 解题思路：

1.**先平方后排序**

顺便复习一下各类排序算法，排序算法总结可参考[排序算法总结](https://itimetraveler.github.io/2017/07/18/%E5%85%AB%E5%A4%A7%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95%E6%80%BB%E7%BB%93%E4%B8%8Ejava%E5%AE%9E%E7%8E%B0/)  。

PS：等下周有空整理一个自己的排序算法总结吧T_T

**代码演示：**

```java
class Solution {
    public int[] sortedSquares(int[] A) {
        int[] ans = new int[A.length];
        for (int i = 0; i < A.length; ++i) {
            ans[i] = A[i] * A[i];
        }
        //直接调用Arrays.sort
        //Arrays.sort()看数组长度 小于47用插入排序； [47,286)用快速排序； 大于等于286：连续性不好用快速排序，连续性好用归并排序 
        //Collections.sort()归并排序或TimSort
        Arrays.sort(ans);
        return ans;
    }
}
```

> 时间复杂度：`O(nlogn)`，其中 n 是数组 A 的长度。
>
> 空间复杂度：`O(logn)`。除了存储答案的数组以外，我们需要`O(logn) `的栈空间进行排序。
>



2. **双指针**

使用两个指针分别指向位置 0 和 n−1，每次比较两个指针对应的数，选择较大的那个**逆序**放入答案并移动指针。

**代码演示：**

```java
class Solution {
    public int[] sortedSquares(int[] A) {
        int left = 0, right = A.length-1;
        int[] res = new int[A.length];
        int index = right;
        while (index >= 0) {
            //说明负数>正数
            if (A[left] + A[right] < 0) {
                res[index] = A[left] * A[left];
                left++;
            }else {
                res[index] = A[right] * A[right];
                right--;
            }
            index--;
        }
        return res;
    }
}
```

> 时间复杂度：`O(n)`，其中 n 是数组 A 的长度。
>
> 空间复杂度：`O(1)`。除了存储答案的数组以外，只需要常量的存储空间。