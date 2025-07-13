package main

import (
	"fmt"
	"math"
)

const (
	TARGET = 24.0 // 目标值
	EPS    = 1e-6 // 浮点数精度
)

// 检查两个浮点数是否相等
func isEqual(a, b float64) bool {
	return math.Abs(a-b) < EPS
}

// 检查浮点数是否为零
func isZero(a float64) bool {
	return math.Abs(a) < EPS
}

// 递归回溯函数
// nums: 当前数字列表
// 返回值: 是否能计算出24
func canGet24(nums []float64) bool {
	n := len(nums)

	// 终止条件：只剩一个数字
	if n == 1 {
		return isEqual(nums[0], TARGET)
	}

	// 枚举所有可能的数字对
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := nums[i], nums[j]

			// 创建新的数字列表（移除选中的两个数字）
			newNums := make([]float64, 0, n-1)
			for k := 0; k < n; k++ {
				if k != i && k != j {
					newNums = append(newNums, nums[k])
				}
			}

			// 尝试四种运算符
			// 1. 加法（交换律，只需计算一次）
			newNums = append(newNums, a+b)
			if canGet24(newNums) {
				return true
			}
			newNums = newNums[:len(newNums)-1] // 回溯

			// 2. 减法（不满足交换律，需要计算两次）
			// a - b
			newNums = append(newNums, a-b)
			if canGet24(newNums) {
				return true
			}
			newNums = newNums[:len(newNums)-1] // 回溯

			// b - a
			newNums = append(newNums, b-a)
			if canGet24(newNums) {
				return true
			}
			newNums = newNums[:len(newNums)-1] // 回溯

			// 3. 乘法（交换律，只需计算一次）
			newNums = append(newNums, a*b)
			if canGet24(newNums) {
				return true
			}
			newNums = newNums[:len(newNums)-1] // 回溯

			// 4. 除法（不满足交换律，需要计算两次）
			// a / b
			if !isZero(b) {
				newNums = append(newNums, a/b)
				if canGet24(newNums) {
					return true
				}
				newNums = newNums[:len(newNums)-1] // 回溯
			}

			// b / a
			if !isZero(a) {
				newNums = append(newNums, b/a)
				if canGet24(newNums) {
					return true
				}
				newNums = newNums[:len(newNums)-1] // 回溯
			}
		}
	}

	return false
}

// 24点游戏主函数
func game24(a, b, c, d int) bool {
	// 转换为浮点数列表
	nums := []float64{float64(a), float64(b), float64(c), float64(d)}

	// 调用递归函数
	return canGet24(nums)
}

func main() {
	var a, b, c, d int

	// 读取输入
	fmt.Scan(&a, &b, &c, &d)

	// 判断是否能得到24
	if game24(a, b, c, d) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
