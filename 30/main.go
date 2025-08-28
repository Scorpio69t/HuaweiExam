package main

import (
	"fmt"
)

// findSubstring 经典解：所有 words 等长，使用按单词长度分组的多起点滑动窗口
func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(s) == 0 || len(words) == 0 {
		return res
	}
	wlen := len(words[0])
	for _, w := range words {
		if len(w) != wlen {
			return res // 题目保证相等，这里防御
		}
	}
	n := len(s)
	m := len(words)
	total := wlen * m
	if n < total {
		return res
	}

	// 统计 words 词频
	target := make(map[string]int)
	for _, w := range words {
		target[w]++
	}

	// 从 wlen 个不同起点出发
	for offset := 0; offset < wlen; offset++ {
		left := offset
		count := 0
		window := make(map[string]int)
		for right := offset; right+wlen <= n; right += wlen {
			w := s[right : right+wlen]
			if target[w] > 0 {
				window[w]++
				count++
				for window[w] > target[w] { // 收缩到合法
					leftWord := s[left : left+wlen]
					window[leftWord]--
					left += wlen
					count--
				}
				if count == m {
					res = append(res, left)
					// 移动一个词，继续寻找下一个
					leftWord := s[left : left+wlen]
					window[leftWord]--
					left += wlen
					count--
				}
			} else { // 不在词表，重置窗口
				window = make(map[string]int)
				count = 0
				left = right + wlen
			}
		}
	}
	return res
}

// findSubstringHash 简洁实现：直接在每个可能起点切分 m 段，计数匹配
func findSubstringHash(s string, words []string) []int {
	res := []int{}
	if len(s) == 0 || len(words) == 0 {
		return res
	}
	wlen := len(words[0])
	for _, w := range words {
		if len(w) != wlen {
			return res
		}
	}
	n := len(s)
	m := len(words)
	total := wlen * m
	if n < total {
		return res
	}
	target := make(map[string]int)
	for _, w := range words {
		target[w]++
	}
	for i := 0; i+total <= n; i++ {
		seen := make(map[string]int)
		ok := true
		for j := 0; j < m; j++ {
			w := s[i+j*wlen : i+(j+1)*wlen]
			if target[w] == 0 {
				ok = false
				break
			}
			seen[w]++
			if seen[w] > target[w] {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println("串联所有单词的子串 测试")
	fmt.Println("====================")

	s1 := "barfoothefoobarman"
	w1 := []string{"foo", "bar"}
	fmt.Println("用例1:", s1, w1)
	fmt.Println("滑窗: ", findSubstring(s1, w1))
	fmt.Println("哈希: ", findSubstringHash(s1, w1))

	s2 := "wordgoodgoodgoodbestword"
	w2 := []string{"word", "good", "best", "word"}
	fmt.Println("\n用例2:", s2, w2)
	fmt.Println("滑窗: ", findSubstring(s2, w2))
	fmt.Println("哈希: ", findSubstringHash(s2, w2))

	s3 := "barfoofoobarthefoobarman"
	w3 := []string{"bar", "foo", "the"}
	fmt.Println("\n用例3:", s3, w3)
	fmt.Println("滑窗: ", findSubstring(s3, w3))
	fmt.Println("哈希: ", findSubstringHash(s3, w3))
}
