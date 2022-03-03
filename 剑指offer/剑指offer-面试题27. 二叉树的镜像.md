#### 原题链接：

https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/



#### 题目描述：

```
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
```



#### 解题思路：

二叉树的镜像即为原二叉树所有左右节点互换位置即可。
我们要做的就是：
1、判断当前节点是否是nil，如果是nil，可以直接返回 （边界条件）
2、如果不是nil，根节点保持不变，把左右子树节点互换，递归

代码演示：

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    //递归，左右节点交换，一行代码的形式，不用设置temp暂时存放左节点
    root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
    
    return root
}

```

> 时间复杂度 O(N) ： 其中 N 为二叉树的节点数量，建立二叉树镜像需要遍历树的所有节点，占用 O(N) 时间。
> 空间复杂度 O(N) ： 最差情况下（当二叉树退化为链表），递归时系统需使用 (N) 大小的栈空间。
>
> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户