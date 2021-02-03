#### [480. 滑动窗口中位数](https://leetcode-cn.com/problems/sliding-window-median/)

中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

[2,3,4]，中位数是 3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 `nums`，有一个长度为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。

```
示例：

给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。

窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
 
 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
 
 提示：
你可以假设 k 始终有效，即：k 始终小于输入的非空数组的元素个数。
与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。
```

#### 解题思路（双堆）

​	维护滑动窗口内的数据，得到滑动窗口内的中位数。比较简单的做法就是将滑动窗口内的数字排序，滑动窗口内的和就等于最中间的数字，或者最中间的两个数字的平均值。**但这样做时间复杂度高，必定会超时。**

**算法思路：**

- 仅考虑中位数，其实我们只需要知道窗口内最中间的数字或者最中间的两个数字就可以了。可以想想如何将数字划分成两堆，一堆是小于中位数的数字集合，另一堆是大于等于中位数的数字集合。
- 利用两个堆来实现上面的这个想法，一个大根堆用于存放小于中位数的数字，另一个用小根堆来存放大于等于中位数的数字。因为需要最中间的数字，所以两个堆的大小彼此是有限制的。
  - 如果当k是偶数的时候，两个堆的大小相等，两个堆的堆顶的平均值就是当前滑动窗口内的中位数
  - 如果当k是奇数的时候，`小根堆的大小 = 大根堆的大小 + 1`.那么小根堆的堆顶就是当前滑动窗口内的中位数。

**算法实现：**

​	我们需要设计一个「数据结构」，用来维护滑动窗口，并且需要提供如下的三个接口：

- `insert(num)`：将一个数`num` 加入数据结构；

- `erase(num)`：将一个数 `num` 移出数据结构；

- `getMedian()`：返回当前数据结构中所有数的中位数。

第一个优先队列 `small` 是一个大根堆，它负责维护所有元素中较小的那一半；第二个优先队列 `large` 是一个小根堆，它负责维护所有元素中较大的那一半。具体地，如果当前需要维护的元素个数为 x，那么 `small` 中维护了 `⌈x/2⌉` 个元素，`large` 中维护了` ⌊x/2⌋ `个元素，其中`⌈y⌉` 和 `⌊y⌋` 分别表示将 y 向上取整和向下取整。也就是说：

**`small` 中的元素个数要么与 `large` 中的元素个数相同，要么比 `large` 中的元素个数恰好多 1 个。**

​	这样设计的好处在于：当二者包含的元素个数相同时，它们各自的堆顶元素的平均值即为中位数；而当 `small` 包含的元素多了一个时，`small` 的堆顶元素即为中位数。这样`getMedian()` 就设计完成了。

​	而对于 `insert(num)` 而言，如果当前两个优先队列都为空，那么根据元素个数的要求，我们必须将这个元素加入 `small`；如果 `small` 非空（显然不会存在 `small` 空而 `large `非空的情况），我们就可以将`num` 与 `small` 的堆顶元素 `top` 比较：

- 如果`num≤top`，我们就将其加入 `smal` 中；

- 如果 `num>top`，我们就将其加入 `large` 中。



​	在成功地加入元素 `num` 之后，两个优先队列的元素个数可能会变得不符合要求。由于我们只加入了一个元素，那么不符合要求的情况只能是下面的二者之一：

- `small` 比 `large` 的元素个数多了 2 个；

- `small` 比 `large` 的元素个数少了 1 个。



​	对于第一种情况，我们将 `small` 的堆顶元素放入 `large`；对于第二种情况，我们将`large` 的堆顶元素放入 `small`，这样就可以解决问题了，`insert(num)` 也就设计完成了。

​	然而对于`erase(num)` 而言，设计起来就不是那么容易了，因为我们知道，优先队列是不支持移出非堆顶元素这一操作的，因此我们可以考虑使用「延迟删除」的技巧，即：

> ​	当我们需要移出优先队列中的某个元素时，我们只将这个删除操作「记录」下来，而不去真的删除这个元素。当这个元素出现在 `small` 或者 `large` 的堆顶时，我们再去将其移出对应的优先队列。
>

