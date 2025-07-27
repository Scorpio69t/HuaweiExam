# 1094. 拼车

## 描述

车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）

给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。

当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。

## 示例 1

输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false
解释：在位置3时，车上已经有2个乘客，再加上3个新乘客，总共5个乘客，超过了容量4。

## 示例 2

输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true
解释：在位置3时，车上已经有2个乘客，再加上3个新乘客，总共5个乘客，刚好等于容量5。

## 示例 3

输入：trips = [[2,1,5],[3,3,7],[1,2,4]], capacity = 6
输出：true
解释：在位置2时，车上已经有2个乘客，再加上1个新乘客，总共3个乘客；在位置3时，车上已经有3个乘客，再加上3个新乘客，总共6个乘客，刚好等于容量6。

## 提示

- 1 <= trips.length <= 1000
- trips[i].length == 3
- 1 <= numPassengersi <= 100
- 0 <= fromi < toi <= 1000
- 1 <= capacity <= 10^5

## 解题思路

### 方法一：差分数组（推荐）

**核心思想**：
- 使用差分数组记录每个位置上下车的人数变化
- 通过前缀和计算每个位置的乘客数量
- 检查是否在任何位置超过容量

**算法步骤**：
1. 创建差分数组，记录每个位置的人数变化
2. 在from位置增加乘客数，在to位置减少乘客数
3. 计算前缀和，得到每个位置的乘客总数
4. 检查是否所有位置的乘客数都不超过容量

**时间复杂度**：O(n + max(to))
**空间复杂度**：O(max(to))

### 方法二：排序 + 模拟

**核心思想**：
- 将所有上下车事件按位置排序
- 模拟车辆行驶过程，维护当前乘客数

**算法步骤**：
1. 将所有上车和下车事件提取出来
2. 按位置排序，相同位置下车优先于上车
3. 遍历所有事件，维护当前乘客数
4. 检查是否在任何时刻超过容量

**时间复杂度**：O(n log n)
**空间复杂度**：O(n)

### 方法三：优先队列

**核心思想**：
- 使用优先队列记录当前在车上的乘客
- 按下车位置排序，及时移除下车的乘客

**算法步骤**：
1. 按上车位置排序所有行程
2. 使用优先队列维护当前在车上的乘客
3. 处理每个行程时，先移除已下车的乘客
4. 检查上车后是否超过容量

**时间复杂度**：O(n log n)
**空间复杂度**：O(n)

## 代码实现

```go
// 方法一：差分数组
func carPooling1(trips [][]int, capacity int) bool {
    // 创建差分数组，最大位置为1000
    diff := make([]int, 1001)
    
    // 记录每个位置的人数变化
    for _, trip := range trips {
        passengers, from, to := trip[0], trip[1], trip[2]
        diff[from] += passengers
        diff[to] -= passengers
    }
    
    // 计算前缀和，检查是否超过容量
    current := 0
    for i := 0; i < 1001; i++ {
        current += diff[i]
        if current > capacity {
            return false
        }
    }
    
    return true
}
```

## 复杂度分析

- **时间复杂度**：O(n + max(to))，其中n是行程数量，max(to)是最大下车位置
- **空间复杂度**：O(max(to))，差分数组的大小

## 测试用例

```go
func main() {
    // 测试用例1
    trips1 := [][]int{{2, 1, 5}, {3, 3, 7}}
    capacity1 := 4
    fmt.Printf("测试用例1: trips=%v, capacity=%d, 结果=%t\n", trips1, capacity1, carPooling1(trips1, capacity1))
    
    // 测试用例2
    trips2 := [][]int{{2, 1, 5}, {3, 3, 7}}
    capacity2 := 5
    fmt.Printf("测试用例2: trips=%v, capacity=%d, 结果=%t\n", trips2, capacity2, carPooling1(trips2, capacity2))
    
    // 测试用例3
    trips3 := [][]int{{2, 1, 5}, {3, 3, 7}, {1, 2, 4}}
    capacity3 := 6
    fmt.Printf("测试用例3: trips=%v, capacity=%d, 结果=%t\n", trips3, capacity3, carPooling1(trips3, capacity3))
}
```
