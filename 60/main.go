package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 方法一：康托展开算法（最优解法）
func getPermutation1(n int, k int) string {
	// 1. 预计算阶乘
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	// 2. 初始化候选数字
	candidates := make([]int, n)
	for i := 0; i < n; i++ {
		candidates[i] = i + 1
	}

	// 3. 转换为0-based索引
	k--

	// 4. 构建结果
	var result strings.Builder

	// 5. 从高位到低位确定每一位数字
	for i := n; i > 0; i-- {
		// 计算当前位应该选择的索引
		index := k / factorial[i-1]

		// 选择对应的候选数字
		result.WriteString(strconv.Itoa(candidates[index]))

		// 从候选列表中移除已选数字
		candidates = append(candidates[:index], candidates[index+1:]...)

		// 更新k值
		k %= factorial[i-1]
	}

	return result.String()
}

// 方法二：递归实现的康托展开
func getPermutation2(n int, k int) string {
	candidates := make([]int, n)
	for i := 0; i < n; i++ {
		candidates[i] = i + 1
	}

	return buildPermutation(candidates, k-1)
}

// 递归构建排列
func buildPermutation(candidates []int, k int) string {
	n := len(candidates)
	if n == 1 {
		return strconv.Itoa(candidates[0])
	}

	// 计算(n-1)!
	factorial := 1
	for i := 1; i < n; i++ {
		factorial *= i
	}

	// 确定当前位
	index := k / factorial
	digit := candidates[index]

	// 移除已选数字
	newCandidates := append([]int{}, candidates[:index]...)
	newCandidates = append(newCandidates, candidates[index+1:]...)

	// 递归处理剩余位
	return strconv.Itoa(digit) + buildPermutation(newCandidates, k%factorial)
}

// 方法三：字典序迭代法
func getPermutation3(n int, k int) string {
	// 初始化第一个排列
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// 迭代k-1次
	for i := 1; i < k; i++ {
		nextPermutation(nums)
	}

	// 转换为字符串
	var result strings.Builder
	for _, num := range nums {
		result.WriteString(strconv.Itoa(num))
	}
	return result.String()
}

// 生成下一个字典序排列
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// 找到第一个递减的位置
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 找到比nums[i]大的最小数
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 反转i+1之后的部分
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 方法四：回溯+剪枝算法
func getPermutation4(n int, k int) string {
	candidates := make([]int, n)
	for i := 0; i < n; i++ {
		candidates[i] = i + 1
	}

	used := make([]bool, n)
	var path []int
	count := 0
	result := ""

	var backtrack func()
	backtrack = func() {
		if len(path) == n {
			count++
			if count == k {
				// 找到第k个排列
				var sb strings.Builder
				for _, num := range path {
					sb.WriteString(strconv.Itoa(num))
				}
				result = sb.String()
			}
			return
		}

		// 剪枝：如果已经找到结果，直接返回
		if result != "" {
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				// 剪枝：如果当前分支的所有排列数都不够到k，跳过
				factorial := 1
				for j := 1; j <= n-len(path)-1; j++ {
					factorial *= j
				}
				if count+factorial < k {
					count += factorial
					continue
				}

				used[i] = true
				path = append(path, candidates[i])
				backtrack()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}

	backtrack()
	return result
}

// 辅助函数：计算阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 测试用例
func createTestCases() []struct {
	n        int
	k        int
	expected string
	name     string
} {
	return []struct {
		n        int
		k        int
		expected string
		name     string
	}{
		{3, 3, "213", "示例1: n=3, k=3"},
		{4, 9, "2314", "示例2: n=4, k=9"},
		{3, 1, "123", "示例3: n=3, k=1"},
		{1, 1, "1", "边界1: n=1"},
		{3, 6, "321", "边界2: 最大排列"},
		{4, 1, "1234", "边界3: 第一个排列"},
		{4, 24, "4321", "边界4: 最后一个排列"},
		{5, 120, "54321", "边界5: n=5最大排列"},
		{2, 1, "12", "边界6: n=2, k=1"},
		{2, 2, "21", "边界7: n=2, k=2"},
	}
}

// 性能测试
func benchmarkAlgorithm(algorithm func(int, int) string, n int, k int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(n, k)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)
	fmt.Printf("%s (n=%d, k=%d): 平均执行时间 %d 纳秒\n", name, n, k, avgTime)
}

