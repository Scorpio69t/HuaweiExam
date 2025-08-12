package main

import (
	"fmt"
	"strings"
	"time"
)

// 解法一：BFS路线图（推荐解法）
// 时间复杂度：O(N²+S)，空间复杂度：O(N²+S)
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	if len(routes) == 0 {
		return -1
	}

	// 构建车站到路线的映射
	stationToRoutes := make(map[int][]int)
	for i, route := range routes {
		for _, station := range route {
			stationToRoutes[station] = append(stationToRoutes[station], i)
		}
	}

	// 检查起点和终点是否存在于路线中
	sourceRoutes, sourceExists := stationToRoutes[source]
	targetRoutes, targetExists := stationToRoutes[target]

	if !sourceExists || !targetExists {
		return -1
	}

	// 检查是否可以直达
	sourceSet := make(map[int]bool)
	for _, route := range sourceRoutes {
		sourceSet[route] = true
	}
	for _, route := range targetRoutes {
		if sourceSet[route] {
			return 1 // 同一条路线，只需一辆公交车
		}
	}

	// BFS搜索
	queue := make([]int, 0)
	visited := make([]bool, len(routes))

	// 将包含起点的所有路线加入队列
	for _, routeIdx := range sourceRoutes {
		queue = append(queue, routeIdx)
		visited[routeIdx] = true
	}

	steps := 1

	for len(queue) > 0 {
		size := len(queue)

		// 处理当前层的所有路线
		for i := 0; i < size; i++ {
			currentRoute := queue[i]

			// 遍历当前路线的所有车站
			for _, station := range routes[currentRoute] {
				// 获取经过该车站的所有路线
				for _, nextRoute := range stationToRoutes[station] {
					if visited[nextRoute] {
						continue
					}

					// 检查是否到达目标
					if contains(routes[nextRoute], target) {
						return steps + 1
					}

					// 标记并加入队列
					visited[nextRoute] = true
					queue = append(queue, nextRoute)
				}
			}
		}

		queue = queue[size:] // 移除已处理的元素
		steps++
	}

	return -1
}

// 解法二：BFS车站图
// 时间复杂度：O(S²)，空间复杂度：O(S²)
func numBusesToDestinationStations(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// 构建车站连接图
	stationGraph := make(map[int]map[int]bool)

	// 同一路线的车站相互连通
	for _, route := range routes {
		for i := 0; i < len(route); i++ {
			if stationGraph[route[i]] == nil {
				stationGraph[route[i]] = make(map[int]bool)
			}
			for j := 0; j < len(route); j++ {
				if i != j {
					stationGraph[route[i]][route[j]] = true
				}
			}
		}
	}

	// 检查起点是否存在
	if stationGraph[source] == nil {
		return -1
	}

	// BFS搜索最短路径
	queue := []int{source}
	visited := make(map[int]bool)
	visited[source] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		steps++

		for i := 0; i < size; i++ {
			currentStation := queue[i]

			// 遍历所有相邻车站
			for nextStation := range stationGraph[currentStation] {
				if nextStation == target {
					return steps
				}

				if !visited[nextStation] {
					visited[nextStation] = true
					queue = append(queue, nextStation)
				}
			}
		}

		queue = queue[size:]
	}

	return -1
}

// 解法三：双向BFS（优化版本）
// 时间复杂度：O(N²+S)，空间复杂度：O(N²+S)
func numBusesToDestinationBidirectional(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// 构建车站到路线的映射
	stationToRoutes := make(map[int][]int)
	for i, route := range routes {
		for _, station := range route {
			stationToRoutes[station] = append(stationToRoutes[station], i)
		}
	}

	sourceRoutes, sourceExists := stationToRoutes[source]
	targetRoutes, targetExists := stationToRoutes[target]

	if !sourceExists || !targetExists {
		return -1
	}

	// 初始化双向搜索
	forwardQueue := make(map[int]bool)
	backwardQueue := make(map[int]bool)
	forwardVisited := make(map[int]int)
	backwardVisited := make(map[int]int)

	// 起点路线
	for _, route := range sourceRoutes {
		forwardQueue[route] = true
		forwardVisited[route] = 1
	}

	// 终点路线
	for _, route := range targetRoutes {
		backwardQueue[route] = true
		backwardVisited[route] = 1

		// 检查是否可以直达
		if forwardVisited[route] > 0 {
			return 1
		}
	}

	// 双向BFS
	for len(forwardQueue) > 0 && len(backwardQueue) > 0 {
		// 选择较小的队列进行扩展
		if len(forwardQueue) > len(backwardQueue) {
			if result := expandQueue(backwardQueue, backwardVisited, forwardVisited, routes, stationToRoutes); result != -1 {
				return result
			}
		} else {
			if result := expandQueue(forwardQueue, forwardVisited, backwardVisited, routes, stationToRoutes); result != -1 {
				return result
			}
		}
	}

	return -1
}

