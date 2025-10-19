# 81. 搜索旋转排序数组 II

## 题目描述

已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。

给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。

你必须尽可能减少整个操作步骤。


## 示例 1：

输入：nums = [2,5,6,0,0,1,2], target = 0
输出：true


## 示例 2：

输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false


## 提示：

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-104 <= target <= 104


## 进阶：

此题与 搜索旋转排序数组 相似，但本题中的 nums  可能包含 重复 元素。这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？

## 解题思路

### 问题深度分析

这是经典的**二分查找算法**问题，也是**旋转数组搜索**的典型应用。核心在于**处理重复元素**，在O(n)时间内搜索目标值。

#### 问题本质

给定旋转后的有序数组（可能包含重复元素），搜索目标值。这是一个**二分查找**问题，需要处理重复元素带来的复杂性。

#### 核心思想

**二分查找 + 重复元素处理**：
1. **二分查找**：使用左右指针缩小搜索范围
2. **重复元素处理**：当左右边界相等时，收缩边界
3. **旋转点判断**：确定哪一半是有序的
4. **目标值搜索**：在有序的一半中搜索目标值

**关键技巧**：
- 使用`left`和`right`指针进行二分查找
- 当`nums[left] == nums[mid] == nums[right]`时，收缩边界
- 判断哪一半是有序的
- 在有序的一半中搜索目标值

#### 关键难点分析

**难点1：重复元素的影响**
- 当`nums[left] == nums[mid] == nums[right]`时，无法判断哪一半有序
- 需要收缩边界来消除重复元素的影响
- 最坏情况下时间复杂度退化为O(n)

**难点2：旋转点的判断**
- 需要准确判断哪一半是有序的
- 左半部分有序：`nums[left] <= nums[mid]`
- 右半部分有序：`nums[mid] <= nums[right]`

**难点3：目标值的搜索**
- 在有序的一半中搜索目标值
- 需要考虑边界条件
- 需要正确处理重复元素

#### 典型情况分析

**情况1：一般情况**
```
nums = [2,5,6,0,0,1,2], target = 0
过程：
1. left=0, right=6, mid=3, nums[3]=0 → 找到
结果: true
```

**情况2：无解情况**
```
nums = [2,5,6,0,0,1,2], target = 3
过程：二分查找后未找到
结果: false
```

**情况3：重复元素较多**
```
nums = [1,1,1,1,1,1,1], target = 1
过程：需要收缩边界处理重复元素
结果: true
```

**情况4：边界情况**
```
nums = [1], target = 1
结果: true
```

#### 算法对比

| 算法     | 时间复杂度 | 空间复杂度 | 特点         |
| -------- | ---------- | ---------- | ------------ |
| 二分查找 | O(n)       | O(1)       | **最优解法** |
| 线性搜索 | O(n)       | O(1)       | 简单但效率低 |
| 哈希表   | O(n)       | O(n)       | 空间复杂度高 |
| 暴力法   | O(n)       | O(1)       | 效率相同     |

注：n为数组长度，最坏情况下时间复杂度为O(n)

### 算法流程图

#### 主算法流程（二分查找）

```mermaid
graph TD
    A[开始: nums, target] --> B[初始化: left=0, right=len-1]
    B --> C[left <= right?]
    C -->|否| D[返回false]
    C -->|是| E[mid = (left+right)/2]
    E --> F[nums[mid] == target?]
    F -->|是| G[返回true]
    F -->|否| H[nums[left] == nums[mid] == nums[right]?]
    H -->|是| I[left++, right--]
    H -->|否| J[nums[left] <= nums[mid]?]
    J -->|是| K[左半部分有序]
    J -->|否| L[右半部分有序]
    K --> M[target在左半部分?]
    L --> N[target在右半部分?]
    M -->|是| O[right = mid-1]
    M -->|否| P[left = mid+1]
    N -->|是| Q[left = mid+1]
    N -->|否| R[right = mid-1]
    I --> C
    O --> C
    P --> C
    Q --> C
    R --> C
```

#### 重复元素处理流程

```mermaid
graph TD
    A[检查重复元素] --> B{nums[left] == nums[mid] == nums[right]?}
    B -->|是| C[收缩边界]
    B -->|否| D[判断有序部分]
    C --> E[left++, right--]
    E --> F[继续二分查找]
    D --> G[确定有序部分]
    G --> H[在有序部分搜索]
    H --> I[更新搜索范围]
    I --> F
```

### 复杂度分析

#### 时间复杂度详解

**二分查找**：O(n)
- 最坏情况：所有元素相同，需要线性搜索
- 平均情况：O(log n)
- 总时间：O(n)

**线性搜索**：O(n)
- 遍历整个数组一次
- 时间复杂度固定为O(n)

#### 空间复杂度详解

