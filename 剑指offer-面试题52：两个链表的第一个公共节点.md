### 原题链接：

https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

### 题目描述：

输入两个链表，找出他们的第一个公共节点，示例：![img](%E5%89%91%E6%8C%87offer-%E9%9D%A2%E8%AF%95%E9%A2%9852%EF%BC%9A%E4%B8%A4%E4%B8%AA%E9%93%BE%E8%A1%A8%E7%9A%84%E7%AC%AC%E4%B8%80%E4%B8%AA%E5%85%AC%E5%85%B1%E8%8A%82%E7%82%B9.assets/160_example_1.png)

输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

### 解题思路：

#### 解法一：双指针法

使用两个指针 node1，node2 分别指向两个链表 headA，headB 的头结点，然后同时分别逐结点遍历，当 node1 到达链表 headA 的末尾时，重新定位到链表 headB 的头结点；当 node2 到达链表 headB 的末尾时，重新定位到链表 headA 的头结点。

这样，当它们相遇时，所指向的结点就是第一个公共结点。

#### 动画图解：

![leetcode_160](%E5%89%91%E6%8C%87offer-%E9%9D%A2%E8%AF%95%E9%A2%9852%EF%BC%9A%E4%B8%A4%E4%B8%AA%E9%93%BE%E8%A1%A8%E7%9A%84%E7%AC%AC%E4%B8%80%E4%B8%AA%E5%85%AC%E5%85%B1%E8%8A%82%E7%82%B9.assets/leetcode_160.gif)

代码演示：

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    nodeA, nodeB := headA, headB
    for nodeA != nodeB {
      if nodeA != nil {
        nodeA = nodeA.Next
      } else {
        nodeA = headB
      }
      if nodeB != nil {
        nodeB = nodeB.Next
      } else {
        nodeB = headA
      }
    }
    return nodeA
}

```

> ```
> 44 ms	7.9 MB
> ```

