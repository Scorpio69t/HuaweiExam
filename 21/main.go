package main

import (
	"fmt"
)

// 单链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法1：迭代-哨兵节点（推荐）：时间 O(m+n)，空间 O(1)
func mergeTwoListsIterative(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

// 解法2：递归：时间 O(m+n)，空间 O(m+n)（递归栈）
func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRecursive(l1, l2.Next)
	return l2
}

// 解法3：原地指针连接（本质与迭代等价，演示另一种写法）
func mergeTwoListsInPlace(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, tail *ListNode
	if l1.Val <= l2.Val {
		head, tail, l1 = l1, l1, l1.Next
	} else {
		head, tail, l2 = l2, l2, l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}
	return head
}

// 工具函数：由切片构建链表
func buildList(nums []int) *ListNode {
	var head, cur *ListNode
	for _, v := range nums {
		node := &ListNode{Val: v}
		if head == nil {
			head = node
			cur = node
		} else {
			cur.Next = node
			cur = node
		}
	}
	return head
}

// 工具函数：链表转切片
func listToSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func runTests() {
	fmt.Println("=== 21. 合并两个有序链表 测试 ===")

	cases := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
		{[]int{-10, -3, 0, 5}, []int{-5, -3, 2}, []int{-10, -5, -3, -3, 0, 2, 5}},
	}

	for i, c := range cases {
		l1 := buildList(c.l1)
		l2 := buildList(c.l2)
		got1 := listToSlice(mergeTwoListsIterative(l1, l2))
		fmt.Printf("用例#%d 迭代: got=%v, want=%v\n", i+1, got1, c.want)

		l1 = buildList(c.l1)
		l2 = buildList(c.l2)
		got2 := listToSlice(mergeTwoListsRecursive(l1, l2))
		fmt.Printf("用例#%d 递归: got=%v, want=%v\n", i+1, got2, c.want)

		l1 = buildList(c.l1)
		l2 = buildList(c.l2)
		got3 := listToSlice(mergeTwoListsInPlace(l1, l2))
		fmt.Printf("用例#%d 原地: got=%v, want=%v\n", i+1, got3, c.want)
	}
}

func main() {
	runTests()
}
// end
