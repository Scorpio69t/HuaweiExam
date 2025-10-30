package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// =========================== 方法一：头插法一趟扫描（最优） ===========================
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	leftNode := pre.Next
	for i := 0; i < right-left; i++ {
		next := leftNode.Next
		leftNode.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

// =========================== 方法二：常规反转子链后接回 ===========================
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	rightNext := rightNode.Next
	// 切下子链
	leftNode := pre.Next
	rightNode.Next = nil
	// 反转子链
	revHead := reverseList(leftNode)
	// 接回
	pre.Next = revHead
	leftNode.Next = rightNext
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// =========================== 方法三：递归（reverseN + 偏移） ===========================
var successor *ListNode

func reverseBetween3(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween3(head.Next, left-1, right-1)
	return head
}

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

// =========================== 方法四：栈辅助 ===========================
func reverseBetween4(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	// 收集区间节点
	stack := []*ListNode{}
	node := pre.Next
	for i := left; i <= right; i++ {
		stack = append(stack, node)
		node = node.Next
	}
	rightNext := node
	// 从栈弹出重建
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pre.Next = n
		pre = pre.Next
	}
	pre.Next = rightNext
	return dummy.Next
}

// =========================== 工具方法与测试 ===========================
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

func toSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func equalSlice(a, b []int) bool {
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

func runTests(methodName string, f func(*ListNode, int, int) *ListNode) int {
	tests := []struct {
		in    []int
		l, r  int
		want  []int
		label string
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}, "mid range"},
		{[]int{5}, 1, 1, []int{5}, "single"},
		{[]int{1, 2, 3}, 1, 3, []int{3, 2, 1}, "all"},
		{[]int{1, 2, 3, 4}, 3, 4, []int{1, 2, 4, 3}, "tail"},
		{[]int{1, 2}, 1, 2, []int{2, 1}, "two"},
		{[]int{1, 2, 3, 4, 5}, 3, 3, []int{1, 2, 3, 4, 5}, "no-op"},
	}
	pass := 0
	for i, tc := range tests {
		head := buildList(tc.in)
		out := f(head, tc.l, tc.r)
		s := toSlice(out)
		ok := equalSlice(s, tc.want)
		status := "✅"
		if !ok {
			status = "❌"
		}
		fmt.Printf("  测试%d(%s): %s\n", i+1, tc.label, status)
		if !ok {
			fmt.Printf("    输入: %v, left=%d, right=%d\n", tc.in, tc.l, tc.r)
			fmt.Printf("    输出: %v\n", s)
			fmt.Printf("    期望: %v\n", tc.want)
		} else {
			pass++
		}
	}
	fmt.Printf("  通过: %d/%d\n\n", pass, len(tests))
	return pass
}

func main() {
	fmt.Println("=== LeetCode 92: 反转链表 II ===\n")
	methods := []struct {
		name string
		fn   func(*ListNode, int, int) *ListNode
	}{
		{"头插法一趟扫描（最优）", reverseBetween1},
		{"常规反转后接回", reverseBetween2},
		{"递归（reverseN）", reverseBetween3},
		{"栈辅助", reverseBetween4},
	}
	for _, m := range methods {
		fmt.Printf("方法：%s\n", m.name)
		runTests(m.name, m.fn)
	}
}
