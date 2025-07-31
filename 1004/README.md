# 1004. 最大连续1的个数 III

## 描述

给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。

## 示例 1

输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。

## 示例 2

输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。

## 提示

- 1 <= nums.length <= 10^5
- nums[i] 不是 0 就是 1
- 0 <= k <= nums.length

## 解题思路

### 方法一：滑动窗口（推荐）

```mermaid
graph TD
    A[开始: nums=[1,1,1,0,0,0,1,1,1,1,0], k=2] --> B[初始化: left=0, right=0, zeros=0]
    B --> C[扩展右边界: right++]
    C --> D{检查 nums[right]}
    D -->|1| E[继续扩展]
    D -->|0| F[zeros++]
    F --> G{zeros > k?}
    G -->|否| E
    G -->|是| H[收缩左边界]
    H --> I{检查 nums[left]}
    I -->|0| J[zeros--]
    I -->|1| K[继续收缩]
    J --> L[left++]
    K --> L
    L --> G
    E --> M[更新最大长度]
    M --> N{right < len(nums)?}
    N -->|是| C
    N -->|否| O[返回最大长度]
    
    style O fill:#90EE90
    style A fill:#E6F3FF
    style B fill:#FFF2CC
```

### 滑动窗口示例

```mermaid
graph LR
    A["初始状态"] --> B["扩展: [1,1,1,0,0,0,1,1,1,1,0]"]
    B --> C["窗口: [1,1,1,0,0] zeros=2"]
    C --> D["扩展: [1,1,1,0,0,0] zeros=3 > k"]
    D --> E["收缩: [1,1,0,0,0] zeros=2"]
    E --> F["扩展: [1,1,0,0,0,1,1,1,1,0]"]
    F --> G["最大窗口: [1,1,1,0,0,1,1,1,1,1,1] 长度=6"]
    
    style G fill:#90EE90
    style A fill:#E6F3FF
```

**核心思想**：
- 使用滑动窗口维护一个包含最多k个0的连续子数组
- 当窗口内的0的个数超过k时，收缩左边界
- 不断扩展右边界，记录最大窗口长度

**算法步骤**：
1. 初始化左右指针 left = 0, right = 0
2. 遍历数组，right指针向右移动
3. 当遇到0时，计数器zeros++
4. 当zeros > k时，收缩左边界直到zeros <= k
5. 更新最大长度 maxLen = max(maxLen, right - left + 1)

**时间复杂度**：O(n)，其中n是数组长度
**空间复杂度**：O(1)

### 方法二：前缀和 + 二分查找

**核心思想**：
- 对于每个位置i，找到最远的j使得区间[i,j]中0的个数不超过k
- 使用前缀和快速计算区间内0的个数
- 使用二分查找找到最远的j

**时间复杂度**：O(n log n)
**空间复杂度**：O(n)

### 方法三：动态规划

**核心思想**：
- dp[i][j]表示以位置i结尾，使用j次翻转机会的最长连续1长度
- 状态转移：如果nums[i]=1，dp[i][j] = dp[i-1][j] + 1
- 如果nums[i]=0且j>0，dp[i][j] = dp[i-1][j-1] + 1

**时间复杂度**：O(nk)
**空间复杂度**：O(nk)

## 代码实现

```go
// 滑动窗口解法
func longestOnes(nums []int, k int) int {
    left, zeros := 0, 0
    maxLen := 0
    
    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            zeros++
        }
        
        // 当0的个数超过k时，收缩左边界
        for zeros > k {
            if nums[left] == 0 {
                zeros--
            }
            left++
        }
        
        maxLen = max(maxLen, right-left+1)
    }
    
    return maxLen
}
```

## 复杂度分析

| 方法        | 时间复杂度 | 空间复杂度 | 适用场景     |
| ----------- | ---------- | ---------- | ------------ |
| 滑动窗口    | O(n)       | O(1)       | 推荐，最优解 |
| 前缀和+二分 | O(n log n) | O(n)       | 需要精确控制 |
| 动态规划    | O(nk)      | O(nk)      | k较小时      |

## 测试用例

```go
func main() {
    // 测试用例1
    nums1 := []int{1,1,1,0,0,0,1,1,1,1,0}
    k1 := 2
    fmt.Printf("测试用例1: nums=%v, k=%d, 结果=%d\n", nums1, k1, longestOnes(nums1, k1))
    
    // 测试用例2
    nums2 := []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
    k2 := 3
    fmt.Printf("测试用例2: nums=%v, k=%d, 结果=%d\n", nums2, k2, longestOnes(nums2, k2))
    
    // 边界测试
    nums3 := []int{1,1,1,1,1}
    k3 := 0
    fmt.Printf("边界测试1: nums=%v, k=%d, 结果=%d\n", nums3, k3, longestOnes(nums3, k3))
    
    nums4 := []int{0,0,0,0,0}
    k4 := 2
    fmt.Printf("边界测试2: nums=%v, k=%d, 结果=%d\n", nums4, k4, longestOnes(nums4, k4))
}
```
