# 213. 打家劫舍 II

## 描述

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

## 示例 1

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

## 示例 2

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

## 示例 3

输入：nums = [1,2,3]
输出：3

## 提示

- 1 <= nums.length <= 100
- 0 <= nums[i] <= 1000

## 解题思路

### 核心分析

这道题是经典动态规划问题"打家劫舍"的环形变种。关键区别是房屋呈环形排列，第一间和最后一间房屋相邻。

### 问题转化

由于房屋呈环形，需要考虑两种情况：
1. 偷第一间房屋：不能偷最后一间房屋
2. 不偷第一间房屋：可以考虑偷最后一间房屋

将环形问题转化为两次线性动态规划：
- 情况1：在 nums[0...n-2] 范围内求最大值（包含第一间，排除最后一间）
- 情况2：在 nums[1...n-1] 范围内求最大值（排除第一间，包含最后一间）

取两种情况的最大值即为答案。

### 算法实现

#### 方法1：动态规划（标准解法）

**状态定义**：
- `dp[i]` 表示前i间房屋能偷到的最大金额
- `dp[i] = max(dp[i-1], dp[i-2] + nums[i])`

**转移方程**：
对于每间房屋，有两种选择：
1. 偷当前房屋：`dp[i-2] + nums[i]`
2. 不偷当前房屋：`dp[i-1]`

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 1 { return nums[0] }
    if n == 2 { return max(nums[0], nums[1]) }
    
    // 情况1：偷第一间，不偷最后一间 [0, n-2]
    case1 := robLinear(nums[:n-1])
    
    // 情况2：不偷第一间，可偷最后一间 [1, n-1]
    case2 := robLinear(nums[1:])
    
    return max(case1, case2)
}

func robLinear(nums []int) int {
    n := len(nums)
    if n == 0 { return 0 }
    if n == 1 { return nums[0] }
    
    dp := make([]int, n)
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], dp[i-2] + nums[i])
    }
    
    return dp[n-1]
}
```

**时间复杂度**：O(n)
**空间复杂度**：O(n)

#### 方法2：空间优化动态规划

由于状态转移只依赖前两个状态，可以用两个变量代替数组。

```go
func robOptimized(nums []int) int {
    n := len(nums)
    if n == 1 { return nums[0] }
    if n == 2 { return max(nums[0], nums[1]) }
    
    case1 := robLinearOptimized(nums[:n-1])
    case2 := robLinearOptimized(nums[1:])
    
    return max(case1, case2)
}

func robLinearOptimized(nums []int) int {
    n := len(nums)
    if n == 0 { return 0 }
    if n == 1 { return nums[0] }
    
    prev2 := nums[0]
    prev1 := max(nums[0], nums[1])
    
    for i := 2; i < n; i++ {
        curr := max(prev1, prev2 + nums[i])
        prev2 = prev1
        prev1 = curr
    }
    
    return prev1
}
```

**时间复杂度**：O(n)
**空间复杂度**：O(1)

#### 方法3：记忆化递归

使用递归思路，配合记忆化避免重复计算。

```go
func robMemo(nums []int) int {
    n := len(nums)
    if n == 1 { return nums[0] }
    if n == 2 { return max(nums[0], nums[1]) }
    
    // 情况1：偷第一间，不偷最后一间
    memo1 := make(map[int]int)
    case1 := robMemoHelper(nums[:n-1], 0, memo1)
    
    // 情况2：不偷第一间，可偷最后一间
    memo2 := make(map[int]int)
    case2 := robMemoHelper(nums[1:], 0, memo2)
    
    return max(case1, case2)
}

func robMemoHelper(nums []int, index int, memo map[int]int) int {
    if index >= len(nums) { return 0 }
    if val, exists := memo[index]; exists {
        return val
    }
    
    // 偷当前房屋
    rob := nums[index] + robMemoHelper(nums, index+2, memo)
    // 不偷当前房屋
    notRob := robMemoHelper(nums, index+1, memo)
    
    result := max(rob, notRob)
    memo[index] = result
    return result
}
```

**时间复杂度**：O(n)
**空间复杂度**：O(n)

#### 方法4：状态机模拟

将问题建模为状态机，明确每个状态的含义。

```go
func robStateMachine(nums []int) int {
    n := len(nums)
    if n == 1 { return nums[0] }
    if n == 2 { return max(nums[0], nums[1]) }
    
    case1 := robStateMachineHelper(nums[:n-1])
    case2 := robStateMachineHelper(nums[1:])
    
    return max(case1, case2)
}

