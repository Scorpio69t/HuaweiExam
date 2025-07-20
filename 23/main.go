package main

import (
	"container/heap"
	"fmt"
	"sort"
	"time"
)

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：分治合并法（推荐）
// 时间复杂度：O(n log k)，空间复杂度：O(log k)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	return divide(lists, 0, len(lists)-1)
}

// 分治递归函数
func divide(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := left + (right-left)/2
	leftList := divide(lists, left, mid)
	rightList := divide(lists, mid+1, right)

	return mergeTwoLists(leftList, rightList)
}

// 合并两个有序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// 连接剩余节点
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

// 方法二：优先队列法（最小堆）
// 时间复杂度：O(n log k)，空间复杂度：O(k)
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKListsHeap(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// 初始化最小堆
	h := &NodeHeap{}
	heap.Init(h)

	// 将所有非空链表的头节点加入堆
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	// 依次取出最小节点
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next

		// 如果该节点有后继，将后继加入堆
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

// 方法三：逐一合并法
// 时间复杂度：O(k*n)，空间复杂度：O(1)
func mergeKListsSequential(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}

	return result
}

// 方法四：数组排序法
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func mergeKListsArray(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// 收集所有节点值
	var values []int
	for _, list := range lists {
		current := list
		for current != nil {
			values = append(values, current.Val)
			current = current.Next
		}
	}

	if len(values) == 0 {
		return nil
	}

	// 排序
	sort.Ints(values)

	// 重建链表
	dummy := &ListNode{}
	current := dummy

	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// 辅助函数：创建链表
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy

	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// 辅助函数：打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print("->")
		}
		current = current.Next
	}
	fmt.Println()
}

// 辅助函数：链表转数组（用于验证结果）
func listToArray(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// 辅助函数：数组比较
func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 性能测试函数
func measureTime(name string, fn func() *ListNode) *ListNode {
	start := time.Now()
	result := fn()
	duration := time.Since(start)
	fmt.Printf("%s: 耗时=%v\n", name, duration)
	return result
}

// 运行所有测试用例
func runTests() {
	fmt.Println("=== 23. 合并K个升序链表 测试用例 ===")

	// 测试用例1：示例1
	fmt.Println("\n--- 测试用例1：示例1 ---")
	lists1 := []*ListNode{
		createList([]int{1, 4, 5}),
		createList([]int{1, 3, 4}),
		createList([]int{2, 6}),
	}

	fmt.Println("输入链表：")
	for i, list := range lists1 {
		fmt.Printf("链表%d: ", i+1)
		printList(list)
	}

	// 测试所有方法
	result1_1 := mergeKLists(copyLists(lists1))
	result1_2 := mergeKListsHeap(copyLists(lists1))
	result1_3 := mergeKListsSequential(copyLists(lists1))
	result1_4 := mergeKListsArray(copyLists(lists1))

	fmt.Println("结果：")
	fmt.Print("分治合并法: ")
	printList(result1_1)
	fmt.Print("优先队列法: ")
	printList(result1_2)
	fmt.Print("逐一合并法: ")
	printList(result1_3)
	fmt.Print("数组排序法: ")
	printList(result1_4)

	// 验证结果一致性
	expected1 := []int{1, 1, 2, 3, 4, 4, 5, 6}
	if arraysEqual(listToArray(result1_1), expected1) &&
		arraysEqual(listToArray(result1_2), expected1) &&
		arraysEqual(listToArray(result1_3), expected1) &&
		arraysEqual(listToArray(result1_4), expected1) {
		fmt.Println("✅ 所有方法结果正确且一致")
	} else {
		fmt.Println("❌ 结果不一致！")
	}

	// 测试用例2：空数组
	fmt.Println("\n--- 测试用例2：空数组 ---")
	lists2 := []*ListNode{}
	result2 := mergeKLists(lists2)
	fmt.Printf("结果: %v (预期: null)\n", result2)

	// 测试用例3：包含空链表
	fmt.Println("\n--- 测试用例3：包含空链表 ---")
	lists3 := []*ListNode{nil}
	result3 := mergeKLists(lists3)
	fmt.Printf("结果: %v (预期: null)\n", result3)

	// 测试用例4：单个链表
	fmt.Println("\n--- 测试用例4：单个链表 ---")
	lists4 := []*ListNode{createList([]int{1, 2, 3})}
	result4 := mergeKLists(lists4)
	fmt.Print("结果: ")
	printList(result4)

	// 测试用例5：重复元素
	fmt.Println("\n--- 测试用例5：重复元素 ---")
	lists5 := []*ListNode{
		createList([]int{1, 1, 2}),
		createList([]int{1, 1, 2}),
	}
	result5 := mergeKLists(lists5)
	fmt.Print("结果: ")
	printList(result5)

	// 性能测试
	fmt.Println("\n--- 性能对比测试 ---")
	performanceTest()
}

// 复制链表数组（因为合并会修改原链表）
func copyLists(lists []*ListNode) []*ListNode {
	copied := make([]*ListNode, len(lists))
	for i, list := range lists {
		copied[i] = copyList(list)
	}
	return copied
}

// 复制单个链表
func copyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	current := dummy

	for head != nil {
		current.Next = &ListNode{Val: head.Val}
		current = current.Next
		head = head.Next
	}

	return dummy.Next
}

