# 33. 搜索旋转排序数组

## 描述

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

## 示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

## 示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

## 示例 3：

输入：nums = [1], target = 0
输出：-1

## 示例 4：

输入：nums = [3,1], target = 1
输出：1

## 提示：

- 1 <= nums.length <= 5000
- -10^4 <= nums[i] <= 10^4
- nums 中的每个值都 独一无二
- 题目数据保证 nums 在预先未知的某个下标上进行了旋转
- -10^4 <= target <= 10^4

## 解题思路

### 方法一：二分查找（推荐）

**核心思想**：
- 旋转数组可以分成两个有序部分
- 通过比较中间元素与边界元素，确定哪一半是有序的
- 在有序的一半中进行二分查找，在无序的一半中继续递归

**算法步骤**：
1. 初始化左右指针 left = 0, right = len(nums) - 1
2. 计算中间位置 mid = (left + right) / 2
3. 如果 nums[mid] == target，返回 mid
4. 判断左半部分是否有序：
   - 如果 nums[left] <= nums[mid]，左半部分有序
   - 如果 target 在左半部分范围内，在左半部分搜索
   - 否则在右半部分搜索
5. 否则右半部分有序：
   - 如果 target 在右半部分范围内，在右半部分搜索
   - 否则在左半部分搜索

**时间复杂度**：O(log n)
**空间复杂度**：O(1)

### 方法二：先找旋转点，再二分查找

**核心思想**：
- 先找到旋转点（最小值的位置）
- 根据 target 与边界值的比较，确定在哪个有序区间搜索
- 在确定的有序区间中进行二分查找

**时间复杂度**：O(log n)
**空间复杂度**：O(1)

### 方法三：线性搜索

**核心思想**：
- 直接遍历数组查找目标值
- 适用于小规模数据或调试

**时间复杂度**：O(n)
**空间复杂度**：O(1)

## 代码实现

```go
// 二分查找解法
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return mid
        }
        
        // 判断左半部分是否有序
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
    
    return -1
}
```

## 复杂度分析

| 方法          | 时间复杂度 | 空间复杂度 | 适用场景     |
| ------------- | ---------- | ---------- | ------------ |
| 二分查找      | O(log n)   | O(1)       | 推荐，最优解 |
| 找旋转点+二分 | O(log n)   | O(1)       | 思路清晰     |
| 线性搜索      | O(n)       | O(1)       | 小规模数据   |

## 算法图解

```
示例: nums = [4,5,6,7,0,1,2], target = 0

第一次二分: left=0, right=6, mid=3
nums[mid]=7, nums[left]=4, nums[right]=2
左半部分[4,5,6,7]有序，但target=0不在范围内
在右半部分[0,1,2]搜索

第二次二分: left=4, right=6, mid=5
nums[mid]=1, nums[left]=0, nums[right]=2
左半部分[0]有序，target=0在范围内
在左半部分搜索

第三次二分: left=4, right=4, mid=4
nums[mid]=0 == target，返回4
```

## 边界情况处理

1. **数组长度为1**：直接比较
2. **数组长度为2**：分别比较两个元素
3. **没有旋转**：数组完全有序
4. **目标值不存在**：返回-1
5. **目标值在边界**：正确处理边界条件

## 测试用例

```go
func main() {
    // 测试用例1
    nums1 := []int{4, 5, 6, 7, 0, 1, 2}
    target1 := 0
    fmt.Printf("测试用例1: nums=%v, target=%d, 结果=%d\n", nums1, target1, search(nums1, target1))
    
    // 测试用例2
    nums2 := []int{4, 5, 6, 7, 0, 1, 2}
    target2 := 3
    fmt.Printf("测试用例2: nums=%v, target=%d, 结果=%d\n", nums2, target2, search(nums2, target2))
    
    // 边界测试
    nums3 := []int{1}
    target3 := 0
    fmt.Printf("边界测试1: nums=%v, target=%d, 结果=%d\n", nums3, target3, search(nums3, target3))
    
    nums4 := []int{3, 1}
    target4 := 1
    fmt.Printf("边界测试2: nums=%v, target=%d, 结果=%d\n", nums4, target4, search(nums4, target4))
}
```
