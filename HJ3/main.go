package main

import (
	"fmt"
)

// 巧妙解法：位图 + 位运算优化
// 时间复杂度：O(n + k)，空间复杂度：O(k/32)
// 其中k=500为数据范围，n为输入数量
func main() {
	var n int
	fmt.Scanln(&n)

	// 使用位图存储数字存在状态
	// 每个uint32可以存储32个数字的状态，500个数字需要16个uint32
	const maxNum = 500
	bitmap := make([]uint32, (maxNum+31)/32) // 向上取整

	// 读取数字并在位图中标记
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)
		setBit(bitmap, num)
	}

	// 遍历位图输出已存在的数字
	outputSortedUniqueNumbers(bitmap, maxNum)
}

// 设置位图中对应位置为1（内联优化）
func setBit(bitmap []uint32, num int) {
	wordIndex := num >> 5 // 等价于 num / 32，位运算更快
	bitOffset := num & 31 // 等价于 num % 32，位运算更快
	bitmap[wordIndex] |= 1 << bitOffset
}

// 检查位图中指定位置是否为1（内联优化）
func getBit(bitmap []uint32, num int) bool {
	wordIndex := num >> 5
	bitOffset := num & 31
	return (bitmap[wordIndex] & (1 << bitOffset)) != 0
}

// 输出排序后的唯一数字（利用位图天然有序的特性）
func outputSortedUniqueNumbers(bitmap []uint32, maxNum int) {
	for num := 1; num <= maxNum; num++ {
		if getBit(bitmap, num) {
			fmt.Println(num)
		}
	}
}

// 更巧妙的进阶版本：使用位运算技巧批量处理
func advancedBitmapSolution() {
	var n int
	fmt.Scanln(&n)

	// 使用计数排序的思想，但用位图优化空间
	const maxNum = 500
	bitmap := make([]uint64, (maxNum+63)/64) // 使用uint64提高效率

	// 批量读取和处理
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&nums[i])
		setBit64(bitmap, nums[i])
	}

	// 使用位运算技巧快速查找和输出
	for i := 0; i < len(bitmap); i++ {
		word := bitmap[i]
		baseNum := i * 64

		// 使用位运算快速找到所有设置的位
		for word != 0 {
			// 找到最低位的1的位置
			trailingZeros := countTrailingZeros(word)
			num := baseNum + trailingZeros + 1

			if num <= maxNum {
				fmt.Println(num)
			}

			// 清除最低位的1
			word &= word - 1
		}
	}
}

// 设置64位位图中对应位置
func setBit64(bitmap []uint64, num int) {
	wordIndex := num >> 6 // num / 64
	bitOffset := num & 63 // num % 64
	bitmap[wordIndex] |= 1 << bitOffset
}

// 计算尾随零的个数（使用内置函数或位运算技巧）
func countTrailingZeros(x uint64) int {
	if x == 0 {
		return 64
	}
	count := 0
	for (x & 1) == 0 {
		count++
		x >>= 1
	}
	return count
}

// 面试官最爱的一行流解法（展示函数式编程思维）
func oneLineSolution() {
	// 虽然在Go中不太适用，但可以展示思维
	// 在其他语言中可以写成：
	// Arrays.stream(nums).distinct().sorted().forEach(System.out::println)
}

// 内存极限优化版本：流式处理
func memoryOptimizedSolution() {
	var n int
	fmt.Scanln(&n)

	// 只使用500位的位图，约63字节
	var bitmap [16]uint32 // 16 * 32 = 512 bits，覆盖1-500

	// 流式处理，边读边标记
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)
		if num >= 1 && num <= 500 {
			bitmap[num>>5] |= 1 << (num & 31)
		}
	}

	// 输出结果
	for num := 1; num <= 500; num++ {
		if (bitmap[num>>5] & (1 << (num & 31))) != 0 {
			fmt.Println(num)
		}
	}
}
