# 5. 最长回文子串

## 描述

给你一个字符串 s，找到 s 中最长的 回文 子串。

## 示例 1

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

## 示例 2

输入：s = "cbbd"
输出："bb"

## 示例 3

输入：s = "a"
输出："a"

## 示例 4

输入：s = "ac"
输出："a"

## 提示

- 1 <= s.length <= 1000
- s 仅由数字和英文字母组成

## 解题思路

### 方法一：中心扩展法（推荐）

**核心思想**：
- 以每个字符为中心，向两边扩展寻找回文
- 考虑奇数长度和偶数长度的回文

**算法步骤**：
1. 遍历字符串的每个位置
2. 以当前位置为中心，向两边扩展
3. 分别处理奇数长度和偶数长度的回文
4. 记录最长回文的起始位置和长度

**时间复杂度**：O(n²)
**空间复杂度**：O(1)

### 方法二：动态规划

**核心思想**：
- 使用dp[i][j]表示s[i:j+1]是否为回文
- 状态转移：dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]

**算法步骤**：
1. 创建二维dp数组
2. 初始化长度为1和2的回文
3. 按长度递增填充dp数组
4. 记录最长回文的位置

**时间复杂度**：O(n²)
**空间复杂度**：O(n²)

### 方法三：Manacher算法

**核心思想**：
- 使用马拉车算法在线性时间内找到最长回文
- 利用回文的对称性质优化计算

**算法步骤**：
1. 预处理字符串，插入分隔符
2. 维护回文半径数组
3. 利用对称性质优化扩展
4. 找到最大回文半径

**时间复杂度**：O(n)
**空间复杂度**：O(n)

### 方法四：暴力解法

**核心思想**：
- 枚举所有可能的子串
- 检查每个子串是否为回文

**算法步骤**：
1. 双重循环枚举所有子串
2. 检查每个子串是否为回文
3. 记录最长回文

**时间复杂度**：O(n³)
**空间复杂度**：O(1)

## 代码实现

```go
// 方法一：中心扩展法
func longestPalindrome1(s string) string {
    if len(s) < 2 {
        return s
    }
    
    start, maxLen := 0, 1
    
    for i := 0; i < len(s); i++ {
        // 奇数长度回文
        len1 := expandAroundCenter(s, i, i)
        // 偶数长度回文
        len2 := expandAroundCenter(s, i, i+1)
        
        maxLenCur := max(len1, len2)
        if maxLenCur > maxLen {
            start = i - (maxLenCur-1)/2
            maxLen = maxLenCur
        }
    }
    
    return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}
```

## 复杂度分析

- **时间复杂度**：O(n²)，其中n是字符串长度
- **空间复杂度**：O(1)，只使用常数个变量

## 测试用例

```go
func main() {
    // 测试用例1
    s1 := "babad"
    fmt.Printf("测试用例1: s=%s, 结果=%s\n", s1, longestPalindrome1(s1))
    
    // 测试用例2
    s2 := "cbbd"
    fmt.Printf("测试用例2: s=%s, 结果=%s\n", s2, longestPalindrome1(s2))
    
    // 测试用例3
    s3 := "a"
    fmt.Printf("测试用例3: s=%s, 结果=%s\n", s3, longestPalindrome1(s3))
}
```

