# 108. 将有序数组转换为二叉搜索树

## 题目描述

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。



## 示例 1：

![btree1](./images/btree1.jpg)

输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

![btree2](./images/btree2.jpg)

## 示例 2：

![btree](./images/btree.jpg)

输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。


## 提示：

- 1 <= nums.length <= 10^4
- -10^4 <= nums[i] <= 10^4
- nums 按 严格递增 顺序排列

## 解题思路

### 问题分析

这是一道经典的**二叉搜索树(BST)构建**问题。关键点：

1. **输入特点**：数组已经按**升序排列**
2. **输出要求**：构建**平衡**二叉搜索树
3. **核心洞察**：有序数组的**中序遍历**恰好是BST的中序遍历结果

### 算法思想

由于数组已排序，我们可以利用**分治思想**：

1. **选择中间元素**作为根节点（保证平衡）
2. **左半部分**递归构建左子树
3. **右半部分**递归构建右子树

这样构建的树自然满足：
- **BST性质**：左子树 < 根 < 右子树
- **平衡性**：左右子树高度差不超过1

### 方法一：递归（中间偏左）

**核心思想**：每次选择中间位置（偏左）作为根节点

```mermaid
graph TD
    A[开始: nums=[-10,-3,0,5,9]] --> B[选择中间索引 mid=2, 值=0]
    B --> C[0作为根节点]
    C --> D[左子数组: -10,-3]
    C --> E[右子数组: 5,9]
    D --> F[选择mid=0, 值=-3]
    F --> G[-3作为左子树根]
    G --> H[左: -10]
    G --> I[右: null]
    E --> J[选择mid=0, 值=5]
    J --> K[5作为右子树根]
    K --> L[左: null]
    K --> M[右: 9]
    
    style A fill:#e1f5ff
    style C fill:#bbdefb
    style G fill:#90caf9
    style K fill:#90caf9
```

**算法步骤**：

```mermaid
flowchart TD
    Start([输入: nums, left, right]) --> Check{left > right?}
    Check -->|是| ReturnNil[返回 nil]
    Check -->|否| CalcMid[计算中间索引<br/>mid = left + right / 2]
    CalcMid --> CreateNode[创建根节点<br/>root = &TreeNode{Val: nums[mid]}]
    CreateNode --> BuildLeft[递归构建左子树<br/>root.Left = sortedArrayToBST nums, left, mid-1]
    BuildLeft --> BuildRight[递归构建右子树<br/>root.Right = sortedArrayToBST nums, mid+1, right]
    BuildRight --> Return([返回 root])
    ReturnNil --> End([结束])
    Return --> End
    
    style Start fill:#e8f5e9
    style Check fill:#fff3e0
    style CreateNode fill:#bbdefb
    style BuildLeft fill:#c8e6c9
    style BuildRight fill:#c8e6c9
    style Return fill:#c5cae9
```

**时间复杂度**：O(n) - 每个节点访问一次  
**空间复杂度**：O(log n) - 递归栈深度

### 方法二：递归（中间偏右）

**核心思想**：选择中间偏右的位置作为根节点

```go
mid := left + (right - left + 1) / 2
```

这样会生成另一种平衡的BST结构。

### 方法三：迭代法（栈模拟）

**核心思想**：使用栈模拟递归过程

```mermaid
flowchart TD
    Start([输入: nums]) --> Init[初始化栈<br/>压入 0, len-1, nil, true]
    Init --> Loop{栈非空?}
    Loop -->|否| End([返回根节点])
    Loop -->|是| Pop[弹出栈顶<br/>left, right, parent, isLeft]
    Pop --> Check{left > right?}
    Check -->|是| Loop
    Check -->|否| CalcMid[mid = left + right / 2]
    CalcMid --> CreateNode[创建节点<br/>node = &TreeNode{Val: nums[mid]}]
    CreateNode --> CheckParent{parent == nil?}
    CheckParent -->|是| SetRoot[设置为根节点]
    CheckParent -->|否| CheckLeft{isLeft?}
    CheckLeft -->|是| SetLeft[parent.Left = node]
    CheckLeft -->|否| SetRight[parent.Right = node]
    SetRoot --> PushRight
    SetLeft --> PushRight
    SetRight --> PushRight[压入右子树参数<br/>mid+1, right, node, false]
    PushRight --> PushLeft[压入左子树参数<br/>left, mid-1, node, true]
    PushLeft --> Loop
    
    style Start fill:#e8f5e9
    style CreateNode fill:#bbdefb
    style Loop fill:#fff3e0
    style End fill:#c5cae9
```

**时间复杂度**：O(n)  
**空间复杂度**：O(log n)

### 方法四：中序遍历模拟

**核心思想**：按中序遍历的顺序构建BST

这种方法需要提前知道树的大小，然后在中序遍历位置填充节点值。

### 复杂度对比

| 方法 | 时间复杂度 | 空间复杂度 | 特点 |
|------|-----------|-----------|------|
| 递归（中间偏左） | O(n) | O(log n) | 代码简洁，易理解 |
| 递归（中间偏右） | O(n) | O(log n) | 生成不同的平衡树 |
| 迭代法 | O(n) | O(log n) | 避免递归栈溢出 |
| 中序遍历 | O(n) | O(log n) | 模拟BST构建过程 |

### 关键点总结

1. **选择中点**：保证平衡性的关键
2. **递归边界**：left > right 时返回 nil
3. **分治思想**：问题规模减半
4. **BST性质**：利用数组有序性

### 扩展问题

1. 如何验证构建的树是否平衡？
2. 如何构建最小高度的BST？
3. 如果数组有重复元素怎么处理？
4. 如何构建完全二叉搜索树？

### 相关题目

- **LeetCode 109**：有序链表转换二叉搜索树
- **LeetCode 110**：平衡二叉树
- **LeetCode 1382**：将二叉搜索树变平衡
