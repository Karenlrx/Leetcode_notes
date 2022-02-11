#### [1365. 有多少小于当前数字的数字](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/)

给你一个数组 `nums`，对于其中每个元素 `nums[i]`，请你统计数组中比它小的所有数字的数目。

换而言之，对于每个 `nums[i] `你必须计算出有效的 j 的数量，其中 j 满足 `j != i` 且` nums[j]` < `nums[i] `。

以数组形式返回答案。

```
示例 1：

输入：nums = [8,1,2,2,3]
输出：[4,0,1,1,3]
解释： 
对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。 
对于 nums[1]=1 不存在比它小的数字。
对于 nums[2]=2 存在一个比它小的数字：（1）。 
对于 nums[3]=2 存在一个比它小的数字：（1）。 
对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
示例 2：

输入：nums = [6,5,4,8]
输出：[2,1,0,3]
示例 3：

输入：nums = [7,7,7,7]
输出：[0,0,0,0]


提示：

2 <= nums.length <= 500
0 <= nums[i] <= 100



```



#### 解题思路：

1. 排序+hash：

   排序完成后，全部的数据从小到大排序了，例如 [8,1,2,2,3] 变更为 [1,2,2,3,8] ，然后对有序的数组进行遍历，例如第一个 2 ，其 index 为 1 ，而比 2 小的元素刚好有 1 个，（后面第 2 个数字 2 ，比它小的也只有数字 1，if 判断那里退出，不进入 put 的逻辑）

代码演示：

```java
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class Solution {
    public int[] smallerNumbersThanCurrent(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>(); // 记录数字 nums[i] 有多少个比它小的数字
        int[] res = Arrays.copyOf(nums, nums.length);
        Arrays.sort(res);
        for (int i = 0; i < res.length; i++) {
            if (!map.containsKey(res[i])) { // 遇到了相同的数字，那么不需要更新该 number 的情况
                map.put(res[i], i);
            }
        }

        for (int i = 0; i < nums.length; i++) {
            res[i] = map.get(nums[i]);
        }

        return res;
    }
}    
```

> 时间复杂度：O(NlogN)，其中 N为数组的长度。排序需要 O(NlogN) 的时间，随后需要 O(N)时间来遍历。
>
> 空间复杂度：O(N)。因为要额外开辟一个数组。
>

2. 计数法：
   - 定义一个 计数数组(counter)，记录“比当前元素个数和当前下标小的元素个数”
   - 根据 `nums`数组 初始化 `counter`数组
   - 计算比 当前数字 小的 `nums`数组中的元素 的个数
   - 根据计数数组生成结果数组

代码演示：

```java
class Solution {
    public int[] smallerNumbersThanCurrent(int[] nums) {
        int[] res = new int[nums.length];
        //0 <= nums[i] <= 100
        int[] counter = new int[101];
        for (int val:nums) counter[val]++;
        //更新counter，计算比当前数字小的nums数组中的元素的个数
        for (int i = 1; i <= 100; i++) {
            counter[i] += counter[i - 1];
        }
        for (int i = 0; i < nums.length; i++) {
            // counter数组记录的是包含自身在内的比自己小的元素个数.
            // 因此,result[i] = counter[nums[i] - 1]
            res[i] = nums[i] == 0 ? 0 : counter[nums[i] - 1];
        }
        return res;
    }
}
```

> 时间复杂度：O(N + K)，其中 K 为值域大小。需要遍历两次原数组，同时遍历一次频次数组 counter 找出前缀和。
>
> 空间复杂度：O(K)。因为要额外开辟一个值域大小的数组。
>