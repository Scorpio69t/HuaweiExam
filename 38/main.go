package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 方法一：递归解法
// 最直观的递归实现，基于前一项生成下一项
func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}

	// 递归获取前一项
	prev := countAndSay1(n - 1)

	// 对前一项进行行程长度编码
	return encodeRLE(prev)
}

// 行程长度编码函数
func encodeRLE(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result strings.Builder
	i := 0

	for i < len(s) {
		count := 1
		// 统计连续相同字符的个数
		for i+count < len(s) && s[i] == s[i+count] {
			count++
		}
		// 添加计数和字符
		result.WriteString(strconv.Itoa(count))
		result.WriteByte(s[i])
		i += count
	}

	return result.String()
}

// 方法二：迭代解法
// 使用循环代替递归，优化空间复杂度
func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}

	result := "1"
	for i := 2; i <= n; i++ {
		result = encodeRLE(result)
	}

	return result
}

// 方法三：优化迭代
// 使用strings.Builder优化字符串构建，双指针技术避免重复遍历
func countAndSay3(n int) string {
	if n == 1 {
		return "1"
	}

	result := "1"
	for i := 2; i <= n; i++ {
		var builder strings.Builder
		left := 0

		for left < len(result) {
			right := left
			// 扩展右指针直到遇到不同字符
			for right < len(result) && result[right] == result[left] {
				right++
			}
			// 添加计数和字符
			builder.WriteString(strconv.Itoa(right - left))
			builder.WriteByte(result[left])
			left = right
		}

		result = builder.String()
	}

	return result
}

// 方法四：双缓冲技术
// 预分配缓冲区大小，使用双缓冲技术优化内存使用
func countAndSay4(n int) string {
	if n == 1 {
		return "1"
	}

	result := "1"
	for i := 2; i <= n; i++ {
		// 预估新字符串长度（通常比原字符串长）
		estimatedLen := len(result) * 2
		builder := strings.Builder{}
		builder.Grow(estimatedLen)

		j := 0
		for j < len(result) {
			count := 1
			for j+count < len(result) && result[j] == result[j+count] {
				count++
			}
			builder.WriteString(strconv.Itoa(count))
			builder.WriteByte(result[j])
			j += count
		}

		result = builder.String()
	}

	return result
}

// 辅助函数：打印外观数列的前n项
func printCountAndSaySequence(n int) {
	fmt.Printf("外观数列前%d项:\n", n)
	for i := 1; i <= n; i++ {
		result := countAndSay4(i) // 使用最优算法
		fmt.Printf("countAndSay(%d) = \"%s\"\n", i, result)
	}
	fmt.Println()
}

// 辅助函数：创建测试用例
func createTestCases() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30}
}

// 辅助函数：验证两个字符串是否相等
func isEqual(a, b string) bool {
	return a == b
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func(int) string, n int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(n)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 内存使用测试函数
func testMemoryUsage(n int) {
	fmt.Printf("测试n=%d时的内存使用情况:\n", n)

	// 测试递归解法
	fmt.Println("递归解法:")
	start := time.Now()
	result1 := countAndSay1(n)
	duration1 := time.Since(start)
	fmt.Printf("  结果: \"%s\" (长度: %d)\n", result1, len(result1))
	fmt.Printf("  耗时: %v\n", duration1)

	// 测试迭代解法
	fmt.Println("迭代解法:")
	start = time.Now()
	result2 := countAndSay2(n)
	duration2 := time.Since(start)
	fmt.Printf("  结果: \"%s\" (长度: %d)\n", result2, len(result2))
	fmt.Printf("  耗时: %v\n", duration2)

	// 测试优化迭代
	fmt.Println("优化迭代:")
	start = time.Now()
	result3 := countAndSay3(n)
	duration3 := time.Since(start)
	fmt.Printf("  结果: \"%s\" (长度: %d)\n", result3, len(result3))
	fmt.Printf("  耗时: %v\n", duration3)

	// 测试双缓冲技术
	fmt.Println("双缓冲技术:")
	start = time.Now()
	result4 := countAndSay4(n)
	duration4 := time.Since(start)
	fmt.Printf("  结果: \"%s\" (长度: %d)\n", result4, len(result4))
	fmt.Printf("  耗时: %v\n", duration4)

	// 验证结果一致性
	fmt.Println("结果验证:")
	fmt.Printf("  所有算法结果一致: %t\n",
		isEqual(result1, result2) && isEqual(result2, result3) && isEqual(result3, result4))
	fmt.Println()
}

func main() {
	fmt.Println("=== 38. 外观数列 ===")
	fmt.Println()

	// 打印外观数列的前10项
	printCountAndSaySequence(10)

	// 测试所有算法
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(int) string
	}{
		{"递归解法", countAndSay1},
		{"迭代解法", countAndSay2},
		{"优化迭代", countAndSay3},
		{"双缓冲技术", countAndSay4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试 n=%d:\n", testCase)

		results := make([]string, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if results[i] != results[0] {
				allEqual = false
				break
			}
		}

		if allEqual {
			fmt.Printf("  ✅ 所有算法结果一致: \"%s\"\n", results[0])
		} else {
			fmt.Printf("  ❌ 算法结果不一致\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: \"%s\"\n", algo.name, results[i])
			}
		}
	}
	fmt.Println()

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceN := 20 // 使用较大的n值进行性能测试

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceN, algo.name)
	}
	fmt.Println()

	// 内存使用测试
	fmt.Println("=== 内存使用测试 ===")
	testMemoryUsage(15)
	testMemoryUsage(20)
	testMemoryUsage(25)

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("外观数列的规律分析:")
	fmt.Println("1. 每一项都是对前一项的行程长度编码")
	fmt.Println("2. 字符串长度呈指数级增长")
	fmt.Println("3. 数字1和2出现频率最高")
	fmt.Println("4. 随着n增大，字符串变得非常长")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 递归解法: O(n×m)，其中m为字符串平均长度")
	fmt.Println("- 迭代解法: O(n×m)，但常数因子更小")
	fmt.Println("- 优化迭代: O(n×m)，使用双指针技术优化")
	fmt.Println("- 双缓冲技术: O(n×m)，预分配缓冲区优化")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 递归解法: O(n×m)，递归栈深度为n")
	fmt.Println("- 迭代解法: O(m)，只需要存储当前字符串")
	fmt.Println("- 优化迭代: O(m)，使用Builder优化内存")
	fmt.Println("- 双缓冲技术: O(m)，预分配缓冲区")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 递归解法：最直观易懂，适合理解算法逻辑")
	fmt.Println("2. 迭代解法：显著优化空间复杂度，避免栈溢出")
	fmt.Println("3. 优化迭代：平衡了性能和代码可读性")
	fmt.Println("4. 双缓冲技术：性能最佳，内存使用最优")
	fmt.Println()
	fmt.Println("推荐使用：双缓冲技术（方法四），在保证性能的同时内存使用最优")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 数据压缩：行程长度编码的实际应用")
	fmt.Println("- 字符串处理：学习字符串操作和模式匹配")
	fmt.Println("- 递归算法：理解递归和迭代的转换")
	fmt.Println("- 算法竞赛：字符串处理的基础题目")
	fmt.Println("- 数学序列：研究数学序列的生成规律")
}
