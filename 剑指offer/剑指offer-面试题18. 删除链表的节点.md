#### 原题链接：

https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/



#### 题目描述：

给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

```
示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.


说明：

题目保证链表中节点的值互不相同
若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
```



#### 解题思路：

简单的链表删除节点操作，注意记录并返回头节点。

注意：删除节点一般都要操作前驱指针，那在遍历的过程中记录前驱指针就非常有必要。

代码演示：

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    if head.Val==val{
        return head.Next
    }
    //记录头节点
    pre:=head   
    for head.Next.Val!=val{
        head=head.Next
    }
    head.Next=head.Next.Next
    return pre
}
```

> 时间复杂度：O(n)  空间复杂度：O(1)
>
> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.9 MB, 在所有 Go 提交中击败了100.00%的用户