// 扩展队列的辅助函数
func expandQueue(queue map[int]bool, visited map[int]int, otherVisited map[int]int,
	routes [][]int, stationToRoutes map[int][]int) int {
	nextQueue := make(map[int]bool)

	for routeIdx := range queue {
		currentSteps := visited[routeIdx]

		for _, station := range routes[routeIdx] {
			for _, nextRoute := range stationToRoutes[station] {
				if visited[nextRoute] > 0 {
					continue
				}

				if otherVisited[nextRoute] > 0 {
					return currentSteps + otherVisited[nextRoute]
				}

				nextQueue[nextRoute] = true
				visited[nextRoute] = currentSteps + 1
			}
		}
	}

	// 清空当前队列，用新队列替代
	for k := range queue {
		delete(queue, k)
	}
	for k := range nextQueue {
		queue[k] = true
	}

	return -1
}

// 解法四：A*搜索（启发式优化）
// 时间复杂度：O(N²+S)，空间复杂度：O(N²+S)
func numBusesToDestinationAStar(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// 构建车站到路线的映射
	stationToRoutes := make(map[int][]int)
	for i, route := range routes {
		for _, station := range route {
			stationToRoutes[station] = append(stationToRoutes[station], i)
		}
	}

	sourceRoutes, sourceExists := stationToRoutes[source]
	targetRoutes, targetExists := stationToRoutes[target]

	if !sourceExists || !targetExists {
		return -1
	}

	// 构建目标路线集合
	targetSet := make(map[int]bool)
	for _, route := range targetRoutes {
		targetSet[route] = true
	}

	// A*搜索使用优先队列
	type Node struct {
		routeIdx int
		steps    int
		priority int // f(n) = g(n) + h(n)
	}

	// 简单优先队列实现
	pq := []Node{}
	visited := make(map[int]bool)

	// 启发式函数：如果路线包含目标站点，启发值为0，否则为1
	heuristic := func(routeIdx int) int {
		if targetSet[routeIdx] {
			return 0
		}
		return 1
	}

	// 初始化起点路线
	for _, routeIdx := range sourceRoutes {
		if targetSet[routeIdx] {
			return 1
		}
		h := heuristic(routeIdx)
		pq = append(pq, Node{routeIdx, 1, 1 + h})
	}

	for len(pq) > 0 {
		// 简单的优先队列取最小值
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i].priority < pq[minIdx].priority {
				minIdx = i
			}
		}

		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)

		if visited[current.routeIdx] {
			continue
		}
		visited[current.routeIdx] = true

		// 扩展当前路线
		for _, station := range routes[current.routeIdx] {
			for _, nextRoute := range stationToRoutes[station] {
				if visited[nextRoute] {
					continue
				}

				if targetSet[nextRoute] {
					return current.steps + 1
				}

				h := heuristic(nextRoute)
				newNode := Node{nextRoute, current.steps + 1, current.steps + 1 + h}
				pq = append(pq, newNode)
			}
		}
	}

	return -1
}

// 辅助函数：检查数组是否包含元素
func contains(arr []int, target int) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}

// 辅助函数：构建路线连通图
func buildRouteGraph(routes [][]int) [][]bool {
	n := len(routes)
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}

	// 检查路线间是否有共同车站
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if hasCommonStation(routes[i], routes[j]) {
				graph[i][j] = true
				graph[j][i] = true
			}
		}
	}

	return graph
}

// 检查两条路线是否有共同车站
func hasCommonStation(route1, route2 []int) bool {
	set := make(map[int]bool)
	for _, station := range route1 {
		set[station] = true
	}
	for _, station := range route2 {
		if set[station] {
			return true
		}
	}
	return false
}

// 公交系统模拟器
type BusSystem struct {
	routes          [][]int
	stationToRoutes map[int][]int
	routeGraph      [][]bool
	totalStations   int
	totalRoutes     int
}

// 创建公交系统
func newBusSystem(routes [][]int) *BusSystem {
	bs := &BusSystem{
		routes:          routes,
		stationToRoutes: make(map[int][]int),
		totalRoutes:     len(routes),
	}

	// 构建车站到路线映射
	stationSet := make(map[int]bool)
	for i, route := range routes {
		for _, station := range route {
			bs.stationToRoutes[station] = append(bs.stationToRoutes[station], i)
			stationSet[station] = true
		}
	}
	bs.totalStations = len(stationSet)

	// 构建路线连通图
	bs.routeGraph = buildRouteGraph(routes)

	return bs
}

// 查找最短路径
func (bs *BusSystem) findShortestPath(source, target int) int {
	return numBusesToDestination(bs.routes, source, target)
}

