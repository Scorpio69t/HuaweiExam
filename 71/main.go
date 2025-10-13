package main

import (
	"fmt"
	"strings"
)

// =========================== 方法一：栈+字符串分割（最优解法） ===========================

// simplifyPath 栈+字符串分割
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func simplifyPath(path string) string {
	// 按'/'分割路径
	parts := strings.Split(path, "/")

	// 使用栈存储有效目录
	stack := make([]string, 0)

	for _, part := range parts {
		if part == "" || part == "." {
			// 跳过空字符串和当前目录
			continue
		} else if part == ".." {
			// 上级目录，弹出栈（如果不为空）
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			// 普通目录或特殊目录名（如...）
			stack = append(stack, part)
		}
	}

	// 构建结果
	if len(stack) == 0 {
		return "/"
	}

	return "/" + strings.Join(stack, "/")
}

// =========================== 方法二：双指针模拟（原地处理） ===========================

// simplifyPath2 双指针模拟
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func simplifyPath2(path string) string {
	parts := strings.Split(path, "/")
	j := 0 // 写指针

	for i := 0; i < len(parts); i++ {
		if parts[i] == "" || parts[i] == "." {
			continue
		} else if parts[i] == ".." {
			if j > 0 {
				j--
			}
		} else {
			parts[j] = parts[i]
			j++
		}
	}

	if j == 0 {
		return "/"
	}

	return "/" + strings.Join(parts[:j], "/")
}

// =========================== 方法三：strings.Builder优化 ===========================

// simplifyPath3 strings.Builder优化
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func simplifyPath3(path string) string {
	parts := strings.Split(path, "/")
	stack := make([]string, 0)

	for _, part := range parts {
		if part == "" || part == "." {
			continue
		} else if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	// 使用StringBuilder
	var builder strings.Builder
	if len(stack) == 0 {
		return "/"
	}

	for _, dir := range stack {
		builder.WriteByte('/')
		builder.WriteString(dir)
	}

	return builder.String()
}

// =========================== 方法四：递归处理 ===========================

// simplifyPath4 递归处理
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func simplifyPath4(path string) string {
	parts := strings.Split(path, "/")
	result := simplifyHelper(parts, 0, []string{})

	if len(result) == 0 {
		return "/"
	}

	return "/" + strings.Join(result, "/")
}

// simplifyHelper 递归辅助函数
func simplifyHelper(parts []string, index int, stack []string) []string {
	if index >= len(parts) {
		return stack
	}

	part := parts[index]

	if part == "" || part == "." {
		return simplifyHelper(parts, index+1, stack)
	} else if part == ".." {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
		return simplifyHelper(parts, index+1, stack)
	} else {
		return simplifyHelper(parts, index+1, append(stack, part))
	}
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 71: 简化路径 ===\n")

	// 测试用例
	testCases := []struct {
		path   string
		expect string
	}{
		{"/home/", "/home"},          // 示例1: 删除尾随斜杠
		{"/home//foo/", "/home/foo"}, // 示例2: 连续斜杠
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"}, // 示例3: 上级目录
		{"/../", "/"},                         // 示例4: 根目录的..
		{"/.../a/../b/c/../d/./", "/.../b/d"}, // 示例5: ...是有效目录名
		{"/", "/"},                            // 边界: 根目录
		{"/a/./b/../../c/", "/c"},             // 复杂情况
		{"/a/../../b/../c//.//", "/c"},        // 多个连续操作
		{"/a//b////c/d//././/..", "/a/b/c"},   // 各种组合
		{"/...", "/..."},                      // ...是目录名
		{"/../../../", "/"},                   // 多个..在根目录
		{"/home/foo/..", "/home"},             // 普通回退
	}

	fmt.Println("方法一：栈+字符串分割")
	runTests(testCases, simplifyPath)

	fmt.Println("\n方法二：双指针模拟")
	runTests(testCases, simplifyPath2)

	fmt.Println("\n方法三：strings.Builder优化")
	runTests(testCases, simplifyPath3)

	fmt.Println("\n方法四：递归处理")
	runTests(testCases, simplifyPath4)

	// 详细示例
	fmt.Println("\n=== 详细示例 ===")
	detailedExample()

	// 路径规则演示
	fmt.Println("\n=== Unix路径规则演示 ===")
	pathRulesDemo()
}

// runTests 运行测试用例
func runTests(testCases []struct {
	path   string
	expect string
}, fn func(string) string) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.path)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}

		fmt.Printf("  测试%d: %s ", i+1, status)
		if status == "❌" {
			fmt.Printf("输入=\"%s\", 输出=\"%s\", 期望=\"%s\"\n", tc.path, result, tc.expect)
		} else {
			fmt.Printf("\"%s\" → \"%s\"\n", tc.path, result)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

// detailedExample 详细示例
func detailedExample() {
	path := "/home/user/Documents/../Pictures"

	fmt.Printf("输入路径: \"%s\"\n\n", path)

	// 分割演示
	parts := strings.Split(path, "/")
	fmt.Println("步骤1: 按'/'分割路径")
	fmt.Printf("  parts = %v\n\n", parts)

	// 栈处理演示
	fmt.Println("步骤2: 逐个处理目录")
	stack := make([]string, 0)

	for i, part := range parts {
		fmt.Printf("  [%d] part=\"%s\"", i, part)

		if part == "" || part == "." {
			fmt.Println(" → 跳过（空或.）")
		} else if part == ".." {
			if len(stack) > 0 {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Printf(" → 上级目录，弹出\"%s\"，栈=%v\n", popped, stack)
			} else {
				fmt.Println(" → 上级目录但栈为空，忽略")
			}
		} else {
			stack = append(stack, part)
			fmt.Printf(" → 压入栈，栈=%v\n", stack)
		}
	}

	// 构建结果
	fmt.Println("\n步骤3: 构建结果")
	result := "/" + strings.Join(stack, "/")
	fmt.Printf("  result = \"%s\"\n", result)
}

// pathRulesDemo 路径规则演示
func pathRulesDemo() {
	examples := []struct {
		desc string
		path string
	}{
		{"删除尾随斜杠", "/home/"},
		{"连续斜杠合并", "/home//foo/"},
		{"当前目录.", "/a/./b"},
		{"上级目录..", "/a/b/../c"},
		{"根目录的..", "/../"},
		{"...是目录名", "/.../"},
		{"复杂组合", "/a/./b/../../c/"},
	}

	for _, ex := range examples {
		result := simplifyPath(ex.path)
		fmt.Printf("%-20s: \"%s\" → \"%s\"\n", ex.desc, ex.path, result)
	}

	// 规则总结
	fmt.Println("\n规则总结:")
	fmt.Println("  1. '.'  表示当前目录，忽略")
	fmt.Println("  2. '..' 表示上级目录，回退一级")
	fmt.Println("  3. '//' 多个斜杠视为单个'/'")
	fmt.Println("  4. '...' 是有效目录名")
	fmt.Println("  5. 结果以'/'开头，不以'/'结尾（根除外）")
}
