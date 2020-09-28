#### 原题链接：

https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/



#### 题目描述：

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

```
示例 ：

输入：s = "We are happy."
输出："We%20are%20happy."
```



解题思路：

- Golang中的string不支持修改字符的操作，所以构造切片存储string（如存在中文字符需要使用[]rune）
- 字符从原有的1个替换成3个，故需要统计空格个数，相应扩充字符串
- 切片属于引用类型，从底层存储来看实际是一个数据结构。如图所示：![选区_025](image/%E9%80%89%E5%8C%BA_025.png)

其中ptr存储的引用数组的首地址，len定义了切片的长度，cap定义切片的容量，故空间复杂度应为O(1)

**解法一：双指针法**

​	利用OriginCount记录原字符串长度，NewCount记录现有的字符串长度，则

- 从后往前遍历原字符串时，首先遇到的元素应放在现有字符串末尾
- 如遇到' '，则依次插入02%，并将currentCount自减



**代码演示：**

```go
func replaceSpace(s string) string {
    count := 0
    originCount := len(s)
    for i , _  := range  s {
        if s[i] == ' ' {
            count++
        }
    }
    for i := 0 ; i < 2*count ; i++ {
        s = s + " "
    } 
    str := []rune(s)
    currentCount := len(str) -1
    for i := originCount - 1  ; i >= 0 ; i-- {
        if str[i] == ' ' {
            str[currentCount] = '0'
            currentCount--
            str[currentCount] = '2'
            currentCount--
            str[currentCount] = '%'
            currentCount--
        }else {
            str[currentCount] = str[i]
            currentCount--
        }
    }
    return string(str)
}

```

> 时间复杂度：O(n)，   空间复杂度：O(1)
>
> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2.7 MB, 在所有 Go 提交中击败了100.00%的用户



**解法二：**

```go
func replaceSpace(s string) string {
    count := 0
    originCount := len(s) 
    for i , _  := range  s {
        if s[i] == ' ' {
            count++
        }
    }
    // for i := 0 ; i < 2*count ; i++ {
    //     s = s + " "
    // } 
    
    //不用扩充字符串，直接对切片进行append，切片 append 操作的本质就是对数组扩容，go 底层会创建一下新的数组 newArr(安装扩容后大小)，将 slice 原来包含的元素拷贝到新的数组 newArr，slice 重新引用到 newArr。旧的arr会被回收，其实本质上来说是一个道理，但不知道为啥相比于解法一会节省内存一点

    str := []rune(s)
    for i := 0; i < 2*count; i++ {
		str = append(str, ' ')
    }
    currentCount := len(str) -1
    for i := originCount - 1  ; i >= 0 ; i-- {
        if str[i] == ' ' {
            
            // 也可以使用golang内置的copy函数对切片进行拷贝
            copy(str[currentCount-2 : currentCount+1], []rune{'%', '2', '0'})
            currentCount = currentCount - 3
        }else {
            str[currentCount] = str[i]
            currentCount--
        }
    }
    return string(str)
}

```

> 时间复杂度：O(n)，   空间复杂度：O(1)
>
> 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
>
> 内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户