

#### 原题链接：

https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

#### 题目描述：

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）

```
示例 ：

输入：head = [1,3,2]
输出：[2,3,1]

限制：
0 <= 链表长度 <= 10000
```



#### 解题思路：

**解法一：递归法**

- 指针移动到尾部，递归操作：tmp := append(tmp, head.Val)

- 递归出口：head.next == nil

  

**代码演示：**

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 递归
//执行用时 :92 ms, 在所有 Go 提交中击败了8.06%的用户
//内存消耗 :73.8 MB, 在所有 Go 提交中击败了100.00%的用户
func reversePrint(head *ListNode) []int {
	if head == nil{
		return nil
	}
    var tmp []int
    tmp = append(tmp , reversePrint(head.Next)...)
    tmp = append(tmp , head.Val)
	return tmp
}

//一行代码的递归
//执行用时 :4 ms, 在所有 Go 提交中击败了73.19%的用户
//内存消耗 :4.6 MB, 在所有 Go 提交中击败了100.00%的用户
//不过不清楚为啥第一种递归内存消耗非常大?
func reversePrint(head *ListNode) []int {
    if head == nil {
        return nil 
    }
    return append(reversePrint(head.Next), head.Val)
}

```



解法二：迭代法

将链表依次存在切片中，再交换对应元素

```go
func reversePrint(head *ListNode) []int {
	if head == nil{
		return nil
	}
    var slice []int
    for head != nil {
        slice = append(slice , head.Val)
        head = head.Next
    }
    for i := 0 ; i < len(slice)/2 ; i++ {
        slice[i] , slice[len(slice) - i-  1] = slice[len(slice) - i-  1] , slice[i]
    }
    return slice
}    
```

> 执行用时 :4 ms, 在所有 Go 提交中击败了73.19%的用户
>
> 内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户



解法三：

与迭代法类似，不过不是将数组存放到切片中，而是先利用count统计切片长度，再对切片进行赋值操作

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }

    count := 0
    newHead := head
    for head != nil {
        count++
        head = head.Next
    }

    slice := make([]int, count)
    i := 0
    for newHead != nil {
        slice[count-i-1] = newHead.Val
        i++
        newHead = newHead.Next 
    }

    return slice
}
```

> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.8 MB, 在所有 Go 提交中击败了100.00%的用户