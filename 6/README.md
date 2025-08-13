# 6. Z 字形变换

## 题目描述

将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

## 示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

## 示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I

## 示例 3：

输入：s = "A", numRows = 1
输出："A"

## 提示：

- 1 <= s.length <= 1000
- s 由英文字母（小写和大写）、',' 和 '.' 组成
- 1 <= numRows <= 1000

## 解题思路

这道题要求将字符串按Z字形排列，然后按行读取。这是一个字符串处理的经典问题。

### 算法分析

这道题的核心思想是**模拟Z字形排列过程**，主要解法包括：

1. **模拟法**：实际构建Z字形矩阵，然后按行读取
2. **方向控制法**：使用方向变量控制字符放置的行数
3. **数学规律法**：利用Z字形的数学规律直接计算字符位置

### 问题本质分析

```mermaid
graph TD
    A[Z字形变换] --> B[字符串重排问题]
    B --> C[Z字形排列]
    B --> D[按行读取]
    
    C --> E[垂直向下填充]
    C --> F[斜向向上填充]
    D --> G[逐行连接结果]
    
    E --> H[行索引递增]
    F --> I[行索引递减]
    G --> J[最终输出字符串]
    
    H --> K[到达底部改变方向]
    I --> L[到达顶部改变方向]
    K --> M[方向控制逻辑]
    L --> M
```

### Z字形排列过程详解

```mermaid
flowchart TD
    A[输入字符串和行数] --> B[创建numRows个字符串构建器]
    B --> C[初始化当前行索引和方向]
    C --> D[遍历字符串字符]
    
    D --> E[将字符添加到当前行]
    E --> F[更新行索引]
    F --> G{是否到达边界?}
    
    G -->|是| H[改变方向]
    G -->|否| I[继续下一字符]
    
    H --> I
    I --> J{还有字符?}
    J -->|是| D
    J -->|否| K[按行连接结果]
    
    K --> L[返回最终字符串]
    
    C --> M[currentRow = 0, direction = 1]
    F --> N[currentRow += direction]
    G --> O[currentRow == 0 或 numRows-1]
    H --> P[direction = -direction]
```

### Z字形排列可视化

```mermaid
graph TD
    A["输入: 'PAYPALISHIRING', numRows = 3"] --> B[Z字形排列过程]
    
    B --> C["第1步: P → 第0行"]
    C --> D["第2步: A → 第1行"]
    D --> E["第3步: Y → 第2行"]
    E --> F["第4步: P → 第1行 (方向改变)"]
    F --> G["第5步: A → 第0行 (方向改变)"]
    G --> H["第6步: L → 第1行"]
    H --> I["第7步: I → 第2行"]
    
    I --> J[继续填充...]
    J --> K["最终排列:"]
    K --> L["P   A   H   N"]
    L --> M["A P L S I I G"]
    M --> N["Y   I   R"]
    
    N --> O["按行读取: PAHNAPLSIIGYIR"]
```

### 方向控制策略

```mermaid
graph TD
    A[方向控制策略] --> B[初始状态]
    B --> C[向下填充阶段]
    C --> D[到达底部]
    D --> E[向上填充阶段]
    E --> F[到达顶部]
    
    B --> G[currentRow = 0, direction = 1]
    C --> H[currentRow += 1]
    D --> I[currentRow == numRows-1]
    I --> J[direction = -1]
    E --> K[currentRow -= 1]
    F --> L[currentRow == 0]
    L --> M[direction = 1]
    
    M --> C
    J --> E
```

### 数学规律法详解

```mermaid
graph TD
    A[数学规律法] --> B[周期长度计算]
    B --> C[垂直字符位置]
    B --> D[斜向字符位置]
    
    B --> E[cycleLen = 2*numRows - 2]
    C --> F[第row行: i = row, row+cycleLen, row+2*cycleLen...]
    D --> G[斜向位置: i + cycleLen - 2*row]
    
    E --> H[第一行和最后一行只有垂直字符]
    F --> I[中间行有垂直和斜向字符]
    G --> J[需要检查边界条件]
    
    H --> K[简化处理逻辑]
    I --> L[双重循环填充]
    J --> L
```

### 各种解法对比

