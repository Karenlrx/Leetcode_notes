####  原题链接：

https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/



#### 题目描述：

定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

```
示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

限制：
0 <= 节点个数 <= 5000
```



#### 解题思路：

- 定义两个指针： pre 和 cur ；pre 在前 cur在后。
- 每次让 pre 的 next 指向 cur ，实现一次局部反转
- 局部反转完成之后， pre 和 cur 同时往前移动一个位置
- 循环上述过程，直至 pre 到达链表尾部

![img](image/9ce26a709147ad9ce6152d604efc1cc19a33dc5d467ed2aae5bc68463fdd2888.gif)

代码演示：

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    //就地逆置
    var cur *ListNode
    pre := head
    for pre != nil{
        next := pre.Next
        pre.Next = cur
        cur = pre
        pre = next
    }
    return cur
}

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
```

> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.5 MB, 在所有 Go 提交中击败了100.00%的用户