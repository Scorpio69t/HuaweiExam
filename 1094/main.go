package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 方法一：差分数组（推荐）
// 时间复杂度：O(n + max(to))，空间复杂度：O(max(to))
func carPooling1(trips [][]int, capacity int) bool {
	// 创建差分数组，最大位置为1000
	diff := make([]int, 1001)

	// 记录每个位置的人数变化
	for _, trip := range trips {
		passengers, from, to := trip[0], trip[1], trip[2]
		diff[from] += passengers
		diff[to] -= passengers
	}

	// 计算前缀和，检查是否超过容量
	current := 0
	for i := 0; i < 1001; i++ {
		current += diff[i]
		if current > capacity {
			return false
		}
	}

	return true
}

// 方法二：排序 + 模拟
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func carPooling2(trips [][]int, capacity int) bool {
	// 创建事件数组
	events := make([][]int, 0, len(trips)*2)

	// 添加上车和下车事件
	for _, trip := range trips {
		passengers, from, to := trip[0], trip[1], trip[2]
		// 上车事件：[位置, 乘客数, 1表示上车]
		events = append(events, []int{from, passengers, 1})
		// 下车事件：[位置, 乘客数, -1表示下车]
		events = append(events, []int{to, passengers, -1})
	}

	// 按位置排序，相同位置下车优先于上车
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		// 相同位置，下车优先
		return events[i][2] < events[j][2]
	})

	// 模拟车辆行驶过程
	current := 0
	for _, event := range events {
		passengers, action := event[1], event[2]
		if action == 1 {
			// 上车
			current += passengers
			if current > capacity {
				return false
			}
		} else {
			// 下车
			current -= passengers
		}
	}

	return true
}

// 方法三：优先队列
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func carPooling3(trips [][]int, capacity int) bool {
	// 按上车位置排序
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

	// 使用优先队列记录当前在车上的乘客
	pq := &PriorityQueue{}
	heap.Init(pq)

	current := 0
	for _, trip := range trips {
		passengers, from, to := trip[0], trip[1], trip[2]

		// 移除已下车的乘客
		for pq.Len() > 0 && (*pq)[0].to <= from {
			item := heap.Pop(pq).(*Passenger)
			current -= item.passengers
		}

		// 添加新乘客
		current += passengers
		if current > capacity {
			return false
		}

		// 将乘客加入优先队列
		heap.Push(pq, &Passenger{passengers: passengers, to: to})
	}

	return true
}

// 乘客结构体
type Passenger struct {
	passengers int
	to         int
}

// 优先队列实现
type PriorityQueue []*Passenger

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].to < pq[j].to
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Passenger)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// 方法四：优化的差分数组
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func carPooling4(trips [][]int, capacity int) bool {
	// 记录所有位置
	positions := make(map[int]bool)
	for _, trip := range trips {
		positions[trip[1]] = true
		positions[trip[2]] = true
	}

	// 将位置排序
	sortedPositions := make([]int, 0, len(positions))
	for pos := range positions {
		sortedPositions = append(sortedPositions, pos)
	}
	sort.Ints(sortedPositions)

	// 创建差分数组
	diff := make(map[int]int)
	for _, trip := range trips {
		passengers, from, to := trip[0], trip[1], trip[2]
		diff[from] += passengers
		diff[to] -= passengers
	}

	// 计算前缀和
	current := 0
	for _, pos := range sortedPositions {
		current += diff[pos]
		if current > capacity {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("=== 1094. 拼车 ===")

	// 测试用例1
	trips1 := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity1 := 4
	fmt.Printf("测试用例1: trips=%v, capacity=%d\n", trips1, capacity1)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips1, capacity1))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips1, capacity1))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips1, capacity1))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips1, capacity1))
	fmt.Println()

	// 测试用例2
	trips2 := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity2 := 5
	fmt.Printf("测试用例2: trips=%v, capacity=%d\n", trips2, capacity2)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips2, capacity2))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips2, capacity2))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips2, capacity2))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips2, capacity2))
	fmt.Println()

	// 测试用例3
	trips3 := [][]int{{2, 1, 5}, {3, 3, 7}, {1, 2, 4}}
	capacity3 := 6
	fmt.Printf("测试用例3: trips=%v, capacity=%d\n", trips3, capacity3)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips3, capacity3))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips3, capacity3))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips3, capacity3))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips3, capacity3))
	fmt.Println()

	// 额外测试用例
	trips4 := [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}
	capacity4 := 11
	fmt.Printf("额外测试: trips=%v, capacity=%d\n", trips4, capacity4)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips4, capacity4))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips4, capacity4))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips4, capacity4))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips4, capacity4))
	fmt.Println()

	// 边界测试用例
	trips5 := [][]int{{1, 0, 1}}
	capacity5 := 1
	fmt.Printf("边界测试: trips=%v, capacity=%d\n", trips5, capacity5)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips5, capacity5))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips5, capacity5))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips5, capacity5))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips5, capacity5))
	fmt.Println()

	// 复杂测试用例
	trips6 := [][]int{{9, 0, 1}, {3, 3, 7}, {4, 1, 5}, {2, 2, 6}}
	capacity6 := 10
	fmt.Printf("复杂测试: trips=%v, capacity=%d\n", trips6, capacity6)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips6, capacity6))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips6, capacity6))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips6, capacity6))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips6, capacity6))
	fmt.Println()

	// 重叠测试用例
	trips7 := [][]int{{2, 1, 5}, {3, 1, 7}, {1, 2, 4}}
	capacity7 := 5
	fmt.Printf("重叠测试: trips=%v, capacity=%d\n", trips7, capacity7)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips7, capacity7))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips7, capacity7))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips7, capacity7))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips7, capacity7))
	fmt.Println()

	// 容量不足测试用例
	trips8 := [][]int{{5, 0, 3}, {3, 2, 5}, {2, 1, 4}}
	capacity8 := 8
	fmt.Printf("容量不足测试: trips=%v, capacity=%d\n", trips8, capacity8)
	fmt.Printf("方法一结果: %t\n", carPooling1(trips8, capacity8))
	fmt.Printf("方法二结果: %t\n", carPooling2(trips8, capacity8))
	fmt.Printf("方法三结果: %t\n", carPooling3(trips8, capacity8))
	fmt.Printf("方法四结果: %t\n", carPooling4(trips8, capacity8))
}
