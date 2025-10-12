package main

import (
	"fmt"
	"regexp"
	"time"
)

// 方法一：有限状态机算法（最优解法）
func isNumber1(s string) bool {
	// 定义状态转移
	// 状态: 0-初始, 1-符号, 2-整数, 3-小数点(前无数字), 4-小数, 5-指数, 6-指数符号, 7-指数数字
	state := 0

	for _, ch := range s {
		switch state {
		case 0: // 初始状态
			if ch == '+' || ch == '-' {
				state = 1 // 符号
			} else if ch >= '0' && ch <= '9' {
				state = 2 // 整数
			} else if ch == '.' {
				state = 3 // 小数点
			} else {
				return false
			}
		case 1: // 符号后
			if ch >= '0' && ch <= '9' {
				state = 2
			} else if ch == '.' {
				state = 3
			} else {
				return false
			}
		case 2: // 整数部分
			if ch >= '0' && ch <= '9' {
				state = 2
			} else if ch == '.' {
				state = 4 // 小数
			} else if ch == 'e' || ch == 'E' {
				state = 5 // 指数
			} else {
				return false
			}
		case 3: // 小数点（前无数字）
			if ch >= '0' && ch <= '9' {
				state = 4
			} else {
				return false
			}
		case 4: // 小数部分
			if ch >= '0' && ch <= '9' {
				state = 4
			} else if ch == 'e' || ch == 'E' {
				state = 5
			} else {
				return false
			}
		case 5: // 指数符号后
			if ch == '+' || ch == '-' {
				state = 6
			} else if ch >= '0' && ch <= '9' {
				state = 7
			} else {
				return false
			}
		case 6: // 指数符号后的符号
			if ch >= '0' && ch <= '9' {
				state = 7
			} else {
				return false
			}
		case 7: // 指数数字
			if ch >= '0' && ch <= '9' {
				state = 7
			} else {
				return false
			}
		}
	}

	// 只有状态2, 4, 7是接受状态
	return state == 2 || state == 4 || state == 7
}

// 方法二：正则表达式算法
func isNumber2(s string) bool {
	// 构建正则表达式
	// 基数部分: [+-]?([0-9]+\.?[0-9]*|\.[0-9]+)
	// 指数部分: ([eE][+-]?[0-9]+)?
	pattern := `^[+-]?([0-9]+\.?[0-9]*|\.[0-9]+)([eE][+-]?[0-9]+)?$`

	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

// 方法三：分段验证算法
func isNumber3(s string) bool {
	// 找到e或E的位置
	ePos := -1
	for i := 0; i < len(s); i++ {
		if s[i] == 'e' || s[i] == 'E' {
			if ePos != -1 {
				return false // 多个e/E
			}
			ePos = i
		}
	}

	if ePos == -1 {
		// 没有指数，验证整个字符串为基数
		return isValidBase(s)
	}

	// 有指数，分别验证基数和指数
	if ePos == 0 || ePos == len(s)-1 {
		return false // e/E不能在开头或结尾
	}

	base := s[:ePos]
	exponent := s[ePos+1:]

	return isValidBase(base) && isValidInteger(exponent)
}

// 验证基数（整数或小数）
func isValidBase(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0
	// 处理符号
	if s[i] == '+' || s[i] == '-' {
		i++
	}

	if i >= len(s) {
		return false
	}

	hasDigit := false
	hasDot := false

	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			hasDigit = true
		} else if s[i] == '.' {
			if hasDot {
				return false // 多个小数点
			}
			hasDot = true
		} else {
			return false // 非法字符
		}
		i++
	}

	return hasDigit
}

// 验证整数
func isValidInteger(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0
	// 处理符号
	if s[i] == '+' || s[i] == '-' {
		i++
	}

	if i >= len(s) {
		return false
	}

	// 必须全是数字
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		i++
	}

	return true
}

// 方法四：标志位遍历算法
func isNumber4(s string) bool {
	if len(s) == 0 {
		return false
	}

	hasNum := false       // 是否有数字
	hasDot := false       // 是否有小数点
	hasE := false         // 是否有指数符号
	hasNumAfterE := false // 指数后是否有数字

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			hasNum = true
			if hasE {
				hasNumAfterE = true
			}
		} else if ch == '.' {
			// 小数点不能在指数后，且只能有一个
			if hasE || hasDot {
				return false
			}
			hasDot = true
		} else if ch == 'e' || ch == 'E' {
			// 指数前必须有数字，且只能有一个指数
			if !hasNum || hasE {
				return false
			}
			hasE = true
		} else if ch == '+' || ch == '-' {
			// 符号只能在开头或指数符号后
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false // 非法字符
		}
	}

	// 必须有数字，且如果有指数，指数后也必须有数字
	return hasNum && (!hasE || hasNumAfterE)
}

