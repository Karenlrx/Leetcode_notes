#### [1030. 距离顺序排列矩阵单元格](https://leetcode-cn.com/problems/matrix-cells-in-distance-order/)

给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。

另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。

返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，`|r1 - r2| + |c1 - c2|`。（你可以按任何满足此条件的顺序返回答案。）

 

```
示例 1：

输入：R = 1, C = 2, r0 = 0, c0 = 0
输出：[[0,0],[0,1]]
解释：从 (r0, c0) 到其他单元格的距离为：[0,1]
示例 2：

输入：R = 2, C = 2, r0 = 0, c0 = 1
输出：[[0,1],[0,0],[1,1],[1,0]]
解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2]
[[0,1],[1,1],[0,0],[1,0]] 也会被视作正确答案。
示例 3：

输入：R = 2, C = 3, r0 = 1, c0 = 2
输出：[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2,2,3]
其他满足题目要求的答案也会被视为正确，例如 [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]]。


提示：

1 <= R <= 100
1 <= C <= 100
0 <= r0 < R
0 <= c0 < C
```

#### 解题思路

1. **直接排序**

先存储矩阵内所有的点，然后将其按照哈曼顿距离直接排序。

**代码演示**

```java
class Solution {
    public int[][] allCellsDistOrder(int R, int C, int r0, int c0) {
        int[][] res = new int[R * C][2];
        for (int i = 0; i < R; i++) {
            for (int j = 0; j < C; j++) {
                //t为二维数组的第t项
                int t = i * C + j;
                res[t][0] = i;
                res[t][1] = j;
            }
        }

        // Arrays.sort(res, new Comparator<int[]>() {
        //     public int compare(int[] arr1, int[] arr2) {
        //         return (Math.abs(arr1[0] - r0) + Math.abs(arr1[1] - c0)) - (Math.abs(arr2[0] - r0) + Math.abs(arr2[1] - c0));
        //     }
        // });
        //使用Lambda表达式
        Arrays.sort(res, (arr1, arr2) -> (Math.abs(arr1[0] - r0) + Math.abs(arr1[1] - c0)) - (Math.abs(arr2[0] - r0) + Math.abs(arr2[1] - c0)));

        return res;
    }
}
```

> 时间复杂度：O(RClog(RC))，存储所有点时间复杂度 O(RC)，排序时间复杂度 O(RClog(RC))。
>
> 空间复杂度：O(log(RC))，即为排序需要使用的栈空间，不考虑返回值的空间占用。



2. **桶排序**

方法一中排序的时间复杂度太高。实际在枚举所有点时，我们可以直接按照哈曼顿距离分桶。这样我们就可以实现线性的桶排序。

- 遍历所有坐标，按照距离的大小分组，每组的距离相等（即放入一个桶中）。
- 按照距离从小到大的原则，遍历所有桶，并输出结果。

**注意：**

- 此解法时间复杂度为 O(R*C)，理论上已达到最快可能
- 实际时间消耗会比预估要差，不同语言便利程度和优化不一，原因如下：
  - 桶的制作涉及大量容器的初始化和存取
  - 桶中要存储大量的坐标信息，不论是直接使用长度为 2 的小数组存储，还是用新的简单数据类，都会耗费很多时间




**代码演示**

```java
class Solution {
    public int[][] allCellsDistOrder(int R, int C, int r0, int c0) {
        //maxDist记录最长的距离，看分成多少个桶
        int maxDist = Math.max(r0, R - 1 - r0) + Math.max(c0, C - 1 - c0);
        List<List<int[]>> bucket = new ArrayList<List<int[]>>();
        for (int i = 0; i <= maxDist; i++) {
            //一共有maxDist+1个桶
            bucket.add(new ArrayList<int[]>());
        }

        for (int i = 0; i < R; i++) {
            for (int j = 0; j < C; j++) {
                //记录点（i,j）到点（r0,c0）的距离
                int d = dist(i, j, r0, c0);
                //相同距离的点放到一个桶中
                bucket.get(d).add(new int[]{i, j});
            }
        }
        int[][] res = new int[R * C][];
        int index = 0;
        for (int i = 0; i <= maxDist; i++) {
            //依次从每个桶中取元素
            for (int[] it : bucket.get(i)) {
                res[index++] = it;
            }
        }
        return res;
    }

    //定义曼哈顿距离
    public int dist(int r1, int c1, int r2, int c2) {
        return Math.abs(r1 - r2) + Math.abs(c1 - c2);
    }
}
```

> 时间复杂度：O(RC)，存储所有点时间复杂度 O(RC)，桶排序时间复杂度 O(RC)。
>
> 空间复杂度：O(RC)，需要存储矩阵内所有点。
>



