# 210. 课程表 II

## 描述

现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

## 示例 1

输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。

## 示例 2

输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。

## 示例 3

输入：numCourses = 1, prerequisites = []
输出：[0]

## 示例 4

输入：numCourses = 3, prerequisites = [[1,0],[1,2],[0,1]]
输出：[]
解释：总共有 3 门课程。要学习课程 1，你需要先完成课程 0 和课程 2。但是要学习课程 0，你需要先完成课程 1，这形成了循环依赖，无法完成所有课程。

## 提示

- 1 <= numCourses <= 2000
- 0 <= prerequisites.length <= numCourses * (numCourses - 1)
- prerequisites[i].length == 2
- 0 <= ai, bi < numCourses
- ai != bi
- 所有[ai, bi] 互不相同

## 解题思路

### 方法一：Kahn算法（拓扑排序）

**核心思想**：
- 使用Kahn算法进行拓扑排序
- 维护每个节点的入度，从入度为0的节点开始
- 使用队列进行BFS遍历

**算法步骤**：
1. 构建邻接表和入度数组
2. 将所有入度为0的节点加入队列
3. 从队列中取出节点，加入结果数组
4. 更新相邻节点的入度，如果入度变为0则加入队列
5. 检查是否所有节点都被访问

**时间复杂度**：O(V + E)
**空间复杂度**：O(V + E)

### 方法二：DFS + 拓扑排序

**核心思想**：
- 使用DFS进行拓扑排序
- 使用三种状态标记节点：未访问(0)、访问中(1)、已访问(2)
- 检测环的存在

**算法步骤**：
1. 构建邻接表
2. 使用DFS遍历所有节点
3. 使用状态数组检测环
4. 按DFS完成顺序构建结果

**时间复杂度**：O(V + E)
**空间复杂度**：O(V + E)

### 方法三：优化的Kahn算法

**核心思想**：
- 优化Kahn算法的实现
- 使用更高效的数据结构

**算法步骤**：
1. 构建邻接表和入度数组
2. 使用栈或队列进行遍历
3. 优化内存使用

**时间复杂度**：O(V + E)
**空间复杂度**：O(V + E)

### 方法四：并查集（不适用）

**注意**：并查集不适用于有向图的拓扑排序，因为拓扑排序需要保持依赖关系的方向性。

## 代码实现

```go
// 方法一：Kahn算法
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
```

## 复杂度分析

- **时间复杂度**：O(V + E)，其中V是节点数，E是边数
- **空间复杂度**：O(V + E)，邻接表和队列的空间

## 测试用例

```go
func main() {
    // 测试用例1
    numCourses1 := 2
    prerequisites1 := [][]int{{1, 0}}
    fmt.Printf("测试用例1: numCourses=%d, prerequisites=%v, 结果=%v\n", 
               numCourses1, prerequisites1, findOrder1(numCourses1, prerequisites1))
    
    // 测试用例2
    numCourses2 := 4
    prerequisites2 := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
    fmt.Printf("测试用例2: numCourses=%d, prerequisites=%v, 结果=%v\n", 
               numCourses2, prerequisites2, findOrder1(numCourses2, prerequisites2))
    
    // 测试用例3
    numCourses3 := 1
    prerequisites3 := [][]int{}
    fmt.Printf("测试用例3: numCourses=%d, prerequisites=%v, 结果=%v\n", 
               numCourses3, prerequisites3, findOrder1(numCourses3, prerequisites3))
}
```
