package main

import (
	"fmt"
)

const (
	INT_MAX = 2147483647
	INT_MIN = -2147483648
)

// divide 使用快倍增（重复加倍减法）实现除法，不使用乘/除/取余
func divide(dividend int, divisor int) int {
	// 边界：除数为0（按题意不会发生），但保险起见
	if divisor == 0 {
		return INT_MAX
	}
	// 溢出边界：INT_MIN / -1
	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}

	// 统一转为 int64 处理，避免中间溢出
	d := int64(dividend)
	dv := int64(divisor)

	// 记录符号
	neg := (d < 0) != (dv < 0)

	// 取绝对值（在 int64 范围内安全）
	ad := abs64(d)
	adv := abs64(dv)

	// 快倍增：每次找到不超过 ad 的 adv 的最大 2^k 倍
	var res int64 = 0
	for ad >= adv {
		tmp := adv
		mul := int64(1)
		for ad >= (tmp << 1) {
			tmp <<= 1
			mul <<= 1
		}
		ad -= tmp
		res += mul
	}

	if neg {
		res = -res
	}
	// 截断到 32 位范围
	if res > int64(INT_MAX) {
		return INT_MAX
	}
	if res < int64(INT_MIN) {
		return INT_MIN
	}
	return int(res)
}

// divideBit 位移长除法：从高位到低位构造结果
func divideBit(dividend int, divisor int) int {
	if divisor == 0 {
		return INT_MAX
	}
	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}

	d := int64(dividend)
	dv := int64(divisor)
	neg := (d < 0) != (dv < 0)
	ad := abs64(d)
	adv := abs64(dv)

	var res int64 = 0
	// 从 31 到 0 尝试（32位）
	for i := 31; i >= 0; i-- {
		if (ad >> i) >= adv {
			res += 1 << uint(i)
			ad -= adv << uint(i)
		}
	}
	if neg {
		res = -res
	}
	if res > int64(INT_MAX) {
		return INT_MAX
	}
	if res < int64(INT_MIN) {
		return INT_MIN
	}
	return int(res)
}

// divideLinear 线性减法（仅作教学/小数值验证，不建议大数据）
func divideLinear(dividend int, divisor int) int {
	if divisor == 0 {
		return INT_MAX
	}
	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}
	if dividend == 0 {
		return 0
	}

	d := int64(dividend)
	dv := int64(divisor)
	neg := (d < 0) != (dv < 0)
	ad := abs64(d)
	adv := abs64(dv)

	var res int64 = 0
	for ad >= adv {
		ad -= adv
		res++
	}
	if neg {
		res = -res
	}
	if res > int64(INT_MAX) {
		return INT_MAX
	}
	if res < int64(INT_MIN) {
		return INT_MIN
	}
	return int(res)
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("两数相除（不使用乘/除/取余）测试")
	fmt.Println("=========================")

	tests := []struct {
		dividend int
		divisor  int
		expect   int
		name     string
	}{
		{10, 3, 3, "10/3"},
		{7, -3, -2, "7/-3"},
		{INT_MIN, -1, INT_MAX, "INT_MIN/-1 溢出"},
		{INT_MIN, 1, INT_MIN, "INT_MIN/1"},
		{0, 1, 0, "0/1"},
		{1, 1, 1, "1/1"},
		{-1010369383, -2147483648, 0, "小于1的绝对值商"},
		{INT_MAX, 2, 1073741823, "INT_MAX/2"},
		{INT_MIN, -3, 715827882, "INT_MIN/-3"},
		{-2147483647, 2, -1073741823, "-2147483647/2"},
	}

	for _, tc := range tests {
		fmt.Printf("\n用例: %s  dividend=%d divisor=%d\n", tc.name, tc.dividend, tc.divisor)
		ans1 := divide(tc.dividend, tc.divisor)
		ans2 := divideBit(tc.dividend, tc.divisor)
		// 线性法只在小值时验证，避免长时间
		var ans3 int
		if abs64(int64(tc.dividend)) < 1000 && abs64(int64(tc.divisor)) > 0 {
			ans3 = divideLinear(tc.dividend, tc.divisor)
			fmt.Printf("divideLinear: %d\n", ans3)
		}
		fmt.Printf("divide     : %d\n", ans1)
		fmt.Printf("divideBit  : %d\n", ans2)
		fmt.Printf("期望       : %d\n", tc.expect)
	}
}
