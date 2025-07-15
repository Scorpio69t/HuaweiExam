# 42. 接雨水

## 📋 题目描述

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

## 🎯 示例

### 示例1
```
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，
在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
```

### 示例2
```
输入：height = [4,2,0,3,2,5]
输出：9
```

### 可视化展示
```
示例1的雨水分布：
   █
 █~█~█
 █~███~█
██~███~██
0102101321
总雨水量: 6

其中：█ 表示柱子，~ 表示雨水
```

## 🔍 解题思路

### 核心思想：木桶原理

接雨水的关键在于：**每个位置能接的雨水量 = min(左边最高柱子, 右边最高柱子) - 当前柱子高度**

### 算法流程

```mermaid
graph TD
    A[开始] --> B{选择算法}
    B -->|暴力法| C[对每个位置寻找左右最高]
    B -->|动态规划| D[预计算左右最高数组]
    B -->|双指针| E[用两个指针同时处理]
    B -->|单调栈| F[维护递减栈处理凹槽]
    
    C --> G[O(n²)时间复杂度]
    D --> H[O(n)时间，O(n)空间]
    E --> I[O(n)时间，O(1)空间]
    F --> J[O(n)时间，O(n)空间]
    
    G --> K[计算结果]
    H --> K
    I --> K
    J --> K
    K --> L[返回总雨水量]
```

## 🚀 算法实现

### 方法1：暴力法（时间复杂度O(n²)）

**思路**：对每个位置，分别寻找左边和右边的最高柱子。

```go
func trapBruteForce(height []int) int {
    for i := 1; i < n-1; i++ {
        // 找左边最高
        leftMax := 0
        for j := 0; j < i; j++ {
            leftMax = max(leftMax, height[j])
        }
        
        // 找右边最高
        rightMax := 0
        for j := i + 1; j < n; j++ {
            rightMax = max(rightMax, height[j])
        }
        
        // 计算当前位置雨水
        water := min(leftMax, rightMax) - height[i]
        if water > 0 {
            totalWater += water
        }
    }
}
```

### 方法2：动态规划（时间复杂度O(n)，空间复杂度O(n)）

**思路**：预计算每个位置的左边最高和右边最高，避免重复计算。

```go
func trapDP(height []int) int {
    // 预计算左边最高
    leftMax := make([]int, n)
    leftMax[0] = height[0]
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i])
    }
    
    // 预计算右边最高
    rightMax := make([]int, n)
    rightMax[n-1] = height[n-1]
    for i := n-2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i])
    }
    
    // 计算雨水
    for i := 1; i < n-1; i++ {
        water := min(leftMax[i], rightMax[i]) - height[i]
        if water > 0 {
            totalWater += water
        }
    }
}
```

### 方法3：双指针（时间复杂度O(n)，空间复杂度O(1)）⭐

**思路**：使用两个指针分别从左右两端向中间移动，动态维护左右最高值。

```go
func trapTwoPointers(height []int) int {
    left, right := 0, n-1
    leftMax, rightMax := 0, 0
    
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                totalWater += leftMax - height[left]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                totalWater += rightMax - height[right]
            }
            right--
        }
    }
}
```

### 方法4：单调栈（时间复杂度O(n)，空间复杂度O(n)）

**思路**：维护一个递减栈，当遇到更高的柱子时，计算能形成的凹槽雨水。

```go
func trapMonotonicStack(height []int) int {
    stack := make([]int, 0)
    
    for i := 0; i < n; i++ {
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            bottom := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            if len(stack) == 0 {
                break
            }
            
            left := stack[len(stack)-1]
            width := i - left - 1
            minHeight := min(height[left], height[i])
            waterHeight := minHeight - height[bottom]
            
            totalWater += width * waterHeight
        }
        stack = append(stack, i)
    }
}
```

## 📊 算法分析

### 时间复杂度对比

| 方法     | 时间复杂度 | 空间复杂度 | 特点                   |
| -------- | ---------- | ---------- | ---------------------- |
| 暴力法   | O(n²)      | O(1)       | 简单直观，效率低       |
| 动态规划 | O(n)       | O(n)       | 空间换时间，易理解     |
| 双指针   | O(n)       | O(1)       | **最优解**，时空都最优 |
| 单调栈   | O(n)       | O(n)       | 思路独特，适合扩展     |

