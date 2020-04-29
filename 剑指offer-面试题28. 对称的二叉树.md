#### 原题链接：

https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/



#### 题目描述：

```
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false

限制：
0 <= 节点个数 <= 1000
```



#### 解题思路：

对称二叉树定义： 对于树中 任意两个对称节点 L 和 R ，一定有：

- L.val = R.val ：即此两对称节点值相等。

- L.left.val = R.right.val ：即 L 的 左子节点 和 R 的 右子节点 对称；

- L.right.val = R.left.val ：即 L 的 右子节点 和 R 的 左子节点 对称。

  根据以上规律，考虑从顶至底递归，判断每对节点是否对称，从而判断树是否为对称二叉树。


![Picture1.png](image/ebf894b723530a89cc9a1fe099f36c57c584d4987b080f625b33e228c0a02bec-Picture1.png)

**算法流程：**
**isSymmetric(root) ：**

特例处理： 若根节点 root 为空，则直接返回 true 。
返回值： 即 recur(root.left, root.right) ;

**recur(L, R) ：**

终止条件：

- 当 L 和 R 同时越过叶节点： 此树从顶至底的节点都对称，因此返回 true；
- 当 L 或 R 中只有一个越过叶节点： 此树不对称，因此返回 false；
- 当节点 L 值!=节点 R 值： 此树不对称，因此返回 false ；

递推工作：

- 判断两节点 L.left 和 R.right 是否对称，即 recur(L.left, R.right) ；
- 判断两节点 L.right 和 R.left 是否对称，即 recur(L.right, R.left) ；
- 返回值： 两对节点都对称时，才是对称树，因此用与逻辑符 && 连接。

动画图解：

![offer28](image/offer28.gif)



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
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return recur(root.Left , root.Right)
}

func recur(L *TreeNode , R *TreeNode) bool {
    //同时越过叶子节点返回的nil
    if L ==nil && R == nil {
        return true
    }
    //如只有一边返回nil或者对应值不相等，返回false
    if (L ==nil && R != nil) || (L != nil && R == nil) || L.Val != R.Val {
        return false
    }
	//判断两节点对应的左节点与右节点是否相等&&右节点月左节点是否相等
    return recur(L.Left , R.Right) && recur(L.Right , R.Left)
}

```

> 时间复杂度 O(N) ： 其中 NN 为二叉树的节点数量，每次执行 recur() 可以判断一对节点是否对称，因此最多调用 N/2 次 recur() 方法。
> 空间复杂度 O(N) ： 最差情况下（见下图），二叉树退化为链表，系统使用 O(N) 大小的栈空间。
>
> 执行用时 :4 ms, 在所有 Go 提交中击败了74.86%的用户
>
> 内存消耗 :2.9 MB, 在所有 Go 提交中击败了100.00%的用户