// 性能对比测试
func performanceTest() {
	// 创建测试数据
	testLists := []*ListNode{
		createList([]int{1, 4, 7, 10, 13}),
		createList([]int{2, 5, 8, 11, 14}),
		createList([]int{3, 6, 9, 12, 15}),
		createList([]int{16, 17, 18, 19, 20}),
		createList([]int{21, 22, 23, 24, 25}),
	}

	fmt.Printf("测试数据：5个链表，每个链表5个元素\n")

	// 测试分治合并法
	result1 := measureTime("分治合并法", func() *ListNode {
		return mergeKLists(copyLists(testLists))
	})

	// 测试优先队列法
	result2 := measureTime("优先队列法", func() *ListNode {
		return mergeKListsHeap(copyLists(testLists))
	})

	// 测试逐一合并法
	result3 := measureTime("逐一合并法", func() *ListNode {
		return mergeKListsSequential(copyLists(testLists))
	})

	// 测试数组排序法
	result4 := measureTime("数组排序法", func() *ListNode {
		return mergeKListsArray(copyLists(testLists))
	})

	// 验证结果一致性
	arr1 := listToArray(result1)
	arr2 := listToArray(result2)
	arr3 := listToArray(result3)
	arr4 := listToArray(result4)

	if arraysEqual(arr1, arr2) && arraysEqual(arr2, arr3) && arraysEqual(arr3, arr4) {
		fmt.Println("✅ 所有方法结果一致")
		fmt.Print("合并结果: ")
		printList(result1)
	} else {
		fmt.Println("❌ 结果不一致！")
	}
}

// 展示分治过程
func demonstrateDivideConquer() {
	fmt.Println("\n--- 分治过程演示 ---")

	lists := []*ListNode{
		createList([]int{1, 4}),
		createList([]int{1, 3}),
		createList([]int{2, 6}),
		createList([]int{5, 7}),
	}

	fmt.Println("原始4个链表：")
	for i, list := range lists {
		fmt.Printf("链表%d: ", i+1)
		printList(list)
	}

	fmt.Println("\n分治合并过程：")
	fmt.Println("第1轮：合并相邻链表对")
	merged1 := mergeTwoLists(copyList(lists[0]), copyList(lists[1]))
	merged2 := mergeTwoLists(copyList(lists[2]), copyList(lists[3]))
	fmt.Print("合并链表1+2: ")
	printList(merged1)
	fmt.Print("合并链表3+4: ")
	printList(merged2)

	fmt.Println("\n第2轮：合并结果")
	final := mergeTwoLists(merged1, merged2)
	fmt.Print("最终结果: ")
	printList(final)
}

func main() {
	runTests()
	demonstrateDivideConquer()
}