​	

「延迟删除」使用到的辅助数据结构一般为哈希表 `delayed`，其中的每个键值对 `(num,freq)`，表示元素 `num` 还需要被删除 `freq` 次。「优先队列 + 延迟删除」有非常多种设计方式，体现在「延迟删除」的时机选择。在本题解中，我们使用一种比较容易编写代码的设计方式，即：

> ​	我们保证在任意操作 `insert(num)`，`erase(num)`，`getMedian() `完成之后（或者说任意操作开始之前），`small` 和`large` 的堆顶元素都是不需要被「延迟删除」的。这样设计的好处在于：我们无需更改 `getMedian()` 的设计，只需要略加修改 `insert(num)` 即可。
>



​	我们首先设计一个辅助函数 `prune(heap)`，它的作用很简单，就是对 `heap` 这个优先队列（`small` 或者`large` 之一），不断地弹出其需要被删除的堆顶元素，并且减少 `delayed` 中对应项的值。在 `prune(heap)` 完成之后，我们就可以保证**`heap` 的堆顶元素是不需要被「延迟删除」的**。

​	这样我们就可以在 `prune(heap)` 的基础上设计另一个辅助函数 `makeBalance()`，它的作用即为调整`small` 和 `large` 中的元素个数，使得二者的元素个数满足要求。由于有了 `erase(num)` 以及「延迟删除」，我们在将一个优先队列的堆顶元素放入另一个优先队列时，第一个优先队列的堆顶元素可能是需要删除的。因此我们就可以用`makeBalance()` 将 `prune(heap)` 封装起来，它的逻辑如下：

- 如果 `small` 和 `large` 中的元素个数满足要求，则不进行任何操作；
- 如果 `small` 比 `large` 的元素个数多了 2 个，那么我们我们将 `small` 的堆顶元素放入 `large`。此时 `small` 的对应元素可能是需要删除的，因此我们调用 `prune(small)`；

- 如果`small` 比 `large` 的元素个数少了 1 个，那么我们将 `large` 的堆顶元素放入 `small`。此时`large`的对应的元素可能是需要删除的，因此我们调用 `prune(large)`。

此时，我们只需要在原先 `insert(num)` 的设计的最后加上一步 `makeBalance() `即可。然而对于 `erase(num)`，我们还是需要进行一些思考的：

- 如果 `num` 与 `small` 和`large`的堆顶元素都不相同，那么 `numl` 是需要被「延迟删除」的，我们将其在哈希表中的值增加 1；

- 否则，例如 `num` 与 `small` 的堆顶元素相同，那么该元素是可以理解被删除的。虽然我们没有实现「立即删除」这个辅助函数，但只要我们将 `num` 在哈希表中的值增加 1，并且调用「延迟删除」的辅助函数 `prune(small)`，那么就相当于实现了「立即删除」的功能。

无论是「立即删除」还是「延迟删除」，其中一个优先队列中的元素个数发生了变化（减少了 1），因此我们还需要用 `makeBalance()` 调整元素的个数。

​	

此时，所有的接口都已经设计完成了。由于 `insert(num)` 和 `erase(num)` 的最后一步都是 `makeBalance()`，而 `makeBalance()` 的最后一步是 `prune(heap)`，因此我们就保证了任意操作完成之后，`small` 和`large`的堆顶元素都是不需要被「延迟删除」的。



**代码演示（Golang）**

