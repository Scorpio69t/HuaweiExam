# 18. 四数之和

## 题目描述

给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

- 0 <= a, b, c, d < n
- a、b、c 和 d 互不相同
- nums[a] + nums[b] + nums[c] + nums[d] == target

你可以按 任意顺序 返回答案 。

## 示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

## 示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

## 提示：

- 1 <= nums.length <= 200
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9

## 解题思路

本题是“kSum”系列的典型代表。最直接的思路是：
- 先对数组排序
- 固定两个数，用双指针在剩余区间内查找另外两个数（3重循环 + 双指针）
- 通过跳过重复元素与剪枝优化，保证不重复且加速

此外，还可以抽象成通用的 kSum 递归框架：当 k==2 时用双指针解决，否则固定一个元素，递归解决 (k-1)Sum。

### 算法一：排序 + 双指针（推荐）
- 排序后，外层两重循环固定 i、j
- 内层用左右指针 left、right 搜索两数和，使四数之和为 target
- 跳过重复元素（i、j、left、right 层面）避免重复解
- 剪枝：用最小可能和/最大可能和与 target 比较，提前 break/continue

```mermaid
flowchart TD
  A[排序nums] --> B[i从0到n-4]
  B --> C{去重/剪枝}
  C -->|不满足| B
  C --> D[j从i+1到n-3]
  D --> E{去重/剪枝}
  E -->|不满足| D
  E --> F[left=j+1,right=n-1]
  F --> G{left<right?}
  G -->|否| D
  G --> H[sum=nums[i]+nums[j]+nums[left]+nums[right]]
  H --> I{sum==target?}
  I -->|是| J[加入解并去重移动]
  I -->|sum<target| K[left++]
  I -->|sum>target| L[right--]
  J --> G
  K --> G
  L --> G
```

### 算法二：通用 kSum（含 2Sum 双指针）
- 排序
- kSum(nums,k,start,target):
  - 若 k==2，用双指针在区间 [start,n) 里求两数和为 target 的所有解
  - 否则，从 start 枚举 i，跳过重复元素，并递归 kSum(nums,k-1,i+1,target-nums[i])
- 结合最小/最大可能和剪枝

```mermaid
flowchart TD
  A[sort(nums)] --> B[kSum(nums,4,0,target)]
  B --> C{k==2?}
  C -->|是| D[2Sum 双指针]
  C -->|否| E[for i from start to n-k]
  E --> F{跳过重复}
  F --> E
  E --> G[递归 kSum(k-1,i+1,target-nums[i])]
  G --> H[前缀拼接nums[i]并收集]
  H --> E
```

### 复杂度
- 排序 + 双指针：时间 O(n^3)，空间 O(1)（不计结果集）
- kSum：最坏也为 O(n^{k-1})，本题 k=4 即 O(n^3)

### 关键细节
- **去重**：
  - i>0 且 nums[i]==nums[i-1] 跳过
  - j>i+1 且 nums[j]==nums[j-1] 跳过
  - 命中答案后 left/right 向内移动并跳过相同值
- **剪枝**：
  - 对当前固定前缀，比较最小可能和与最大可能和与 target 的关系进行 break/continue
- **long 溢出处理**：
  - 计算和时转为 int64 避免大数相加溢出

### 代码实现要点
1. 先排序，明确单调结构便于双指针
2. 严格实现四处去重逻辑，保证不重复
3. 使用 int64 做加法再比较，避免溢出
4. 通过下界/上界剪枝减少无效搜索

本仓库 `18/main.go` 中提供：
- `fourSum`：排序 + 双指针实现
- `fourSumKSum` + `kSum`：通用 kSum 版本
- 主函数内含示例用例并输出结果，便于自检
