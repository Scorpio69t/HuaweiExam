package main

import (
	"fmt"
	"strings"
)

// =========================== 方法一：贪心+分段格式化（最优解法） ===========================

// fullJustify 贪心+分段格式化
// 时间复杂度：O(n)，n为单词总数
// 空间复杂度：O(1)，不计结果数组
func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	i := 0

	for i < len(words) {
		// 确定当前行包含的单词范围 [i, j)
		j := i
		lineLen := 0

		for j < len(words) {
			newLen := lineLen + len(words[j])
			if j > i {
				newLen += 1 // 单词间至少一个空格
			}
			if newLen > maxWidth {
				break
			}
			lineLen = newLen
			j++
		}

		// 格式化当前行
		isLastLine := j == len(words)
		line := formatLine(words, i, j, maxWidth, isLastLine)
		result = append(result, line)

		i = j
	}

	return result
}

// formatLine 格式化一行
func formatLine(words []string, start, end, maxWidth int, isLastLine bool) string {
	numWords := end - start

	// 最后一行或只有一个单词，左对齐
	if isLastLine || numWords == 1 {
		return leftAlign(words, start, end, maxWidth)
	}

	// 两端对齐
	return justify(words, start, end, maxWidth)
}

// justify 两端对齐
func justify(words []string, start, end, maxWidth int) string {
	numWords := end - start

	// 计算单词总长度
	totalWordsLen := 0
	for i := start; i < end; i++ {
		totalWordsLen += len(words[i])
	}

	// 计算总空格数和间隙数
	totalSpaces := maxWidth - totalWordsLen
	gaps := numWords - 1

	// 每个间隙的基础空格数和额外空格数
	baseSpaces := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	// 构建结果
	var builder strings.Builder
	for i := start; i < end; i++ {
		builder.WriteString(words[i])

		if i < end-1 {
			// 添加基础空格
			for j := 0; j < baseSpaces; j++ {
				builder.WriteByte(' ')
			}

			// 左侧优先，添加额外空格
			if extraSpaces > 0 {
				builder.WriteByte(' ')
				extraSpaces--
			}
		}
	}

	return builder.String()
}

// leftAlign 左对齐
func leftAlign(words []string, start, end, maxWidth int) string {
	var builder strings.Builder

	for i := start; i < end; i++ {
		builder.WriteString(words[i])

		if i < end-1 {
			builder.WriteByte(' ')
		}
	}

	// 右侧填充空格
	for builder.Len() < maxWidth {
		builder.WriteByte(' ')
	}

	return builder.String()
}

// =========================== 方法二：模拟排版 ===========================

// fullJustify2 模拟排版
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func fullJustify2(words []string, maxWidth int) []string {
	result := make([]string, 0)
	currentLine := make([]string, 0)
	currentLen := 0

	for i, word := range words {
		// 检查是否可以加入当前行
		needLen := currentLen + len(word)
		if len(currentLine) > 0 {
			needLen += 1 // 单词间空格
		}

		if needLen <= maxWidth {
			// 可以加入当前行
			currentLine = append(currentLine, word)
			currentLen = needLen
		} else {
			// 当前行已满，格式化并加入结果
			line := formatLine2(currentLine, maxWidth, false)
			result = append(result, line)

			// 开始新行
			currentLine = []string{word}
			currentLen = len(word)
		}

		// 最后一个单词
		if i == len(words)-1 {
			line := formatLine2(currentLine, maxWidth, true)
			result = append(result, line)
		}
	}

	return result
}

// formatLine2 格式化一行（方法二）
func formatLine2(words []string, maxWidth int, isLastLine bool) string {
	if len(words) == 0 {
		return strings.Repeat(" ", maxWidth)
	}

	// 最后一行或只有一个单词
	if isLastLine || len(words) == 1 {
		line := strings.Join(words, " ")
		return line + strings.Repeat(" ", maxWidth-len(line))
	}

	// 两端对齐
	totalLen := 0
	for _, word := range words {
		totalLen += len(word)
	}

	totalSpaces := maxWidth - totalLen
	gaps := len(words) - 1
	baseSpaces := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	var builder strings.Builder
	for i, word := range words {
		builder.WriteString(word)

		if i < len(words)-1 {
			spaces := baseSpaces
			if i < extraSpaces {
				spaces++
			}
			builder.WriteString(strings.Repeat(" ", spaces))
		}
	}

	return builder.String()
}

// =========================== 方法三：递归分行 ===========================

// fullJustify3 递归分行
// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为行数
func fullJustify3(words []string, maxWidth int) []string {
	return justifyHelper(words, 0, maxWidth)
}

// justifyHelper 递归辅助函数
func justifyHelper(words []string, start int, maxWidth int) []string {
	if start >= len(words) {
		return []string{}
	}

	// 确定当前行包含的单词
	end := start
	lineLen := 0

	for end < len(words) {
		newLen := lineLen + len(words[end])
		if end > start {
			newLen += 1
		}
		if newLen > maxWidth {
			break
		}
		lineLen = newLen
		end++
	}

	// 格式化当前行
	isLastLine := end == len(words)
	line := formatLine(words, start, end, maxWidth, isLastLine)

	// 递归处理剩余单词
	restLines := justifyHelper(words, end, maxWidth)

	return append([]string{line}, restLines...)
}