// 获取系统统计信息
func (bs *BusSystem) getStats() map[string]interface{} {
	return map[string]interface{}{
		"total_routes":   bs.totalRoutes,
		"total_stations": bs.totalStations,
		"avg_route_len":  bs.getAverageRouteLength(),
		"connectivity":   bs.getConnectivity(),
	}
}

// 计算平均路线长度
func (bs *BusSystem) getAverageRouteLength() float64 {
	total := 0
	for _, route := range bs.routes {
		total += len(route)
	}
	return float64(total) / float64(bs.totalRoutes)
}

// 计算连通性
func (bs *BusSystem) getConnectivity() float64 {
	connections := 0
	for i := 0; i < bs.totalRoutes; i++ {
		for j := i + 1; j < bs.totalRoutes; j++ {
			if bs.routeGraph[i][j] {
				connections++
			}
		}
	}
	totalPairs := bs.totalRoutes * (bs.totalRoutes - 1) / 2
	if totalPairs == 0 {
		return 0
	}
	return float64(connections) / float64(totalPairs)
}

// 测试函数
func testBusRoutes() {
	testCases := []struct {
		routes   [][]int
		source   int
		target   int
		expected int
		desc     string
	}{
		{
			[][]int{{1, 2, 7}, {3, 6, 7}},
			1, 6, 2,
			"示例1：需要换乘一次",
		},
		{
			[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}},
			15, 12, -1,
			"示例2：无法到达",
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			1, 6, -1,
			"两条独立路线：无法到达",
		},
		{
			[][]int{{1, 2, 3, 4, 5}},
			1, 5, 1,
			"单条路线：直达",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}},
			1, 4, 3,
			"链式连接：需要多次换乘",
		},
		{
			[][]int{{1, 2, 3}, {2, 4, 5}, {3, 5, 6}},
			1, 6, 2,
			"网状结构：多条路径",
		},
		{
			[][]int{{1}, {1}},
			1, 1, 0,
			"起终点相同",
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {1, 6}},
			2, 5, 3,
			"桥接路线：需要三次换乘",
		},
		{
			[][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			0, 9, 1,
			"长路线测试",
		},
		{
			[][]int{{1, 2}, {1, 3}, {1, 4}},
			2, 4, 2,
			"星型结构：中心换乘",
		},
	}

	fmt.Println("=== 公交路线测试 ===")
	fmt.Println()

	for i, tc := range testCases {
		// 测试不同算法
		result1 := numBusesToDestination(tc.routes, tc.source, tc.target)
		result2 := numBusesToDestinationBidirectional(tc.routes, tc.source, tc.target)
		result3 := numBusesToDestinationAStar(tc.routes, tc.source, tc.target)

		status := "✅"
		if result1 != tc.expected {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Printf("路线: %v\n", tc.routes)
		fmt.Printf("起点: %d, 终点: %d\n", tc.source, tc.target)
		fmt.Printf("期望: %d, 实际: %d\n", tc.expected, result1)

		// 验证算法一致性
		consistent := result1 == result2 && result2 == result3
		fmt.Printf("算法一致性: %t (BFS:%d, 双向:%d, A*:%d)\n",
			consistent, result1, result2, result3)

		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 50))
	}
}

// 性能测试
func benchmarkBusRoutes() {
	fmt.Println()
	fmt.Println("=== 性能测试 ===")
	fmt.Println()

	// 构造测试数据
	testData := []struct {
		routes [][]int
		source int
		target int
		desc   string
	}{
		{
			generateRoutes(10, 5, 10),
			1, 45,
			"小规模：10条路线",
		},
		{
			generateRoutes(50, 10, 100),
			5, 450,
			"中等规模：50条路线",
		},
		{
			generateRoutes(100, 20, 200),
			10, 950,
			"大规模：100条路线",
		},
		{
			generateLinearRoutes(200),
			1, 399,
			"最坏情况：线性路线",
		},
	}

	algorithms := []struct {
		name string
		fn   func([][]int, int, int) int
	}{
		{"BFS路线图", numBusesToDestination},
		{"双向BFS", numBusesToDestinationBidirectional},
		{"A*搜索", numBusesToDestinationAStar},
	}

	for _, data := range testData {
		fmt.Printf("%s:\n", data.desc)

		// 创建公交系统分析
		bs := newBusSystem(data.routes)
		stats := bs.getStats()
		fmt.Printf("  路线数: %d, 车站数: %d, 平均长度: %.1f, 连通性: %.2f\n",
			stats["total_routes"], stats["total_stations"],
			stats["avg_route_len"], stats["connectivity"])

		for _, algo := range algorithms {
			start := time.Now()
			result := algo.fn(data.routes, data.source, data.target)
			duration := time.Since(start)

			fmt.Printf("  %s: 结果=%d, 耗时=%v\n", algo.name, result, duration)
		}
		fmt.Println()
	}
}