func main() {
	fmt.Println("=== 60. 排列序列 ===")
	fmt.Println()

	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(int, int) string
	}{
		{"康托展开算法", getPermutation1},
		{"递归实现算法", getPermutation2},
		{"字典序迭代算法", getPermutation3},
		{"回溯剪枝算法", getPermutation4},
	}

	// 正确性测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)
		fmt.Printf("  输入: n=%d, k=%d\n", testCase.n, testCase.k)

		results := make([]string, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.n, testCase.k)
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
			fmt.Printf("  ✅ 所有算法结果一致且正确: %s\n", results[0])
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			fmt.Printf("  预期: %s\n", testCase.expected)
			for i, algo := range algorithms {
				fmt.Printf("  %s: %s\n", algo.name, results[i])
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	perfTests := []struct {
		n int
		k int
	}{
		{5, 60},
		{6, 360},
		{7, 2520},
	}

	for _, test := range perfTests {
		fmt.Printf("测试规模 n=%d, k=%d:\n", test.n, test.k)
		for _, algo := range algorithms[:3] { // 只测试前3个算法，回溯法较慢
			benchmarkAlgorithm(algo.fn, test.n, test.k, algo.name)
		}
		fmt.Println()
	}

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("排列序列问题的特点:")
	fmt.Println("1. 康托展开：数学方法直接计算")
	fmt.Println("2. 分组思想：每个数字开头形成一组")
	fmt.Println("3. 阶乘定位：利用阶乘快速定位")
	fmt.Println("4. 贪心选择：每位选择最优数字")
	fmt.Println()

	fmt.Println("=== 康托展开示例 ===")
	fmt.Println("以n=4, k=9为例：")
	fmt.Println("第1组(以1开头): 1-6个排列  (3! = 6)")
	fmt.Println("第2组(以2开头): 7-12个排列 (3! = 6)")
	fmt.Println("第3组(以3开头): 13-18个排列")
	fmt.Println("第4组(以4开头): 19-24个排列")
	fmt.Println()
	fmt.Println("k=9在第2组，选择2作为第一位")
	fmt.Println("剩余[1,3,4]，k'=3，选择3作为第二位")
	fmt.Println("剩余[1,4]，k'=1，选择4作为第三位")
	fmt.Println("剩余[1]，选择1作为第四位")
	fmt.Println("结果: 2314")
	fmt.Println()

	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 康托展开: O(n²)，每次删除候选数字O(n)")
	fmt.Println("- 递归实现: O(n²)，递归n层，每层O(n)")
	fmt.Println("- 字典序迭代: O(k×n)，迭代k次，每次O(n)")
	fmt.Println("- 回溯剪枝: O(k×n)，生成k个排列")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 康托展开: O(n)，候选列表+阶乘数组")
	fmt.Println("- 递归实现: O(n)，递归栈深度")
	fmt.Println("- 字典序迭代: O(n)，存储当前排列")
	fmt.Println("- 回溯剪枝: O(n)，递归栈+路径")
	fmt.Println()

	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 数学建模：康托展开转化为数学计算")
	fmt.Println("2. 预计算：提前计算阶乘表")
	fmt.Println("3. 贪心思想：每位选择当前最优")
	fmt.Println("4. 索引转换：注意1-based和0-based")
	fmt.Println("5. 候选管理：动态维护剩余数字")
	fmt.Println("6. 分治思想：确定一位后递归处理")
	fmt.Println()

	fmt.Println("=== 数学背景 ===")
	fmt.Println("康托展开（Cantor Expansion）:")
	fmt.Println("- 全排列到自然数的双射")
	fmt.Println("- 阶乘数系统（阶乘进制）")
	fmt.Println("- 第i位基数是i!，取值范围[0,i]")
	fmt.Println()
	fmt.Println("正向康托展开（排列→数字）:")
	fmt.Println("X = a[n]×(n-1)! + a[n-1]×(n-2)! + ... + a[1]×0!")
	fmt.Println()
	fmt.Println("逆向康托展开（数字→排列）:")
	fmt.Println("通过除法和取模逐位确定数字")
	fmt.Println()

	fmt.Println("推荐使用：康托展开算法（方法一），时间效率最高")
}
