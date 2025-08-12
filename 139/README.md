# 139. 单词拆分

## 描述

给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

 
## 示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

## 示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。

## 示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
 
## 提示：

- 1 <= s.length <= 300
- 1 <= wordDict.length <= 1000
- 1 <= wordDict[i].length <= 20
- s 和 wordDict[i] 仅由小写英文字母组成
- wordDict 中的所有字符串 互不相同

## 解题思路

### 算法分析

这是一道经典的**动态规划**问题，核心在于判断字符串能否通过字典中的单词进行分割组合。

#### 核心思想

1. **状态定义**：dp[i] 表示字符串 s[0...i-1] 能否被字典中的单词拆分
2. **状态转移**：如果 dp[j] = true 且 s[j...i-1] 在字典中，则 dp[i] = true
3. **边界条件**：dp[0] = true（空字符串可以被拆分）
4. **目标结果**：dp[n]（整个字符串是否可拆分）

#### 算法对比

| 算法 | 时间复杂度 | 空间复杂度 | 特点 |
|------|------------|------------|------|
| 动态规划 | O(n²×m) | O(n) | 经典解法，易理解 |
| DFS+记忆化 | O(n²×m) | O(n) | 递归思维，避免重复计算 |
| BFS | O(n²×m) | O(n) | 层次遍历，直观 |
| Trie+DP | O(n×k+∑len) | O(∑len) | 字典较大时优化查找 |

注：n为字符串长度，m为字典大小，k为平均单词长度

### 算法流程图

```mermaid
graph TD
    A[输入: 字符串s, 字典wordDict] --> B[初始化dp数组]
    B --> C[dp[0] = true 空字符串可拆分]
    C --> D[遍历位置i: 1到n]
    D --> E[遍历分割点j: 0到i-1]
    E --> F{dp[j] == true?}
    F -->|否| G[继续下一个j]
    F -->|是| H{s[j:i] 在字典中?}
    H -->|否| G
    H -->|是| I[dp[i] = true]
    I --> J[break 内层循环]
    G --> K{j < i-1?}
    K -->|是| E
    K -->|否| L{i < n?}
    J --> L
    L -->|是| D
    L -->|否| M[返回dp[n]]
```

### DFS+记忆化算法流程

```mermaid
graph TD
    A[DFS开始: 位置start] --> B{start == len(s)?}
    B -->|是| C[返回true 完全匹配]
    B -->|否| D{memo[start] 已计算?}
    D -->|是| E[返回memo[start]]
    D -->|否| F[遍历从start开始的所有子串]
    F --> G{当前子串在字典中?}
    G -->|否| H[尝试下一个子串]
    G -->|是| I[递归DFS子串结束位置]
    I --> J{递归返回true?}
    J -->|是| K[memo[start] = true]
    J -->|否| H
    K --> L[返回true]
    H --> M{还有子串?}
    M -->|是| F
    M -->|否| N[memo[start] = false]
    N --> O[返回false]
```

### BFS算法流程

```mermaid
graph TD
    A[BFS开始] --> B[队列初始化: queue = [0]]
    B --> C[visited数组标记已访问位置]
    C --> D{队列非空?}
    D -->|否| E[返回false 无法拆分]
    D -->|是| F[取出队列首元素start]
    F --> G{start == len(s)?}
    G -->|是| H[返回true 完全匹配]
    G -->|否| I[遍历从start开始的子串]
    I --> J{子串在字典中?}
    J -->|否| K[尝试下一个子串]
    J -->|是| L{end位置已访问?}
    L -->|是| K
    L -->|否| M[标记end已访问]
    M --> N[end加入队列]
    N --> K
    K --> O{还有子串?}
    O -->|是| I
    O -->|否| D
```

### Trie优化算法流程

```mermaid
graph TD
    A[构建Trie树] --> B[插入所有字典单词]
    B --> C[DP算法开始]
    C --> D[dp[0] = true]
    D --> E[遍历位置i: 1到n]
    E --> F[从Trie根节点开始]
    F --> G[向前查找: j从i-1到0]
    G --> H{s[j]对应的子节点存在?}
    H -->|否| I[break 无法继续]
    H -->|是| J[移动到子节点]
    J --> K{当前节点是单词结尾?}
    K -->|否| L[继续向前查找]
    K -->|是| M{dp[j] == true?}
    M -->|否| L
    M -->|是| N[dp[i] = true]
    N --> O[break 找到一种拆分]
    L --> P{j > 0?}
    P -->|是| G
    P -->|否| Q{i < n?}
    O --> Q
    I --> Q
    Q -->|是| E
    Q -->|否| R[返回dp[n]]
```

### 复杂度分析

#### 时间复杂度对比
- **动态规划**：O(n²×m)，需要双重循环+字典查找
- **DFS+记忆化**：O(n²×m)，避免重复子问题计算
- **BFS**：O(n²×m)，广度优先遍历所有可能
- **Trie+DP**：O(n²×k)，Trie优化字典查找

#### 空间复杂度对比
- **动态规划**：O(n)，dp数组
- **DFS+记忆化**：O(n)，递归栈+记忆化数组
- **BFS**：O(n)，队列+访问标记
- **Trie+DP**：O(∑word_len)，Trie树存储

### 应用场景

1. **文本分词**：中文分词、英文单词切分
2. **密码学**：密码字典攻击，模式匹配
3. **编译原理**：词法分析，token识别
4. **搜索引擎**：查询词分割，同义词匹配
5. **自然语言处理**：句子解析，语义分析

### 算法优化

#### 1. 哈希集合优化
```go
// 将wordDict转为map，O(1)查找
wordSet := make(map[string]bool)
for _, word := range wordDict {
    wordSet[word] = true
}
```

#### 2. 早期剪枝
```go
// 如果剩余长度小于最短单词，直接返回false
if len(s)-start < minWordLen {
    return false
}
```

#### 3. 长度过滤
```go
// 只考虑可能的单词长度
for length := minLen; length <= maxLen && start+length <= len(s); length++ {
    // 检查长度为length的子串
}
```

#### 4. 记忆化优化
```go
// 使用memo避免重复计算
if val, exists := memo[start]; exists {
    return val
}
```

### 测试用例设计

#### 基础测试
- 正常拆分：`"leetcode"`, `["leet", "code"]`
- 重复使用：`"applepenapple"`, `["apple", "pen"]`
- 无法拆分：`"catsandog"`, `["cats", "dog", "sand", "and", "cat"]`

#### 边界测试
- 空字符串：`""`, `["a"]`
- 单字符：`"a"`, `["a"]`
- 完全匹配：`"word"`, `["word"]`
- 字典为空：`"abc"`, `[]`

#### 性能测试
- 长字符串：300字符的字符串
- 大字典：1000个单词的字典
- 最坏情况：每个位置都需要回溯

### 实现技巧

1. **状态压缩**：只需要一维dp数组
2. **逆向思维**：从后往前考虑剩余部分
3. **剪枝优化**：提前终止无效分支
4. **数据结构选择**：根据字典大小选择HashMap或Trie
5. **边界处理**：正确处理空字符串和越界情况