// 测试用例
func createTestCases() []struct {
	input    string
	expected bool
	name     string
} {
	return []struct {
		input    string
		expected bool
		name     string
	}{
		// 有效数字
		{"0", true, "基础: 单个数字"},
		{"2", true, "基础: 整数"},
		{"+100", true, "基础: 正号整数"},
		{"-456", true, "基础: 负号整数"},
		{"0.1", true, "小数: 标准小数"},
		{"3.14", true, "小数: 圆周率"},
		{"+3.14", true, "小数: 正号小数"},
		{"-0.1", true, "小数: 负号小数"},
		{"4.", true, "小数: 整数后跟小数点"},
		{".9", true, "小数: 小数点开头"},
		{"-.9", true, "小数: 负号小数点开头"},
		{"2e10", true, "科学: 整数指数"},
		{"2E10", true, "科学: 大写E"},
		{"3e+7", true, "科学: 正指数"},
		{"6e-1", true, "科学: 负指数"},
		{"53.5e93", true, "科学: 小数指数"},
		{"-123.456e789", true, "科学: 负小数大指数"},
		{"0089", true, "边界: 前导零"},
		{"+.8", true, "边界: 符号小数点"},

		// 无效数字
		{"abc", false, "无效: 纯字母"},
		{"1a", false, "无效: 数字字母"},
		{"1 ", false, "无效: 有空格"},
		{"1e", false, "无效: 指数无数字"},
		{"e3", false, "无效: 开头是指数"},
		{"99e2.5", false, "无效: 指数是小数"},
		{"--6", false, "无效: 双负号"},
		{"-+3", false, "无效: 负正号"},
		{"+-3", false, "无效: 正负号"},
		{"95a54e53", false, "无效: 包含字母"},
		{".", false, "无效: 仅小数点"},
		{"e", false, "无效: 仅指数符号"},
		{"1e2e3", false, "无效: 多个指数"},
		{"1.2.3", false, "无效: 多个小数点"},
		{"+", false, "无效: 仅正号"},
		{"-", false, "无效: 仅负号"},
		{"1+2", false, "无效: 中间有符号"},
		{"1-2", false, "无效: 中间有负号"},
		{"1e+", false, "无效: 指数后仅符号"},
		{"", false, "边界: 空字符串"},
	}
}

// 性能测试
func benchmarkAlgorithm(algorithm func(string) bool, input string, name string) {
	iterations := 10000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(input)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)
	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

func main() {
	fmt.Println("=== 65. 有效数字 ===")
	fmt.Println()

	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(string) bool
	}{
		{"有限状态机算法", isNumber1},
		{"正则表达式算法", isNumber2},
		{"分段验证算法", isNumber3},
		{"标志位遍历算法", isNumber4},
	}

	// 正确性测试
	fmt.Println("=== 算法正确性测试 ===")
	passCount := 0
	failCount := 0

	for _, testCase := range testCases {
		results := make([]bool, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.input)
		}

		// 检查所有算法结果是否一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if results[i] != results[0] {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		isValid := results[0] == testCase.expected

		if allEqual && isValid {
			passCount++
			fmt.Printf("✅ %s: \"%s\" = %v\n", testCase.name, testCase.input, results[0])
		} else {
			failCount++
			fmt.Printf("❌ %s: \"%s\"\n", testCase.name, testCase.input)
			fmt.Printf("   预期: %v\n", testCase.expected)
			for i, algo := range algorithms {
				fmt.Printf("   %s: %v\n", algo.name, results[i])
			}
		}
	}

	fmt.Println()
	fmt.Printf("测试统计: 通过 %d/%d, 失败 %d/%d\n",
		passCount, len(testCases), failCount, len(testCases))
	fmt.Println()

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	perfTests := []string{
		"53.5e93",
		"-123.456e789",
		"+3.14",
		"0.0089",
	}

	for _, test := range perfTests {
		fmt.Printf("测试输入: \"%s\"\n", test)
		for _, algo := range algorithms {
			benchmarkAlgorithm(algo.fn, test, algo.name)
		}
		fmt.Println()
	}

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("有效数字问题的特点:")
	fmt.Println("1. 有限状态机：状态转移清晰")
	fmt.Println("2. 正则表达式：代码简洁")
	fmt.Println("3. 分段验证：拆分基数和指数")
	fmt.Println("4. 标志位遍历：直观但复杂")
	fmt.Println()

	fmt.Println("=== 状态机示例 ===")
	fmt.Println("状态转移:")
	fmt.Println("初始 → 符号/数字/小数点")
	fmt.Println("整数 → 数字/小数点/指数")
	fmt.Println("小数 → 数字/指数")
	fmt.Println("指数 → 符号/数字")
	fmt.Println()
	fmt.Println("接受状态: 整数(2), 小数(4), 指数数字(7)")
	fmt.Println()

	fmt.Println("=== 常见陷阱 ===")
	fmt.Println("1. 小数点前后必须有数字: \".\" 无效")
	fmt.Println("2. 指数部分不能是小数: \"1e2.3\" 无效")
	fmt.Println("3. 符号只能在开头或指数后: \"1+2\" 无效")
	fmt.Println("4. 指数符号后必须有数字: \"1e\" 无效")
	fmt.Println("5. 不能有多个小数点或指数: \"1.2.3\", \"1e2e3\" 无效")
	fmt.Println()

	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 有限状态机: O(n)，每个字符处理一次")
	fmt.Println("- 正则表达式: O(n)，正则匹配")
	fmt.Println("- 分段验证: O(n)，遍历+验证")
	fmt.Println("- 标志位遍历: O(n)，单次遍历")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 有限状态机: O(1)，常数状态变量")
	fmt.Println("- 正则表达式: O(1)，常数空间")
	fmt.Println("- 分段验证: O(1)，常数标志")
	fmt.Println("- 标志位遍历: O(1)，常数标志")
	fmt.Println()

	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 状态机设计：画出状态转移图")
	fmt.Println("2. 接受状态：明确有效的结束状态")
	fmt.Println("3. 边界检查：每个转移都要验证")
	fmt.Println("4. 标志管理：跟踪关键信息")
	fmt.Println("5. 测试覆盖：准备充分的测试用例")
	fmt.Println()

	fmt.Println("=== 应用场景 ===")
	fmt.Println("1. 编译器词法分析：数字token识别")
	fmt.Println("2. 配置文件解析：验证数值配置")
	fmt.Println("3. 表单验证：前后端数据校验")
	fmt.Println("4. 数据清洗：过滤无效数据")
	fmt.Println("5. 命令行解析：验证输入参数")
	fmt.Println()

	fmt.Println("推荐使用：有限状态机算法（方法一），最标准最清晰")
}
