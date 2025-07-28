package main

import "fmt"

// 方法一：Kahn算法（拓扑排序）
// 时间复杂度：O(V + E)，空间复杂度：O(V + E)
func findOrder1(numCourses int, prerequisites [][]int) []int {
	// 构建邻接表和入度数组
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	// 将所有入度为0的节点加入队列
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)
	count := 0

	// BFS遍历
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)
		count++

		// 更新相邻节点的入度
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if count == numCourses {
		return result
	}
	return []int{}
}

// 方法二：DFS + 拓扑排序
// 时间复杂度：O(V + E)，空间复杂度：O(V + E)
func findOrder2(numCourses int, prerequisites [][]int) []int {
	// 构建邻接表
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		graph[from] = append(graph[from], to)
	}

	// 状态数组：0=未访问，1=访问中，2=已访问
	visited := make([]int, numCourses)
	result := make([]int, numCourses)
	index := numCourses - 1

	// DFS遍历所有节点
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfs(graph, i, visited, result, &index) {
				return []int{} // 检测到环
			}
		}
	}

	return result
}

func dfs(graph [][]int, node int, visited []int, result []int, index *int) bool {
	visited[node] = 1 // 标记为访问中

	for _, neighbor := range graph[node] {
		if visited[neighbor] == 1 {
			return false // 检测到环
		}
		if visited[neighbor] == 0 {
			if !dfs(graph, neighbor, visited, result, index) {
				return false
			}
		}
	}

	visited[node] = 2 // 标记为已访问
	result[*index] = node
	*index--

	return true
}

// 方法三：优化的Kahn算法
// 时间复杂度：O(V + E)，空间复杂度：O(V + E)
func findOrder3(numCourses int, prerequisites [][]int) []int {
	// 构建邻接表和入度数组
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	// 使用栈进行遍历
	stack := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			stack = append(stack, i)
		}
	}

	result := make([]int, 0)
	count := 0

	// 栈遍历
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current)
		count++

		// 更新相邻节点的入度
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				stack = append(stack, neighbor)
			}
		}
	}

	if count == numCourses {
		return result
	}
	return []int{}
}

// 方法四：改进的DFS（使用递归栈）
// 时间复杂度：O(V + E)，空间复杂度：O(V + E)
func findOrder4(numCourses int, prerequisites [][]int) []int {
	// 构建邻接表
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		graph[from] = append(graph[from], to)
	}

	// 状态数组：0=未访问，1=访问中，2=已访问
	visited := make([]int, numCourses)
	result := make([]int, numCourses)
	index := numCourses - 1

	// DFS遍历所有节点
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfsImproved(graph, i, visited, result, &index) {
				return []int{} // 检测到环
			}
		}
	}

	return result
}

func dfsImproved(graph [][]int, node int, visited []int, result []int, index *int) bool {
	visited[node] = 1 // 标记为访问中

	for _, neighbor := range graph[node] {
		if visited[neighbor] == 1 {
			return false // 检测到环
		}
		if visited[neighbor] == 0 {
			if !dfsImproved(graph, neighbor, visited, result, index) {
				return false
			}
		}
	}

	visited[node] = 2 // 标记为已访问
	result[*index] = node
	*index--

	return true
}

// 方法五：使用map优化邻接表
// 时间复杂度：O(V + E)，空间复杂度：O(V + E)
func findOrder5(numCourses int, prerequisites [][]int) []int {
	// 使用map构建邻接表
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses)

	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	// 将所有入度为0的节点加入队列
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)
	count := 0

	// BFS遍历
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)
		count++

		// 更新相邻节点的入度
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if count == numCourses {
		return result
	}
	return []int{}
}

func main() {
	fmt.Println("=== 210. 课程表 II ===")

	// 测试用例1
	numCourses1 := 2
	prerequisites1 := [][]int{{1, 0}}
	fmt.Printf("测试用例1: numCourses=%d, prerequisites=%v\n", numCourses1, prerequisites1)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses1, prerequisites1))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses1, prerequisites1))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses1, prerequisites1))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses1, prerequisites1))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses1, prerequisites1))
	fmt.Println()

	// 测试用例2
	numCourses2 := 4
	prerequisites2 := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Printf("测试用例2: numCourses=%d, prerequisites=%v\n", numCourses2, prerequisites2)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses2, prerequisites2))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses2, prerequisites2))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses2, prerequisites2))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses2, prerequisites2))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses2, prerequisites2))
	fmt.Println()

	// 测试用例3
	numCourses3 := 1
	prerequisites3 := [][]int{}
	fmt.Printf("测试用例3: numCourses=%d, prerequisites=%v\n", numCourses3, prerequisites3)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses3, prerequisites3))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses3, prerequisites3))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses3, prerequisites3))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses3, prerequisites3))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses3, prerequisites3))
	fmt.Println()

	// 测试用例4（有环）
	numCourses4 := 3
	prerequisites4 := [][]int{{1, 0}, {1, 2}, {0, 1}}
	fmt.Printf("测试用例4: numCourses=%d, prerequisites=%v\n", numCourses4, prerequisites4)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses4, prerequisites4))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses4, prerequisites4))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses4, prerequisites4))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses4, prerequisites4))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses4, prerequisites4))
	fmt.Println()

	// 额外测试用例
	numCourses5 := 3
	prerequisites5 := [][]int{{0, 1}, {0, 2}, {1, 2}}
	fmt.Printf("额外测试: numCourses=%d, prerequisites=%v\n", numCourses5, prerequisites5)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses5, prerequisites5))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses5, prerequisites5))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses5, prerequisites5))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses5, prerequisites5))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses5, prerequisites5))
	fmt.Println()

	// 边界测试用例
	numCourses6 := 0
	prerequisites6 := [][]int{}
	fmt.Printf("边界测试: numCourses=%d, prerequisites=%v\n", numCourses6, prerequisites6)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses6, prerequisites6))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses6, prerequisites6))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses6, prerequisites6))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses6, prerequisites6))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses6, prerequisites6))
	fmt.Println()

	// 复杂测试用例
	numCourses7 := 6
	prerequisites7 := [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}, {5, 4}}
	fmt.Printf("复杂测试: numCourses=%d, prerequisites=%v\n", numCourses7, prerequisites7)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses7, prerequisites7))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses7, prerequisites7))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses7, prerequisites7))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses7, prerequisites7))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses7, prerequisites7))
	fmt.Println()

	// 无依赖测试用例
	numCourses8 := 4
	prerequisites8 := [][]int{}
	fmt.Printf("无依赖测试: numCourses=%d, prerequisites=%v\n", numCourses8, prerequisites8)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses8, prerequisites8))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses8, prerequisites8))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses8, prerequisites8))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses8, prerequisites8))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses8, prerequisites8))
	fmt.Println()

	// 自环测试用例
	numCourses9 := 2
	prerequisites9 := [][]int{{0, 0}}
	fmt.Printf("自环测试: numCourses=%d, prerequisites=%v\n", numCourses9, prerequisites9)
	fmt.Printf("方法一结果: %v\n", findOrder1(numCourses9, prerequisites9))
	fmt.Printf("方法二结果: %v\n", findOrder2(numCourses9, prerequisites9))
	fmt.Printf("方法三结果: %v\n", findOrder3(numCourses9, prerequisites9))
	fmt.Printf("方法四结果: %v\n", findOrder4(numCourses9, prerequisites9))
	fmt.Printf("方法五结果: %v\n", findOrder5(numCourses9, prerequisites9))
}
