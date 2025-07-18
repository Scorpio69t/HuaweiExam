# 70. 爬楼梯

## 描述

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

## 示例 1

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

## 示例 2

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶

## 提示

- 1 <= n <= 45

## 解题思路

### 核心分析

这是一道经典的动态规划入门题目，本质上是**斐波那契数列**的变形。

### 问题建模

要到达第n阶楼梯，可以从两个位置到达：
1. 从第(n-1)阶爬1步
2. 从第(n-2)阶爬2步

因此：`f(n) = f(n-1) + f(n-2)`

这正是斐波那契数列的递推关系！

### 算法实现

#### 方法1：动态规划（标准解法）

**状态定义**：
- `dp[i]` 表示到达第i阶楼梯的方法数
- `dp[i] = dp[i-1] + dp[i-2]`

**边界条件**：
- `dp[1] = 1`（只有一种方法：爬1步）
- `dp[2] = 2`（两种方法：1+1 或 2）

```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 2
    
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    
    return dp[n]
}
```

**时间复杂度**：O(n)
**空间复杂度**：O(n)

#### 方法2：空间优化动态规划（最优解）

由于状态转移只依赖前两个状态，可以用两个变量代替数组。

```go
func climbStairsOptimized(n int) int {
    if n <= 2 {
        return n
    }
    
    prev2 := 1  // f(1)
    prev1 := 2  // f(2)
    
    for i := 3; i <= n; i++ {
        curr := prev1 + prev2
        prev2 = prev1
        prev1 = curr
    }
    
    return prev1
}
```

**时间复杂度**：O(n)
**空间复杂度**：O(1)

#### 方法3：递归 + 记忆化

使用递归思路，配合记忆化避免重复计算。

```go
func climbStairsMemo(n int) int {
    memo := make(map[int]int)
    return climbStairsMemoHelper(n, memo)
}

func climbStairsMemoHelper(n int, memo map[int]int) int {
    if n <= 2 {
        return n
    }
    
    if val, exists := memo[n]; exists {
        return val
    }
    
    result := climbStairsMemoHelper(n-1, memo) + climbStairsMemoHelper(n-2, memo)
    memo[n] = result
    return result
}
```

**时间复杂度**：O(n)
**空间复杂度**：O(n)

#### 方法4：数学公式（斐波那契通项公式）

使用贝尔纳公式直接计算斐波那契数列第n项。

```go
func climbStairsFormula(n int) int {
    if n <= 2 {
        return n
    }
    
    sqrt5 := math.Sqrt(5)
    phi := (1 + sqrt5) / 2        // 黄金比例
    psi := (1 - sqrt5) / 2        // 共轭黄金比例
    
    // 斐波那契通项公式：F(n) = (φ^n - ψ^n) / √5
    // 但这里是 F(n+1)，因为我们的序列是 f(1)=1, f(2)=2
    result := (math.Pow(phi, float64(n+1)) - math.Pow(psi, float64(n+1))) / sqrt5
    
    return int(math.Round(result))
}
```

**时间复杂度**：O(1)
**空间复杂度**：O(1)

#### 方法5：矩阵快速幂

使用矩阵快速幂计算斐波那契数列，适合处理大数。

```go
func climbStairsMatrix(n int) int {
    if n <= 2 {
        return n
    }
    
    // 转换矩阵: [[1,1],[1,0]]
    base := [][]int{{1, 1}, {1, 0}}
    result := matrixPower(base, n-1)
    
    // result * [F(2), F(1)] = [F(n+1), F(n)]
    return result[0][0]*2 + result[0][1]*1
}

func matrixPower(matrix [][]int, n int) [][]int {
    size := len(matrix)
    result := make([][]int, size)
    for i := range result {
        result[i] = make([]int, size)
        result[i][i] = 1  // 单位矩阵
    }
    
    base := make([][]int, size)
    for i := range base {
        base[i] = make([]int, size)
        copy(base[i], matrix[i])
    }
    
    for n > 0 {
        if n&1 == 1 {
            result = matrixMultiply(result, base)
        }
        base = matrixMultiply(base, base)
        n >>= 1
    }
    
    return result
}

func matrixMultiply(a, b [][]int) [][]int {
    size := len(a)
    result := make([][]int, size)
    for i := range result {
        result[i] = make([]int, size)
        for j := 0; j < size; j++ {
            for k := 0; k < size; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result
}
```

**时间复杂度**：O(log n)
**空间复杂度**：O(1)

## 复杂度分析