**解法三：BFS**
可以把所有的坐标看作树的结点，距离相等的结点位于树的同一层
而对于每一层的结点，它们的距离 dist 可以分为行距离和列距离，且 `rowDist + colDist = dist `必然成立
使 rowDist 从 0 到 dist 递增，colDist 相应有不同的值，可以得到不同的坐标：
横坐标为：`r0 - rowDist 或 r0 + rowDist`
纵坐标为：`c0 - colDist 或 c0 + colDist`
注意特殊情况：rowDist 或 colDist 为 0 时，每组只有一个正确值
对步骤 3 中，所有在矩阵范围内的坐标进行记录
注意：

此解法不关心最大距离，只要步骤 4 中记录的结果达到 R * C 的数量就可以终止搜索。
理论上此解法并不比桶排序优秀，但是代码中极少创建额外的容器和对象，所以实际的运行效率不会太差

**代码演示**

```java
class Solution {
    public int[][] allCellsDistOrder(int R, int C, int r0, int c0) {
        int[][] res = new int[R * C][2];
        int dist = 0;
        int cnt = 0;
        int[] factor = {-1, 1};
        while (cnt < R * C) {
            for (int rowDist = 0; rowDist <= dist; rowDist++) {
                int colDist = dist - rowDist;
                for (int i = 0; i < 2; i++) {
                    int row = r0 + factor[i] * rowDist;
                    for (int j = 0; j < 2; j++) {
                        int col = c0 + factor[j] * colDist;
                        if (row >= 0 && row < R && col >= 0 && col < C) {
                            res[cnt][0] = row;
                            res[cnt][1] = col;
                            cnt++;
                        }
                        if (colDist == 0) break;
                    }
                    if (rowDist == 0) break;
                }
            }
            dist++;
        }

        return res;
    }

}
```

> 时间复杂度: O((R+C)<sup>2</sup>)，因为对每一种距离 dist，rowDist 都要进行从 0 开始递增到 dist 的遍历操作，而距离可能的最大值为 R + C
> 此解法时间复杂度大于 O(R * C) 的原因是：每种距离可能产生多个不在矩阵内的坐标，但搜索算法必须依次检查予以排除。
>
> 空间复杂度：O(1)，不考虑创建额外数组存储res。



**解法四：几何法（类 BFS）**
如果把矩阵当作二维直角坐标系中的图形，而且把所有不在矩阵内的点也考虑进来，那么所有到 (r0, c0) 点的“距离”相等的整数坐标有明显的规律：

![Snipaste_2020-03-04_23-51-22.png](https://pic.leetcode-cn.com/47aacf8273ec9c560510012f74be0fe5a617b7517d3b191d7f34ce8837d907ea-Snipaste_2020-03-04_23-51-22.png)

可以看到，它们的坐标都在一个正方形的边上（包括顶点），而且正方形的上下顶点 row 值为 r0，左右顶点 col 值为 c0。
这样，只要保证每次找到一个正方形的顶点，然后按照规律“画出”这个正方形即可，画图步骤如下:

保存 4 个向量标明画线的方向
出发点为 (r0 - 1, c0)
按照 1 中的向量指示方向画线，遇到一个正方形的顶点就更换为下一个向量（向左转 90°）
在上述的画线步骤中，不断检查线上的整数坐标，如果符合要求就进行记录。

**注意：**

顶点的判断方法有两组，分别对应和 r0 或 c0 是否相等
对每个距离 dist 都要画出正方形检查，检查的点数量是 8 * dist，而最大距离可能是 R + C，所以时间复杂度为 O((R+C)<sup>2</sup>)
此解法代码中看似没有按照距离分层遍历，实际每个初始顶点的求解过程中已经包含了按照距离分层的想法，实际极其类似 BFS
此解法要检查的点理论上多于 BFS，尤其是 (r0, c0) 位于矩阵一角时会明显偏慢（最后要画很多很大的正方形）

**代码演示**

```java
class Solution {
    public int[][] allCellsDistOrder(int R, int C, int r0, int c0) {
        int[][] res = new int[R * C][2];
        res[0][0] = r0;
        res[0][1] = c0;
        int[] dr = {1, 1, -1, -1};
        int[] dc = {1, -1, -1, 1};
        int row = r0;
        int col = c0;
        var cnt = 1;
        while (cnt < R * C) {
            row--;
            for (int i = 0; i < 4; i++) {
                while ((i % 2 == 0 && row != r0) || (i % 2 != 0 && col != c0)) {
                    if (row >= 0 && row < R && col >= 0 && col < C) {
                        res[cnt][0] = row;
                        res[cnt][1] = col;
                        cnt++;
                    }
                    row += dr[i];
                    col += dc[i];
                }
            }
        }
        return res;
    }
}
```

> 时间复杂度：O((R+C) <sup>2</sup>)，我们需要遍历矩阵内所有点，同时也会遍历部分超过矩阵部分的点。在最坏情况下，给定的单元格位于矩阵的一个角，例如 (0,0)(0,0)，此时最大的曼哈顿距离为 `R+C−2`，需要遍历的点数为 `2(R+C−2)(R+C−1)+1`，因此时间复杂度为 O((R+C) <sup>2</sup>)。
>
> 空间复杂度：O(1)，不考虑返回值的空间占用。
>