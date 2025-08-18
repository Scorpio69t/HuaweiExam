# 15. 三数之和

## 题目描述

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

## 示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

## 示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。

## 示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。

## 提示：

- 3 <= nums.length <= 3000
- -10^5 <= nums[i] <= 10^5

## 解题思路

### 方法一：排序 + 双指针（推荐）

核心思想：
- 先对数组从小到大排序；
- 枚举第一个数 i，随后在区间 [i+1, n-1] 用双指针 l、r 寻找两数之和为 -nums[i]；
- 为避免重复：
  - i 若与前一个数相同则跳过；
  - 每次找到一个解后，移动 l 与 r 并跳过与上一位置相同的值。

步骤：
1. 排序 nums；
2. 枚举 i in [0..n-3]：若 nums[i] > 0，直接 break；若 i>0 且 nums[i]==nums[i-1]，continue；
3. 令 l=i+1, r=n-1，计算 s=nums[i]+nums[l]+nums[r]：
   - s==0：记录三元组并跳过重复的 l、r；
   - s<0：l++；
   - s>0：r--。

复杂度：
- 时间 O(n^2)，空间 O(1) 额外空间（不计返回结果）。

### 方法二：排序 + 哈希辅助查找（备选）

思路：
- 固定 i 后，在右侧用哈希集合寻找 two-sum 的补数 target=-nums[i]；
- 小心控制去重。实现上不如双指针直观，一般作为备选或对拍。

时间复杂度同样为 O(n^2)。

## 代码实现（Go）

```go
// 方法一：排序 + 双指针
func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 3 { return res }
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}
		if nums[i] > 0 { // 剪枝
			break
		}
		l, r := i+1, n-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] { l++ }
				r--
				for l < r && nums[r] == nums[r+1] { r-- }
			} else if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
```

## 复杂度分析

- 时间复杂度：O(n^2)
- 空间复杂度：O(1)（不计输出结果）

## 测试用例

```go
func main() {
	cases := [][]int{
		{-1, 0, 1, 2, -1, -4}, // [[-1,-1,2],[-1,0,1]]
		{0, 1, 1},             // []
		{0, 0, 0},             // [[0,0,0]]
	}
	for _, c := range cases {
		fmt.Println(threeSum(c))
	}
}
```

## 边界与注意点

- 多个零：如 [0,0,0,0] 结果应只有 [0,0,0]
- 大量重复值时的去重处理
- 排序后才能使用双指针并简化去重逻辑

## 运行

在当前目录下运行：

```bash
go run main.go
```