| 方法       | 时间复杂度 | 空间复杂度 | 优缺点               |
| ---------- | ---------- | ---------- | -------------------- |
| 标准DP     | O(n)       | O(n)       | 思路清晰，易理解     |
| 空间优化DP | O(n)       | O(1)       | 最实用的解法 ⭐       |
| 记忆化递归 | O(n)       | O(n)       | 自顶向下，递归栈开销 |
| 数学公式   | O(1)       | O(1)       | 最快，但有精度问题   |
| 矩阵快速幂 | O(log n)   | O(1)       | 适合大数，复杂度低   |

## 核心要点

1. **斐波那契本质**：问题等价于求斐波那契数列第(n+1)项
2. **状态转移**：每个状态只依赖前两个状态
3. **空间优化**：可以用O(1)空间代替O(n)空间
4. **边界处理**：n=1和n=2的特殊情况

## 数学推导

### 递推关系证明

设 `f(n)` 表示到达第n阶楼梯的方法数：

**递推关系**：
```
f(n) = f(n-1) + f(n-2)  (n ≥ 3)
```

**初始条件**：
```
f(1) = 1
f(2) = 2
```

**证明**：
要到达第n阶，只能从两个位置到达：
- 从第(n-1)阶爬1步：有 `f(n-1)` 种方法
- 从第(n-2)阶爬2步：有 `f(n-2)` 种方法
- 总计：`f(n-1) + f(n-2)` 种方法

### 斐波那契数列对应关系

爬楼梯序列：1, 2, 3, 5, 8, 13, 21, 34, ...
斐波那契序列：1, 1, 2, 3, 5, 8, 13, 21, ...

**关系**：`climbStairs(n) = fibonacci(n+1)`

### 通项公式推导

斐波那契数列通项公式：
```
F(n) = (φ^n - ψ^n) / √5
```

其中：
- φ = (1 + √5) / 2 ≈ 1.618（黄金比例）
- ψ = (1 - √5) / 2 ≈ -0.618

因此：
```
climbStairs(n) = F(n+1) = (φ^(n+1) - ψ^(n+1)) / √5
```

## 执行流程图

```mermaid
graph TD
    A[开始: 输入n] --> B{边界判断}
    B -->|n ≤ 2| C[返回n]
    B -->|n > 2| D[选择算法]
    
    D --> E[标准DP]
    D --> F[空间优化DP]
    D --> G[记忆化递归]
    D --> H[数学公式]
    D --> I[矩阵快速幂]
    
    E --> J[创建dp数组]
    F --> K[使用两个变量]
    G --> L[递归+缓存]
    H --> M[黄金比例公式]
    I --> N[矩阵乘法]
    
    J --> O[循环计算f1-fn]
    K --> P[循环更新prev1,prev2]
    L --> Q[递归计算子问题]
    M --> R[直接计算结果]
    N --> S[快速幂计算]
    
    O --> T[返回dp[n]]
    P --> T
    Q --> T
    R --> T
    S --> T
    C --> U[结束]
    T --> U
```

## 实际应用

1. **组合计数**：计算特定约束下的方案数
2. **路径规划**：网格中的路径计数问题
3. **动态规划优化**：状态压缩的经典例子
4. **算法面试**：考察DP基础的经典题目

## 扩展变形

1. **步数扩展**：如果可以爬1、2、3步怎么办？
2. **限制条件**：某些台阶不能踩怎么处理？
3. **成本问题**：每步有不同成本，求最小成本
4. **二维扩展**：在网格中从左上到右下的路径数

## 测试用例设计

```go
// 基础测试
n=1 → 1
n=2 → 2
n=3 → 3
n=4 → 5
n=5 → 8

// 边界测试
n=1 → 1 (最小值)
n=45 → 1836311903 (题目限制最大值)

// 斐波那契验证
n=10 → 89
n=20 → 10946
n=30 → 1346269

// 性能测试
大数值测试各算法效率对比
```

## 数学性质

### 黄金比例的美

斐波那契数列中相邻两项的比值趋近于黄金比例φ：
```
lim(n→∞) F(n+1)/F(n) = φ = (1+√5)/2 ≈ 1.618
```

### 奇偶性质

```
F(n) 为偶数 ⟺ n ≡ 0 (mod 3)
```

### 整除性质

```
gcd(F(m), F(n)) = F(gcd(m, n))
```

### 平方和性质

```
F(1)² + F(2)² + ... + F(n)² = F(n) × F(n+1)
```

