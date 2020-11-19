#### [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

```
示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
```



#### 解题思路

使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。

右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。

注意到以下性质：

左指针左边均为非零数；

右指针左边直到左指针处均为零。

因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。



**代码演示：**

```java
class Solution {
    public void moveZeroes(int[] nums) {
        int n = nums.length;
        //left指针指向非0元素后面的第一个0元素，即左边均为非零数
        int left = 0;
        for(int right = 0; right < n; right++) {
            //将非0元素往前移动，则left+1
            if (nums[right] != 0) {
                swap(nums, left, right);
                left++;
            }
            //否则right指向的为0，不交换
        }
    }

    public void swap(int[] nums, int left, int right) {
        int temp = nums[left];
        nums[left] = nums[right];
        nums[right] = temp;
    }
}
```

> - 时间复杂度：O*(*n)，其中 n 为序列长度。每个位置至多被遍历两次。
> - 空间复杂度：O(1)。只需要常数的空间存放若干变量。