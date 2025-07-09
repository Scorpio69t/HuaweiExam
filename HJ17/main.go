package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 解析指令并返回移动后的坐标
func processInstructions(input string) (int, int) {
	x, y := 0, 0 // 初始位置 (0,0)

	// 按分号分割指令
	instructions := strings.Split(input, ";")

	for _, instruction := range instructions {
		// 跳过空指令
		if len(instruction) == 0 {
			continue
		}

		// 验证指令格式
		if !isValidInstruction(instruction) {
			continue
		}

		// 解析指令
		direction := instruction[0]
		distanceStr := instruction[1:] // 去掉方向字符
		distance, _ := strconv.Atoi(distanceStr)

		// 根据方向更新坐标
		switch direction {
		case 'A': // 向左
			x -= distance
		case 'D': // 向右
			x += distance
		case 'W': // 向上
			y += distance
		case 'S': // 向下
			y -= distance
		}
	}

	return x, y
}

// 验证指令是否合法
func isValidInstruction(instruction string) bool {
	// 指令长度至少为2（方向+数字）
	if len(instruction) < 2 {
		return false
	}

	// 检查方向字符
	direction := instruction[0]
	if direction != 'A' && direction != 'D' && direction != 'W' && direction != 'S' {
		return false
	}

	// 提取数字部分
	distanceStr := instruction[1:]

	// 检查数字部分是否全为数字
	for _, char := range distanceStr {
		if char < '0' || char > '9' {
			return false
		}
	}

	// 解析数字并检查范围
	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return false
	}

	// 数字必须在1-99范围内
	if distance <= 0 || distance >= 100 {
		return false
	}

	return true
}

// 测试函数
func testCases() {
	fmt.Println("=== 测试用例 ===")

	testCases := []string{
		"A10;S20;W10;D30;X;A1A;B10A11;;A10;",
		"ABC;AKL;DA1;D001;W023;A100;S00;",
		"A00;S01;W2;",
	}

	expectedResults := []string{
		"10,-10",
		"0,0",
		"0,1",
	}

	for i, testCase := range testCases {
		x, y := processInstructions(testCase)
		result := fmt.Sprintf("%d,%d", x, y)
		expected := expectedResults[i]

		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("输入: %s\n", testCase)
		fmt.Printf("输出: %s\n", result)
		fmt.Printf("期望: %s\n", expected)
		if result == expected {
			fmt.Println("✓ 通过")
		} else {
			fmt.Println("✗ 失败")
		}
		fmt.Println()
	}
}

func main() {
	// 运行测试用例
	testCases()

	// 交互式输入
	fmt.Println("=== 交互式测试 ===")
	fmt.Println("请输入指令字符串（按Ctrl+C退出）:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			continue
		}

		x, y := processInstructions(input)
		fmt.Printf("最终坐标: %d,%d\n", x, y)
		fmt.Println("请输入下一个指令字符串:")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入时出错: %v\n", err)
	}
}
