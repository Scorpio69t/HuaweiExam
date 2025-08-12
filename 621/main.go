package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// 方法一：贪心+数学公式（推荐）
// 时间复杂度：O(n)，空间复杂度：O(1)
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}

	// 统计每个任务的频率
	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}

	// 找到最大频率
	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	// 统计具有最大频率的任务数量
	count := 0
	for _, f := range freq {
		if f == maxFreq {
			count++
		}
	}

	// 计算最短时间
	// 公式：(maxFreq-1) * (n+1) + count
	minTime := (maxFreq-1)*(n+1) + count

	// 如果任务总数更多，返回任务总数
	if minTime < len(tasks) {
		return len(tasks)
	}
	return minTime
}

// 方法二：模拟调度（使用优先队列）
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func leastIntervalSimulation(tasks []byte, n int) int {
	// 统计任务频率
	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}

	// 创建最大堆
	pq := &MaxHeap{}
	heap.Init(pq)

	// 将任务加入优先队列
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			heap.Push(pq, &Task{task: byte('A' + i), count: freq[i]})
		}
	}

	time := 0
	// 记录每个任务的冷却时间
	coolDown := make(map[byte]int)

	for pq.Len() > 0 {
		// 找到可以执行的任务
		temp := []*Task{}

		for pq.Len() > 0 {
			task := heap.Pop(pq).(*Task)
			if coolDown[task.task] <= time {
				// 可以执行这个任务
				task.count--
				if task.count > 0 {
					heap.Push(pq, task)
				}
				coolDown[task.task] = time + n + 1
				break
			} else {
				// 还不能执行，暂时保存
				temp = append(temp, task)
			}
		}

		// 将暂时不能执行的任务放回队列
		for _, task := range temp {
			heap.Push(pq, task)
		}

		time++
	}

	return time
}

// 方法三：优化贪心（单次遍历）
// 时间复杂度：O(n)，空间复杂度：O(1)
func leastIntervalOptimized(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}

	freq := make([]int, 26)
	maxFreq := 0
	count := 0

	// 单次遍历完成统计和计算
	for _, task := range tasks {
		freq[task-'A']++
		if freq[task-'A'] > maxFreq {
			maxFreq = freq[task-'A']
			count = 1
		} else if freq[task-'A'] == maxFreq {
			count++
		}
	}

	minTime := (maxFreq-1)*(n+1) + count
	if minTime < len(tasks) {
		return len(tasks)
	}
	return minTime
}

// 方法四：数学公式简化版
// 时间复杂度：O(n)，空间复杂度：O(1)
func leastIntervalMath(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}

	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}

	maxFreq := 0
	count := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
			count = 1
		} else if f == maxFreq {
			count++
		}
	}

	// 使用标准公式
	minTime := (maxFreq-1)*(n+1) + count
	if minTime < len(tasks) {
		return len(tasks)
	}
	return minTime
}

// 优先队列相关结构
type Task struct {
	task  byte
	count int
}

type MaxHeap []*Task

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Task))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func runTests() {
	type testCase struct {
		tasks    []byte
		n        int
		expected int
		desc     string
	}

	tests := []testCase{
		{[]byte("AAABBB"), 2, 8, "示例1"},
		{[]byte("ACADBB"), 1, 6, "示例2"},
		{[]byte("AAAAAABCDEFG"), 2, 16, "示例3"},
		{[]byte("ABC"), 0, 3, "冷却时间为0"},
		{[]byte("AAA"), 1, 5, "单一任务"},
		{[]byte{}, 1, 0, "空数组"},
		{[]byte("AABBCC"), 2, 6, "多任务相同频率"},
		{[]byte("AB"), 100, 2, "长冷却时间"},
		{[]byte("ABCDEF"), 2, 6, "无重复任务"},
		{[]byte("AAAA"), 2, 10, "单一任务长冷却"},
	}

	fmt.Println("=== 621. 任务调度器 - 测试 ===")
	for i, tc := range tests {
		r1 := leastInterval(tc.tasks, tc.n)
		r2 := leastIntervalSimulation(tc.tasks, tc.n)
		r3 := leastIntervalOptimized(tc.tasks, tc.n)
		r4 := leastIntervalMath(tc.tasks, tc.n)

		ok := (r1 == tc.expected) && (r2 == tc.expected) && (r3 == tc.expected) && (r4 == tc.expected)
		status := "✅"
		if !ok {
			status = "❌"
		}

		fmt.Printf("用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: tasks=%v, n=%d\n", string(tc.tasks), tc.n)
		fmt.Printf("期望: %d\n", tc.expected)
		fmt.Printf("贪心+数学: %d, 模拟调度: %d, 优化贪心: %d, 数学简化: %d\n", r1, r2, r3, r4)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 50))
	}
}

func main() {
	runTests()
}
