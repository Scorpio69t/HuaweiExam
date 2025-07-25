# 560. 和为 K 的子数组

## 描述

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。

## 示例 1

输入：nums = [1,1,1], k = 2
输出：2
解释：和为 2 的子数组为 [1,1] 和 [1,1]。

## 示例 2

输入：nums = [1,2,3], k = 3
输出：2
解释：和为 3 的子数组为 [1,2] 和 [3]。

## 示例 3

输入：nums = [1,-1,0], k = 0
输出：3
解释：和为 0 的子数组为 [1,-1], [-1,0] 和 [0]。

## 提示

- 1 <= nums.length <= 2 * 10^4
- -1000 <= nums[i] <= 1000
- -10^7 <= k <= 10^7

## 解题思路

### 方法一：前缀和 + 哈希表（推荐）

**核心思想**：
- 使用前缀和数组快速计算任意区间的和
- 利用哈希表记录每个前缀和出现的次数
- 对于当前位置i，查找前缀和为 `prefix[i] - k` 的位置个数

**算法步骤**：
1. 初始化前缀和数组 prefix[0] = 0
2. 遍历数组，计算前缀和 prefix[i] = prefix[i-1] + nums[i-1]
3. 对于每个位置i，查找哈希表中 `prefix[i] - k` 的个数
4. 将当前前缀和加入哈希表

**时间复杂度**：O(n)，其中n是数组长度
**空间复杂度**：O(n)

### 方法二：暴力枚举

**核心思想**：
- 枚举所有可能的子数组起点和终点
- 计算每个子数组的和，统计等于k的个数

**时间复杂度**：O(n²)
**空间复杂度**：O(1)

### 方法三：滑动窗口（仅适用于非负数数组）

**核心思想**：
- 当数组元素都是非负数时，可以使用滑动窗口
- 维护一个窗口，当窗口和小于k时扩展，大于k时收缩

**时间复杂度**：O(n)
**空间复杂度**：O(1)

## 代码实现

```go
// 前缀和 + 哈希表解法
func subarraySum(nums []int, k int) int {
    count := 0
    prefixSum := 0
    prefixMap := make(map[int]int)
    prefixMap[0] = 1 // 初始化，空数组的前缀和为0
    
    for _, num := range nums {
        prefixSum += num
        // 查找前缀和为 prefixSum - k 的个数
        if val, exists := prefixMap[prefixSum-k]; exists {
            count += val
        }
        // 将当前前缀和加入哈希表
        prefixMap[prefixSum]++
    }
    
    return count
}
```

## 复杂度分析

| 方法          | 时间复杂度 | 空间复杂度 | 适用场景           |
| ------------- | ---------- | ---------- | ------------------ |
| 前缀和+哈希表 | O(n)       | O(n)       | 推荐，通用解法     |
| 暴力枚举      | O(n²)      | O(1)       | 小规模数据         |
| 滑动窗口      | O(n)       | O(1)       | 仅适用于非负数数组 |

## 算法图解

```
数组: [1, 1, 1], k = 2

前缀和: [0, 1, 2, 3]
索引:   0  1  2  3

查找过程:
i=1: prefixSum=1, 查找 prefixSum-k=1-2=-1, 不存在
i=2: prefixSum=2, 查找 prefixSum-k=2-2=0, 存在1个
i=3: prefixSum=3, 查找 prefixSum-k=3-2=1, 存在1个

结果: 2
```

## 测试用例

```go
func main() {
    // 测试用例1
    nums1 := []int{1, 1, 1}
    k1 := 2
    fmt.Printf("测试用例1: nums=%v, k=%d, 结果=%d\n", nums1, k1, subarraySum(nums1, k1))
    
    // 测试用例2
    nums2 := []int{1, 2, 3}
    k2 := 3
    fmt.Printf("测试用例2: nums=%v, k=%d, 结果=%d\n", nums2, k2, subarraySum(nums2, k2))
    
    // 测试用例3
    nums3 := []int{1, -1, 0}
    k3 := 0
    fmt.Printf("测试用例3: nums=%v, k=%d, 结果=%d\n", nums3, k3, subarraySum(nums3, k3))
    
    // 边界测试
    nums4 := []int{1}
    k4 := 1
    fmt.Printf("边界测试1: nums=%v, k=%d, 结果=%d\n", nums4, k4, subarraySum(nums4, k4))
    
    nums5 := []int{1, 2, 3, 4, 5}
    k5 := 9
    fmt.Printf("边界测试2: nums=%v, k=%d, 结果=%d\n", nums5, k5, subarraySum(nums5, k5))
}
```
