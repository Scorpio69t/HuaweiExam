package main

import (
	"fmt"
)

// longestValidParenthesesStack 栈（索引栈）解法
// 思路：维护一个索引栈，栈底为“最后一个不可能成为有效子串起点”的下标哨兵（初始为 -1）。
// 遍历字符：
// - 遇到 '(' 下标压栈
// - 遇到 ')' 弹栈；若栈空，压入当前下标作为新的无效边界；否则用 i - 栈顶 计算当前最大长度
// 时间 O(n)，空间 O(n)
func longestValidParenthesesStack(s string) int {
	maxLen := 0
	stack := make([]int, 0, len(s))
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// pop
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 当前 i 作为新的无效边界
				stack = append(stack, i)
			} else {
				// 有效长度 = i - 栈顶
				curr := i - stack[len(stack)-1]
				if curr > maxLen {
					maxLen = curr
				}
			}
		}
	}
	return maxLen
}

// longestValidParenthesesDP 动态规划
// dp[i] 表示“以 i 位置字符结尾”的最长有效括号长度
// 转移：
//   - 若 s[i] == ')' 且 s[i-1] == '('，则 dp[i] = dp[i-2] + 2
//   - 若 s[i] == ')' 且 s[i-1] == ')'，并且 i - dp[i-1] - 1 >= 0 且 s[i - dp[i-1] - 1] == '('，
//     则 dp[i] = dp[i-1] + 2 + dp[i - dp[i-1] - 2]
//
// 时间 O(n)，空间 O(n)
func longestValidParenthesesDP(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	maxLen := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				val := 2
				if i >= 2 {
					val += dp[i-2]
				}
				dp[i] = val
			} else {
				// s[i-1] == ')'
				pre := i - dp[i-1] - 1
				if pre >= 0 && s[pre] == '(' {
					val := dp[i-1] + 2
					if pre-1 >= 0 {
						val += dp[pre-1]
					}
					dp[i] = val
				}
			}
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}
	return maxLen
}

// longestValidParenthesesTwoPass 计数双向扫描
// 从左到右统计 left/right 个数，相等时更新长度，right > left 时清零；
// 再从右到左做同样逻辑（对称处理 "(((" 这类前缀过多左括号的情况）。
// 时间 O(n)，空间 O(1)
func longestValidParenthesesTwoPass(s string) int {
	maxLen := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			curr := 2 * right
			if curr > maxLen {
				maxLen = curr
			}
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			curr := 2 * left
			if curr > maxLen {
				maxLen = curr
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLen
}

// longestValidParenthesesGreedy 扩展：在线修正法（基于索引边界）
// 与栈法等价思路：维护最后一个不匹配右括号的位置 lastInvalid，遇到 '(' 入栈索引；
// 遇到 ')' 出栈，若栈空更新 lastInvalid，否则用 i - 栈顶计算长度。
func longestValidParenthesesGreedy(s string) int {
	maxLen := 0
	stack := make([]int, 0, len(s))
	lastInvalid := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				lastInvalid = i
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					curr := i - lastInvalid
					if curr > maxLen {
						maxLen = curr
					}
				} else {
					curr := i - stack[len(stack)-1]
					if curr > maxLen {
						maxLen = curr
					}
				}
			}
		}
	}
	return maxLen
}

// 统一封装，便于批量验证
type method struct {
	name string
	fn   func(string) int
}

func main() {
	tests := []string{
		"",
		"(",
		")",
		"()",
		"(()",
		")()())",
		"()()",
		"()(())",
		"())((())",
		"((((",
		"))))",
		"()(()",
		"(()())",
		"())()(()())",
	}

	methods := []method{
		{"索引栈", longestValidParenthesesStack},
		{"动态规划", longestValidParenthesesDP},
		{"双向计数", longestValidParenthesesTwoPass},
		{"在线修正", longestValidParenthesesGreedy},
	}

	fmt.Println("32. 最长有效括号 - 多解法对比")
	for _, t := range tests {
		fmt.Printf("输入: %q\n", t)
		best := 0
		for _, m := range methods {
			got := m.fn(t)
			if got > best {
				best = got
			}
			fmt.Printf("  %-6s => %d\n", m.name, got)
		}
		fmt.Printf("  最优 => %d\n", best)
		fmt.Println("------------------------------")
	}
}
