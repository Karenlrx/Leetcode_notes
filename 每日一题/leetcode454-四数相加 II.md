#### [454. 四数相加 II](https://leetcode-cn.com/problems/4sum-ii/)

给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组` (i, j, k, l)` ，使得` A[i] + B[j] + C[k] + D[l] = 0`。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 `0 ≤ N ≤ 500 `。所有整数的范围在 -2^28^ 到 2^28^ - 1 之间，最终结果不会超过 2^31^ - 1 。

```
例如:

输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:

1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
```

#### 解题思路

采用分为两组，HashMap 存一组，另一组和 HashMap 进行比对。
这样的话情况就可以分为三种：

- HashMap 存一个数组，如 A。然后计算三个数组之和，如 BCD。时间复杂度为：O(n)+O(n^3^)，得到 O(n^3^).
- HashMap 存三个数组之和，如 ABC。然后计算一个数组，如 D。时间复杂度为：O(n^3^)+O(n)，得到 O(n^3^).
- HashMap存两个数组之和，如AB。然后计算两个数组之和，如 CD。时间复杂度为：O(n^2^)+O(n^2^)，得到 O(n^2^).

根据第二点我们可以得出要存两个数组算两个数组。
我们以存 AB 两数组之和为例。首先求出 A 和 B 任意两数之和 sumAB，以 sumAB 为 key，sumAB 出现的次数为 value，存入 hashmap 中。
然后计算 C 和 D 中任意两数之和的相反数 sumCD，在 hashmap 中查找是否存在 key 为 sumCD。
算法时间复杂度为 O(n^2^)。

**代码演示（JAVA）**

```java
import java.util.HashMap;

class Solution {
    public int fourSumCount(int[] A, int[] B, int[] C, int[] D) {
        int N = A.length;
        if (N == 0 && B.length != N && C.length != N && D.length != N) return 0;
        HashMap<Integer, Integer> map = new HashMap<>();
        int res = 0;
        //HashMap存储AB的和sumAB
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                int sumAB = A[i] + B[j];
                //key为sumAB，value为的到sumAB的组合次数
                if (map.containsKey(sumAB)) {
                    map.put(sumAB, map.get(sumAB)+1);
                }else {
                    map.put(sumAB,1);
                }
            }
        }
        //根据HashMap判断有没有匹配的res
        for (int k = 0; k < N; k++) {
            for (int l = 0; l < N; l++) {
                int sumCD = C[k] + D[l];
                if (map.containsKey(-sumCD)) res += map.get(-sumCD); 
            }
        }
        return res;
    }
}
```

**代码演示（Golang）**

```
func fourSumCount(A []int, B []int, C []int, D []int) int {
	sumAB := make(map[int]int)
	res := 0
	for _, v := range A {
		for _, w := range B {
			sumAB[v+w]++
		}
	}
	for _, v := range C {
		for _, w := range D {
			res += sumAB[-(v+w)]
		}
	}
	return res
}

```

> 时间复杂度：O(n^2^)。我们使用了两次二重循环，时间复杂度均为 O(n^2^)。在循环中对哈希映射进行的修改以及查询操作的期望时间复杂度均为 O(1)，因此总时间复杂度为 O(n^2^)。
>
> 空间复杂度：O(n^2^)，即为哈希映射需要使用的空间。在最坏的情况下，A[i]+B[j]A[i]+B[j] 的值均不相同，因此值的个数为 n^2^ ，也就需要 O(n^2^) 的空间。
>