// 生成测试路线
func generateRoutes(numRoutes, avgLength, maxStation int) [][]int {
	routes := make([][]int, numRoutes)

	for i := 0; i < numRoutes; i++ {
		length := avgLength + (i%5 - 2) // 长度在avgLength±2之间变化
		if length < 2 {
			length = 2
		}

		route := make([]int, length)
		start := (i * maxStation / numRoutes) + 1

		for j := 0; j < length; j++ {
			route[j] = start + j*2 // 确保站点分布
		}

		routes[i] = route
	}

	return routes
}

// 生成线性路线（最坏情况）
func generateLinearRoutes(numRoutes int) [][]int {
	routes := make([][]int, numRoutes)

	for i := 0; i < numRoutes; i++ {
		// 每条路线只连接相邻的两个站点
		routes[i] = []int{i*2 + 1, i*2 + 2}
	}

	return routes
}

// 路径可视化
func visualizePath(routes [][]int, source, target int) {
	fmt.Println()
	fmt.Println("=== 路径可视化 ===")

	bs := newBusSystem(routes)
	result := bs.findShortestPath(source, target)

	fmt.Printf("从车站 %d 到车站 %d:\n", source, target)
	fmt.Printf("最少换乘次数: %d\n", result)

	if result == -1 {
		fmt.Println("无法到达目标车站")
		return
	}

	fmt.Println("\n路线信息:")
	for i, route := range routes {
		fmt.Printf("路线 %d: %v\n", i, route)

		hasSource := contains(route, source)
		hasTarget := contains(route, target)

		if hasSource && hasTarget {
			fmt.Printf("  * 直达路线!\n")
		} else if hasSource {
			fmt.Printf("  * 包含起点\n")
		} else if hasTarget {
			fmt.Printf("  * 包含终点\n")
		}
	}

	// 显示连通信息
	fmt.Println("\n路线连通性:")
	for i := 0; i < len(routes); i++ {
		connections := []int{}
		for j := 0; j < len(routes); j++ {
			if i != j && hasCommonStation(routes[i], routes[j]) {
				connections = append(connections, j)
			}
		}
		if len(connections) > 0 {
			fmt.Printf("路线 %d 连接到: %v\n", i, connections)
		}
	}
}

// 算法比较演示
func demonstrateAlgorithms() {
	fmt.Println()
	fmt.Println("=== 算法实现对比 ===")

	routes := [][]int{
		{1, 2, 7},
		{3, 6, 7},
		{2, 4, 6},
		{4, 8, 9},
	}
	source, target := 1, 9

	fmt.Printf("测试路线: %v\n", routes)
	fmt.Printf("起点: %d, 终点: %d\n", source, target)

	algorithms := []struct {
		name string
		fn   func([][]int, int, int) int
		desc string
	}{
		{"BFS路线图", numBusesToDestination, "以路线为节点的图搜索"},
		{"双向BFS", numBusesToDestinationBidirectional, "从两端同时搜索"},
		{"A*搜索", numBusesToDestinationAStar, "启发式搜索优化"},
	}

	for _, algo := range algorithms {
		start := time.Now()
		result := algo.fn(routes, source, target)
		duration := time.Since(start)

		fmt.Printf("\n%s (%s):\n", algo.name, algo.desc)
		fmt.Printf("  结果: %d\n", result)
		fmt.Printf("  耗时: %v\n", duration)
	}
}

func main() {
	fmt.Println("815. 公交路线 - 多种解法实现")
	fmt.Println("==============================")

	// 基础功能测试
	testBusRoutes()

	// 性能对比测试
	benchmarkBusRoutes()

	// 路径可视化
	routes := [][]int{{1, 2, 7}, {3, 6, 7}}
	visualizePath(routes, 1, 6)

	// 算法对比演示
	demonstrateAlgorithms()

	// 展示算法特点
	fmt.Println()
	fmt.Println("=== 算法特点分析 ===")
	fmt.Println("1. BFS路线图：以路线为节点，保证最短路径")
	fmt.Println("2. 双向BFS：从两端搜索，理论上快一倍")
	fmt.Println("3. A*搜索：启发式优化，适合特定场景")
	fmt.Println("4. 车站图：直观但空间复杂度较高")

	fmt.Println()
	fmt.Println("=== 公交路线问题技巧 ===")
	fmt.Println("• 图建模：合理选择节点类型(路线vs车站)")
	fmt.Println("• BFS搜索：保证找到最少换乘次数")
	fmt.Println("• 预处理：构建高效的数据结构")
	fmt.Println("• 优化策略：双向搜索和启发式方法")
	fmt.Println("• 边界处理：起终点相同、无法到达等情况")
	fmt.Println("• 系统设计：可扩展的公交系统架构")
}