```go
//要实现heap底下的Interface接口，查看源码得：
/*type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}*/
//所以需要实现sort.Interface接口 ==> 通过sort.IntSlice实现
//需要实现Push、Pop方法
type hp struct {
	sort.IntSlice	//sort.IntSlice是一个实现了sort.Interface接口的[]int
	size int
}
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { h.size++; heap.Push(h, v) }
func (h *hp) pop() int           { h.size--; return heap.Pop(h).(int) }

// prune()的作用是对small/large队列，不断弹出其需要的被删除的堆顶元素，并且减少 delayed 中对应项的值。
//在 prune() 完成之后，我们就可以保证 heap 的堆顶元素是不需要被「延迟删除」的。
func (h *hp) prune() {
	for h.Len() > 0 {
		//num为堆顶元素
		num := h.IntSlice[0]
		//small是大根堆，所以num取反
		if h == small {
			num = -num
		}
		//判断堆顶元素是否在延迟删除队列中，在的话就必须pop（因为num在small/large中实际已经被删除，但是如果不是堆顶元素不能随意删除）
		if d, ok := delayed[num]; ok {
			//说明不止一个待删除的num
			if d > 1 {
				delayed[num]--
			} else {
				//说明只有一个待删除的num，不能直接delayed[num]--，因为理论上上删除唯一的num后ok=false（即delay列表中不存在num）
				//而实际delayed[num]--还是会返回ok=true，所以需要删除num这个key
				delete(delayed, num)
			}
			heap.Pop(h)
		} else {
			break
		}
	}
}

var delayed map[int]int
var small, large *hp

//平衡small和large个数
func makeBalance() {
	// 调整 small 和 large 中的元素个数，使得二者的元素个数满足要求
	if small.size > large.size+1 { // small 比 large 元素多 2 个
		//因为small的每个数为原数取反，所以从small转移到large或从large转移到small，push的时候都需要取反
		large.push(-small.pop())
		small.prune() // small 堆顶元素被移除，需要进行 prune
	} else if small.size < large.size { // large 比 small 元素多 1 个
		small.push(-large.pop())
		large.prune() // large 堆顶元素被移除，需要进行 prune
	}
}

//insert为插入元素
func insert(num int) {
	//num属于small
	if small.Len() == 0 || num <= -small.IntSlice[0] {
		small.push(-num)
	} else {
		//num属于large
		large.push(num)
	}
	//插入后判断是否个数平衡
	makeBalance()
}

//清除元素
func erase(num int) {
	delayed[num]++
	if num <= -small.IntSlice[0] {
		small.size--
		//如果删除的元素num是堆顶元素
		if num == -small.IntSlice[0] {
			small.prune()
		}
	} else {
		large.size--
		if num == large.IntSlice[0] {
			large.prune()
		}
	}
	//清除后判断是否个数平衡
	makeBalance()
}

//取中位数
func getMedian(k int) float64 {
	//当窗口大小为奇数，返回small堆顶元素(注意取反)
	if k%2 != 0 {
		return float64(-small.IntSlice[0])
	}
	//窗口大小为偶数，返回(small与large堆顶元素和)/2
	return float64(-small.IntSlice[0]+large.IntSlice[0]) / 2
}

func medianSlidingWindow(nums []int, k int) []float64 {
	delayed = map[int]int{} // 哈希表，记录「延迟删除」的元素，key 为元素，value 为需要删除的次数
	small = &hp{}           // 大根堆，维护较小的一半元素
	large = &hp{}           // 小根堆，维护较大的一半元素

	//插入元素，开始计算滑动窗口内的中位数
	for _, num := range nums[:k] {
		insert(num)
	}
	n := len(nums)
	ans := make([]float64, 1, n-k+1)
	//初始化,记录从0到k-1位置的中位数
	ans[0] = getMedian(k)
	for i := k; i < n; i++ {
		//滑动窗口右移，添加最右侧元素
		insert(nums[i])
		//删除滑动窗口最左侧元素
		erase(nums[i-k])
		ans = append(ans, getMedian(k))
	}
	return ans
}
```

> 由于「延迟删除」的存在，`small` 比 `large` 在最坏情况下可能包含所有的 n 个元素，即没有一个元素被真正删除了。因此优先队列的大小是 O(n) 而不是 O(k) 的，其中 n 是数组 `nums` 的长度。
>
> 时间复杂度：O(n log n)。`insert(num)` 和 `erase(num)` 单次时间复杂度为 O(log n)，`getMedian()` 单次时间复杂度为 O(1)。因此总时间复杂度为 O(n log n)。
>
> 空间复杂度：O(n)。即为 `small`，`large` 和 `delayed` 需要使用的空间。
>