func robStateMachineHelper(nums []int) int {
    n := len(nums)
    if n == 0 { return 0 }
    if n == 1 { return nums[0] }
    
    // 状态：[当前位置是否偷窃][累计金额]
    // true: 偷了当前房屋
    // false: 没偷当前房屋
    
    robbed := nums[0]    // 偷了第0间房屋的最大金额
    notRobbed := 0       // 没偷第0间房屋的最大金额
    
    for i := 1; i < n; i++ {
        newRobbed := notRobbed + nums[i]  // 偷当前房屋
        newNotRobbed := max(robbed, notRobbed)  // 不偷当前房屋
        
        robbed = newRobbed
        notRobbed = newNotRobbed
    }
    
    return max(robbed, notRobbed)
}
```

**时间复杂度**：O(n)
**空间复杂度**：O(1)

## 复杂度分析

| 方法       | 时间复杂度 | 空间复杂度 | 优缺点               |
| ---------- | ---------- | ---------- | -------------------- |
| 标准DP     | O(n)       | O(n)       | 思路清晰，易理解     |
| 空间优化DP | O(n)       | O(1)       | 最优解，推荐使用     |
| 记忆化递归 | O(n)       | O(n)       | 自顶向下，递归栈开销 |
| 状态机     | O(n)       | O(1)       | 状态清晰，逻辑直观   |

## 核心要点

1. **环形转化**：将环形问题分解为两个线性子问题
2. **状态转移**：每个位置有偷或不偷两种选择
3. **边界处理**：特殊处理长度为1和2的情况
4. **最优子结构**：当前最优解依赖于之前的最优解

## 数学证明

### 最优子结构证明

设 `f(i)` 表示前i间房屋的最大收益：

**递推关系**：
```
f(i) = max(f(i-1), f(i-2) + nums[i])
```

**证明**：
- 如果偷第i间房屋，则不能偷第i-1间，最大收益为 `f(i-2) + nums[i]`
- 如果不偷第i间房屋，最大收益为 `f(i-1)`
- 两者取最大值即为最优解

### 环形约束处理

对于环形数组 `[0, 1, 2, ..., n-1]`：

**情况分析**：
1. 若偷房屋0，则不能偷房屋n-1，问题等价于求 `[0, 1, ..., n-2]` 的最大值
2. 若不偷房屋0，则可偷房屋n-1，问题等价于求 `[1, 2, ..., n-1]` 的最大值

**正确性**：任何最优解必然属于上述两种情况之一。

## 执行流程图

```mermaid
graph TD
    A[开始: 输入nums数组] --> B{数组长度判断}
    B -->|n=1| C[返回nums[0]]
    B -->|n=2| D[返回max nums[0], nums[1]]
    B -->|n≥3| E[分解为两个子问题]
    
    E --> F[情况1: nums[0...n-2] 包含第一间，排除最后一间]
    E --> G[情况2: nums[1...n-1] 排除第一间，包含最后一间]
    
    F --> H[线性DP求解子问题1]
    G --> I[线性DP求解子问题2]
    
    H --> J[得到结果1]
    I --> K[得到结果2]
    
    J --> L[返回max 结果1, 结果2]
    K --> L
    C --> M[结束]
    D --> M
    L --> M
```

## 实际应用

1. **资源分配**：在有约束条件下的最优资源分配
2. **调度问题**：环形任务调度的最优化
3. **投资策略**：在周期性投资中的风险控制
4. **游戏设计**：回合制游戏中的策略优化

## 扩展思考

1. **多环问题**：如果是多层环形结构怎么处理？
2. **动态约束**：如果相邻约束会动态变化呢？
3. **多维扩展**：二维网格上的环形约束问题
4. **概率版本**：每间房屋被发现的概率不同

## 测试用例设计

```go
// 基础测试
[2,3,2] → 3
[1,2,3,1] → 4
[1,2,3] → 3

// 边界测试
[5] → 5
[1,2] → 2
[2,1] → 2

// 极值测试
[1,1,1,1] → 2
[1000,1,1,1000] → 2000
[0,0,0] → 0

// 递增序列
[1,2,3,4,5] → 9 (选择1,3,5)

// 递减序列
[5,4,3,2,1] → 8 (选择5,3,1)
```