// =========================== 方法四：预计算空格数组 ===========================

// fullJustify4 预计算空格数组
// 时间复杂度：O(n)
// 空间复杂度：O(k)，k为每行单词数
func fullJustify4(words []string, maxWidth int) []string {
	result := make([]string, 0)
	i := 0

	for i < len(words) {
		// 确定当前行包含的单词
		j := i
		lineLen := 0

		for j < len(words) {
			newLen := lineLen + len(words[j])
			if j > i {
				newLen += 1
			}
			if newLen > maxWidth {
				break
			}
			lineLen = newLen
			j++
		}

		// 构建当前行
		line := buildLine(words, i, j, maxWidth, j == len(words))
		result = append(result, line)

		i = j
	}

	return result
}

// buildLine 使用空格数组构建行
func buildLine(words []string, start, end, maxWidth int, isLastLine bool) string {
	numWords := end - start

	// 最后一行或单个单词
	if isLastLine || numWords == 1 {
		return leftAlign(words, start, end, maxWidth)
	}

	// 计算每个间隙的空格数
	totalLen := 0
	for i := start; i < end; i++ {
		totalLen += len(words[i])
	}

	totalSpaces := maxWidth - totalLen
	gaps := numWords - 1

	// 预计算每个间隙的空格数
	spaces := make([]int, gaps)
	baseSpaces := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	for i := 0; i < gaps; i++ {
		spaces[i] = baseSpaces
		if i < extraSpaces {
			spaces[i]++
		}
	}

	// 构建结果
	var builder strings.Builder
	for i := start; i < end; i++ {
		builder.WriteString(words[i])

		if i < end-1 {
			builder.WriteString(strings.Repeat(" ", spaces[i-start]))
		}
	}

	return builder.String()
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 68: 文本左右对齐 ===\n")

	// 测试用例
	testCases := []struct {
		words    []string
		maxWidth int
		expect   []string
	}{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expect:   []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expect:   []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			expect:   []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
		{
			words:    []string{"a"},
			maxWidth: 1,
			expect:   []string{"a"},
		},
		{
			words:    []string{"a", "b"},
			maxWidth: 3,
			expect:   []string{"a b"},
		},
		{
			words:    []string{"a", "b", "c"},
			maxWidth: 5,
			expect:   []string{"a b c"},
		},
		{
			words:    []string{"Listen", "to", "many,", "speak", "to", "a", "few."},
			maxWidth: 6,
			expect:   []string{"Listen", "to    ", "many, ", "speak ", "to   a", "few.  "},
		},
	}

	fmt.Println("方法一：贪心+分段格式化")
	runTests(testCases, fullJustify)

	fmt.Println("\n方法二：模拟排版")
	runTests(testCases, fullJustify2)

	fmt.Println("\n方法三：递归分行")
	runTests(testCases, fullJustify3)

	fmt.Println("\n方法四：预计算空格数组")
	runTests(testCases, fullJustify4)

	// 详细输出示例
	fmt.Println("\n=== 详细输出示例 ===")
	detailedExample()
}

// runTests 运行测试用例
func runTests(testCases []struct {
	words    []string
	maxWidth int
	expect   []string
}, fn func([]string, int) []string) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.words, tc.maxWidth)
		status := "✅"

		// 比较结果
		if len(result) != len(tc.expect) {
			status = "❌"
		} else {
			for j := range result {
				if result[j] != tc.expect[j] {
					status = "❌"
					break
				}
			}
		}

		if status == "✅" {
			passCount++
		}

		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: %v, maxWidth=%d\n", tc.words, tc.maxWidth)
			fmt.Printf("    输出: %v\n", result)
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

// detailedExample 详细输出示例
func detailedExample() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16

	fmt.Printf("输入: words = %v\n", words)
	fmt.Printf("      maxWidth = %d\n\n", maxWidth)

	result := fullJustify(words, maxWidth)

	fmt.Println("输出:")
	for i, line := range result {
		// 显示行号和内容
		fmt.Printf("  行%d: \"%s\" (长度:%d)\n", i+1, line, len(line))

		// 显示空格分布
		spaceCount := 0
		for _, ch := range line {
			if ch == ' ' {
				spaceCount++
			}
		}
		fmt.Printf("       空格数: %d\n", spaceCount)
	}

	// 分析每行的空格分配
	fmt.Println("\n空格分配分析:")
	fmt.Println("  第1行: \"This    is    an\"")
	fmt.Println("         单词: [This, is, an], 长度: 4+2+2=8")
	fmt.Println("         总空格: 16-8=8, 间隙: 2")
	fmt.Println("         每个间隙: 8/2=4个空格")

	fmt.Println("\n  第2行: \"example  of text\"")
	fmt.Println("         单词: [example, of, text], 长度: 7+2+4=13")
	fmt.Println("         总空格: 16-13=3, 间隙: 2")
	fmt.Println("         基础空格: 3/2=1, 额外: 3%2=1")
	fmt.Println("         第1个间隙: 1+1=2, 第2个间隙: 1")

	fmt.Println("\n  第3行: \"justification.  \"")
	fmt.Println("         单词: [justification.], 长度: 14")
	fmt.Println("         最后一行，左对齐，右侧填充2个空格")
}