```mermaid
graph TD
    A[解法对比] --> B[模拟法]
    A --> C[方向控制法]
    A --> D[数学规律法]
    
    B --> E[时间O_n空间O_n*numRows]
    C --> F[时间O_n空间O_n]
    D --> G[时间O_n空间O_n]
    
    E --> H[直观易懂]
    F --> I[空间效率高]
    G --> J[性能最优]
    
    H --> K[适合理解算法]
    I --> K[推荐实现]
    J --> K[适合优化]
```

### 算法流程图

```mermaid
flowchart TD
    A[开始] --> B{numRows == 1?}
    B -->|是| C[直接返回原字符串]
    B -->|否| D{numRows >= len(s)?}
    
    D -->|是| C
    D -->|否| E[创建numRows个字符串构建器]
    
    E --> F[初始化currentRow=0, direction=1]
    F --> G[遍历字符串字符]
    
    G --> H[将字符添加到当前行]
    H --> I[更新行索引]
    I --> J{到达边界?}
    
    J -->|是| K[改变方向]
    J -->|否| L{还有字符?}
    
    K --> L
    L -->|是| G
    L -->|否| M[按行连接结果]
    
    M --> N[返回最终字符串]
    C --> O[结束]
    N --> O
```

### 边界情况处理

```mermaid
graph TD
    A[边界情况] --> B[numRows = 1]
    A --> C[numRows >= len(s)]
    A --> D[空字符串]
    A --> E[单字符字符串]
    
    B --> F[直接返回原字符串]
    C --> F
    D --> G[返回空字符串]
    E --> H[正常处理]
    
    F --> I[避免不必要的计算]
    G --> I
    H --> I
```

### 时间复杂度分析

```mermaid
graph TD
    A[时间复杂度分析] --> B[字符遍历]
    B --> C[每个字符处理O_1]
    C --> D[总时间O_n]
    
    D --> E[n是字符串长度]
    E --> F[线性时间复杂度]
    F --> G[最优解法]
```

### 空间复杂度分析

```mermaid
graph TD
    A[空间复杂度分析] --> B[额外空间使用]
    B --> C[字符串构建器数组]
    C --> D[空间O_n]
    
    D --> E[n是字符串长度]
    E --> F[空间效率合理]
    F --> G[可接受范围]
```

### 关键优化点

```mermaid
graph TD
    A[优化策略] --> B[提前返回]
    A --> C[字符串构建器]
    A --> D[方向控制优化]
    
    B --> E[边界情况直接返回]
    C --> F[避免字符串拼接开销]
    D --> G[减少条件判断]
    
    E --> H[提高执行效率]
    F --> H
    G --> H
```

### 实际应用场景

```mermaid
graph TD
    A[应用场景] --> B[文本排版]
    A --> C[数据可视化]
    A --> D[密码学]
    A --> E[图像处理]
    
    B --> F[文本分栏显示]
    C --> G[数据表格排列]
    D --> H[字符重排加密]
    E --> I[像素矩阵变换]
    
    F --> J[核心算法组件]
    G --> J
    H --> J
    I --> J
```

### 测试用例设计

```mermaid
graph TD
    A[测试用例] --> B[基础功能]
    A --> C[边界情况]
    A --> D[性能测试]
    
    B --> E[多行Z字形]
    B --> F[单行情况]
    B --> G[不同字符串长度]
    
    C --> H[numRows = 1]
    C --> I[numRows >= len(s)]
    C --> J[空字符串]
    
    D --> K[最大长度字符串]
    D --> L[最大行数]
    
    E --> M[验证正确性]
    F --> M
    G --> M
    H --> M
    I --> M
    J --> M
    K --> N[验证性能]
    L --> N
```

### 代码实现要点

1. **方向控制逻辑**：
   - 使用direction变量控制行索引变化
   - 在边界处改变方向

2. **字符串构建器**：
   - 使用strings.Builder提高效率
   - 避免频繁的字符串拼接

3. **边界条件处理**：
   - numRows = 1时直接返回
   - numRows >= len(s)时直接返回

4. **行索引管理**：
   - 当前行索引范围：0 到 numRows-1
   - 方向改变条件：到达顶部或底部

5. **结果构建**：
   - 按行顺序连接所有行的内容
   - 使用strings.Builder提高效率

这个问题的关键在于**理解Z字形的排列规律**和**掌握方向控制技巧**，通过模拟Z字形的填充过程，实现字符串的重排和重组。
