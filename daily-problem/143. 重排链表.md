#### [143. 重排链表](https://leetcode-cn.com/problems/reorder-list/)

给定一个单链表·`L：L0→L1→…→Ln-1→Ln `，
将其重新排列后变为： `L0→Ln→L1→Ln-1→L2→Ln-2→…`

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

```
示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
```



#### 解题思路：

1. 找到原链表的中点（参考[「876. 链表的中间结点」](https://leetcode-cn.com/problems/middle-of-the-linked-list/)）。

   - 我们可以使用快慢指针来 O(N) 地找到链表的中间节点。快指针一次走两步，慢指针一次走一步，当快指针走到终点的话，慢指针会刚好到中点。如果节点个数是偶数的话，slow 走到的是左端点，利用这一点，我们可以把奇数和偶数的情况合并，不需要分开考虑。

2. 将原链表的右半端反转（参考[「206. 反转链表」](https://leetcode-cn.com/problems/reverse-linked-list/)）。

   - 我们可以使用迭代法实现链表的反转。

     ![img](https://pic.leetcode-cn.com/9ce26a709147ad9ce6152d604efc1cc19a33dc5d467ed2aae5bc68463fdd2888.gif)

3. 将原链表的两端合并。

   - 因为两链表长度相差不超过 1，因此直接合并即可。

举例为：

```
1 -> 2 -> 3 -> 4 -> 5 -> 6
第一步，将链表平均分成两半
1 -> 2 -> 3
4 -> 5 -> 6
    
第二步，将第二个链表逆序
1 -> 2 -> 3
6 -> 5 -> 4
    
第三步，依次连接两个链表
1 -> 6 -> 2 -> 5 -> 3 -> 4
```

**代码演示：**

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public void reorderList(ListNode head) {
        if (head == null || head.next == null || head.next.next == null) {
            return;
        }
        //找中点，链表分成两个
        ListNode slow = head;
        ListNode fast = head;
        //不能只用fast.next.next != null，因为当fast指向倒数第一个节点fast.next==null,所以fast.next.next会报空指针错
        while (fast.next != null && fast.next.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }

        ListNode newHead = slow.next;
        slow.next = null;
        
        //第二个链表倒置
        newHead = reverseList(newHead);
        
        //链表节点依次连接
        while (newHead != null) {
            ListNode temp = newHead.next;
            newHead.next = head.next;
            head.next = newHead;
            head = newHead.next;
            newHead = temp;
        }

    }

    public ListNode reverseList(ListNode head) {
        //申请节点，pre和 cur，pre指向null
        ListNode pre = head;
        ListNode cur = null;
        ListNode tmp = null;
        while(pre!=null) {
            //记录当前节点的下一个节点
            tmp = pre.next;
            //然后将pre指向cur 
            pre.next = cur;
            //pre和cur节点都前进一位
            cur = pre;
            pre = tmp;
        }
        return pre;
    }
}
```

> - 时间复杂度：O(N)，其中 N是链表中的节点数。
> - 空间复杂度：O(1)。