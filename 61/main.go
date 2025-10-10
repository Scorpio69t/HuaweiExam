package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// ===== 方法一：闭环取模 + 断开（最优常规解） =====
// 思路：
// 1) 边界：空/单节点/ k==0 -> 原链表
// 2) 计算长度 len，并找到尾节点 tail，使 tail.Next=head，形成环
// 3) k %= len，若 k==0 直接断环返回
// 4) 新尾在第 len-k 个节点，新头为新尾的 Next，将新尾.Next 置空
// 时间 O(n)，空间 O(1)
func rotateRightRing(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 计算长度并拿到尾
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	// 闭环
	tail.Next = head
	// 新尾: 走 n-k-1 步到达；新头: 新尾.Next
	steps := n - k - 1
	newTail := head
	for i := 0; i < steps; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

// ===== 方法二：快慢指针（双指针等效解） =====
// 让 fast 先走 k%len 步，然后 slow 与 fast 同步前进直到 fast 到尾；
// slow 停在“新尾”，slow.Next 为“新头”。最后把尾部与头断开并连接尾到原头。
// 时间 O(n)，空间 O(1)
func rotateRightTwoPointers(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 计算长度
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// 同步走到尾
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// slow 是新尾，slow.Next 新头
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

// ===== 方法三：数组收集节点指针（易理解） =====
// 将节点指针装入切片，按索引重连；空间 O(n)，适合讲解与对拍
func rotateRightArray(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	nodes := make([]*ListNode, 0, 64)
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	n := len(nodes)
	k %= n
	if k == 0 {
		return head
	}
	// 新头索引
	newHeadIdx := (n - k) % n
	newHead := nodes[newHeadIdx]
	// 断开与重连
	nodes[(newHeadIdx-1+n)%n].Next = nil
	nodes[n-1].Next = nodes[0]
	return newHead
}

// ===== 构建/打印/辅助 =====
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

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func main() {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
		{[]int{}, 5, []int{}},
		{[]int{1}, 0, []int{1}},
		{[]int{1}, 3, []int{1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2}, 3, []int{2, 1}},
	}

	methods := []struct {
		name string
		fn   func(*ListNode, int) *ListNode
	}{
		{"闭环取模", rotateRightRing},
		{"快慢双指针", rotateRightTwoPointers},
		{"数组重连", rotateRightArray},
	}

	fmt.Println("61. 旋转链表 - 多解法对比")
	for _, c := range cases {
		fmt.Printf("in=%v k=%d\n", c.in, c.k)
		for _, m := range methods {
			got := listToSlice(m.fn(buildList(c.in), c.k))
			status := "✅"
			if fmt.Sprint(got) != fmt.Sprint(c.want) {
				status = "❌"
			}
			fmt.Printf("  %-8s => %v %s\n", m.name, got, status)
		}
		fmt.Printf("  期望 => %v\n", c.want)
		fmt.Println("------------------------------")
	}
}