**二分查找**：O(1)
- 只使用常数额外空间
- 原地搜索
- 总空间：O(1)

### 关键优化技巧

#### 技巧1：二分查找（最优解法）

```go
func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return true
        }
        
        // 处理重复元素
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
        } else if nums[left] <= nums[mid] {
            // 左半部分有序
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // 右半部分有序
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return false
}
```

**优势**：
- 时间复杂度：O(n)
- 空间复杂度：O(1)
- 处理重复元素

#### 技巧2：线性搜索

```go
func search(nums []int, target int) bool {
    for _, num := range nums {
        if num == target {
            return true
        }
    }
    return false
}
```

**特点**：简单直接，但效率较低

#### 技巧3：哈希表

```go
func search(nums []int, target int) bool {
    numMap := make(map[int]bool)
    for _, num := range nums {
        numMap[num] = true
    }
    return numMap[target]
}
```

**特点**：使用哈希表，空间复杂度高

#### 技巧4：优化版二分查找

```go
func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return true
        }
        
        // 处理重复元素
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
            continue
        }
        
        if nums[left] <= nums[mid] {
            // 左半部分有序
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // 右半部分有序
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return false
}
```

**特点**：优化重复元素处理逻辑

### 边界情况处理

1. **空数组**：返回false
2. **单元素**：直接比较
3. **全部相同**：需要特殊处理
4. **目标值在边界**：正确处理边界条件
5. **重复元素较多**：收缩边界处理

### 测试用例设计

#### 基础测试
```
输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
说明: 一般情况
```

#### 简单情况
```
输入: nums = [1], target = 1
输出: true
说明: 单元素情况
```

#### 特殊情况
```
输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
说明: 无解情况
```

#### 边界情况
```
输入: nums = [], target = 0
输出: false
说明: 空数组情况
```

### 常见错误与陷阱

#### 错误1：重复元素处理错误

```go
// ❌ 错误：没有处理重复元素
if nums[left] <= nums[mid] {
    // 直接判断，可能出错
}

// ✅ 正确：先处理重复元素
if nums[left] == nums[mid] && nums[mid] == nums[right] {
    left++
    right--
} else if nums[left] <= nums[mid] {
    // 然后判断有序部分
}
```

#### 错误2：边界条件错误

```go
// ❌ 错误：边界条件不正确
if nums[left] <= target && target < nums[mid] {
    right = mid // 错误：应该是mid-1
}

// ✅ 正确：边界条件正确
if nums[left] <= target && target < nums[mid] {
    right = mid - 1
}
```

#### 错误3：循环条件错误

```go
// ❌ 错误：循环条件不正确
for left < right { // 可能漏掉某些情况
    // ...
}

// ✅ 正确：使用正确的循环条件
for left <= right {
    // ...
}
```

### 实战技巧总结

1. **二分查找模板**：左右指针 + 中点判断
2. **重复元素处理**：收缩边界消除影响
3. **有序部分判断**：准确判断哪一半有序
4. **边界处理**：正确处理各种边界情况
5. **时间复杂度**：最坏情况下O(n)

### 进阶扩展

#### 扩展1：返回目标值索引

```go
func searchIndex(nums []int, target int) int {
    // 返回目标值的索引，未找到返回-1
    // ...
}
```

#### 扩展2：统计目标值出现次数

```go
func searchCount(nums []int, target int) int {
    // 统计目标值在数组中出现的次数
    // ...
}
```

#### 扩展3：支持多个目标值

```go
func searchMultiple(nums []int, targets []int) []bool {
    // 同时搜索多个目标值
    // ...
}
```

### 应用场景

1. **数据搜索**：在旋转数组中搜索数据
2. **算法竞赛**：二分查找基础
3. **系统设计**：高效数据检索
4. **数据分析**：快速数据查找
5. **游戏开发**：关卡数据搜索

## 代码实现

本题提供了四种不同的解法，重点掌握二分查找算法。

## 测试结果

| 测试用例 | 二分查找 | 线性搜索 | 哈希表 | 优化版 |
| -------- | -------- | -------- | ------ | ------ |
| 基础测试 | ✅        | ✅        | ✅      | ✅      |
| 简单情况 | ✅        | ✅        | ✅      | ✅      |
| 特殊情况 | ✅        | ✅        | ✅      | ✅      |
| 边界情况 | ✅        | ✅        | ✅      | ✅      |

## 核心收获

1. **二分查找**：旋转数组搜索的经典应用
2. **重复元素处理**：收缩边界消除影响
3. **有序部分判断**：准确判断哪一半有序
4. **边界处理**：各种边界情况的考虑
5. **时间复杂度**：最坏情况下O(n)

## 应用拓展

- 数据搜索和检索
- 算法竞赛基础
- 系统设计应用
- 数据分析技术
- 游戏开发优化