### 双指针核心原理

**关键洞察**：对于位置i，只要知道 `min(leftMax[i], rightMax[i])`，就能计算雨水量。

双指针的精妙之处：
1. 当 `height[left] < height[right]` 时，左边的雨水量只取决于 `leftMax`
2. 当 `height[right] <= height[left]` 时，右边的雨水量只取决于 `rightMax`
3. 这样就避免了预计算整个数组的最高值

### 双指针过程可视化

以数组 `[0,1,0,2,1,0,1,3,2,1,2,1]` 为例：

```
步骤  左指针  右指针  左最高  右最高  当前水量  总水量
1     0      11     0      1      0       0
2     1      11     1      1      0       0
3     2      11     1      1      1       1
4     3      11     2      1      0       1
5     4      11     2      1      1       2
6     5      11     2      1      2       4
7     6      11     2      1      1       5
8     7      11     3      1      0       5
9     8      10     3      2      1       6
10    9      10     3      2      1       6
```

## 🎯 关键技巧

### 1. 双指针的移动策略
```go
if height[left] < height[right] {
    // 处理左边，因为左边的限制更小
    left++
} else {
    // 处理右边，因为右边的限制更小
    right--
}
```

### 2. 单调栈的应用
- 维护递减栈
- 遇到更高柱子时计算凹槽
- 适合处理局部最优问题

### 3. 边界条件处理
- 数组长度 ≤ 2：无法形成凹槽，返回0
- 首尾位置：不能接雨水
- 负数处理：题目保证非负整数

## 🔧 实际应用

### 1. 工程问题
- 雨水收集系统设计
- 地形分析和排水规划
- 建筑物屋顶设计

### 2. 算法竞赛
- 单调栈的经典应用
- 双指针技巧的典型场景
- 动态规划的空间优化

### 3. 数据结构
- 栈的应用场景
- 数组处理技巧
- 空间优化策略

## 📈 性能优化

### 1. 空间优化
```go
// 动态规划 -> 双指针
// O(n) 空间 -> O(1) 空间
```

### 2. 时间优化
```go
// 暴力法 -> 双指针
// O(n²) 时间 -> O(n) 时间
```

### 3. 缓存优化
```go
// 预计算常用值
leftMax := make([]int, n)
rightMax := make([]int, n)
```

## 🧪 测试用例

### 基础测试
```go
testCases := []struct {
    input    []int
    expected int
}{
    {[]int{0,1,0,2,1,0,1,3,2,1,2,1}, 6},
    {[]int{4,2,0,3,2,5}, 9},
    {[]int{3,0,2,0,4}, 7},
    {[]int{}, 0},
    {[]int{1}, 0},
    {[]int{1,2}, 0},
}
```

### 边界测试
```go
extremeCases := []struct {
    input    []int
    expected int
}{
    {[]int{0,0,0}, 0},           // 全为0
    {[]int{3,3,3}, 0},           // 全相等
    {[]int{1,2,3,4,5}, 0},       // 递增
    {[]int{5,4,3,2,1}, 0},       // 递减
    {[]int{3,0,0,0,3}, 9},       // 大凹槽
}
```

## 💡 扩展思考

### 1. 变种问题
- 接雨水II（2D版本）
- 柱状图中最大的矩形
- 最大矩形面积

### 2. 优化方向
- 并行处理大数组
- 内存映射处理超大数据
- GPU加速计算

### 3. 实际应用扩展
- 图像处理中的填充算法
- 游戏中的地形生成
- 金融中的价格分析

## 🎯 总结

接雨水问题是一个经典的算法题，展示了多种解题思路：

1. **暴力法**：直观但效率低，适合理解问题
2. **动态规划**：空间换时间，思路清晰
3. **双指针**：最优解，时空复杂度都最优 ⭐
4. **单调栈**：思路独特，适合扩展应用

**核心思想**：每个位置的雨水量 = min(左边最高, 右边最高) - 当前高度

**最佳实践**：推荐使用双指针法，因为它在时间和空间上都是最优的。

通过这道题，我们学会了：
- 双指针技巧的精妙应用
- 单调栈的使用场景
- 动态规划的空间优化
- 算法复杂度的权衡

---
