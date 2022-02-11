## [106. 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

#### 题目描述

根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

```
    3
   / \
  9  20
    /  \
   15   7
```



#### 解题思路

1. 递归

   - 为了高效查找根节点元素在中序遍历数组中的下标，选择创建Hashmap来存储中序序列，即建立一个（元素，下标）键值对的哈希表。

   - 定义递归函数findRoot(int left, int right) 表示当前递归到中序序列中当前子树的左右边界，递归入口为findRoot(0, length-1) ：
     - 如果 left > right，说明子树为空，返回空节点。
     - 选择后序遍历的最后一个节点作为根节点。

   利用哈希表 O(1)查询当根节点在中序遍历中下标为 index。从left 到 index - 1 属于左子树，从 index + 1 到 right 属于右子树。

   根据后序遍历逻辑，递归创建右子树 和左子树。**注意这里有需要先创建右子树，再创建左子树的依赖关系。**可以理解为在后序遍历的数组中整个数组是先存储左子树的节点，再存储右子树的节点，最后存储根节点，如果按每次选择「后序遍历的最后一个节点」为根节点，则先被构造出来的应该为右子树。

   - 返回根节点 root。

   ```java
   /**
    * Definition for a binary tree node.
    * public class TreeNode {
    *     int val;
    *     TreeNode left;
    *     TreeNode right;
    *     TreeNode(int x) { val = x; }
    * }
    */
   public class Solution {
       int postIndex;
       int[] postorder;
       int[] inorder;
       Map<Integer, Integer> map = new HashMap<Integer, Integer>();
   
       public TreeNode buildTree(int[] inorder, int[] postorder) {
           //需要更新postorder节inorder
           this.postorder = postorder;
           this.inorder = inorder;
           int i = 0;
           postIndex = postorder.length - 1;
           //将中序序列存储到map中（inorder.val, i）
           for (int arr : inorder) {
               map.put(arr, i++);
           }
           return findRoot(0, postorder.length - 1);
       }
   
       public TreeNode findRoot(int left, int right) {
           if (left > right) return null;
           //root为postorder的最后一个节点
           int rootVal = postorder[postIndex];
           TreeNode root = new TreeNode(rootVal);
           //输出一个root，postIndex-1
           postIndex--;
           root.right = findRoot(map.get(rootVal) + 1, right);
           root.left = findRoot(left, map.get(rootVal) - 1);
   
           return root;
       }
   }
   ```
