package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var a, b, c, d, e, error, private int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "~")
		if len(parts) != 2 {
			error++
			continue
		}

		ip, mask := parts[0], parts[1]

		// 先检查IP格式
		if !isValidIP(ip) {
			error++
			continue
		}

		ipParts := strings.Split(ip, ".")
		first, _ := strconv.Atoi(ipParts[0])

		// 特殊IP处理：0.*.*.* 和 127.*.*.* 直接跳过，不检查掩码
		if first == 0 || first == 127 {
			continue
		}

		// 检查掩码格式
		if !isValidMask(mask) {
			error++
			continue
		}

		// 检查私有IP
		if isPrivateIP(ip) {
			private++
		}

		// 分类统计
		if first >= 1 && first <= 127 {
			a++
		} else if first >= 128 && first <= 191 {
			b++
		} else if first >= 192 && first <= 223 {
			c++
		} else if first >= 224 && first <= 239 {
			d++
		} else if first >= 240 && first <= 255 {
			e++
		}
	}

	fmt.Printf("%d %d %d %d %d %d %d", a, b, c, d, e, error, private)
}

func isValidIP(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if part == "" {
			return false
		}
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func isValidMask(mask string) bool {
	parts := strings.Split(mask, ".")
	if len(parts) != 4 {
		return false
	}

	var binary string
	for _, part := range parts {
		if part == "" {
			return false
		}
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
		binary += fmt.Sprintf("%08b", num)
	}

	if binary == "11111111111111111111111111111111" || binary == "00000000000000000000000000000000" {
		return false
	}

	oneFound := false
	zeroFound := false
	for _, bit := range binary {
		if bit == '1' {
			if zeroFound {
				return false
			}
			oneFound = true
		} else {
			zeroFound = true
		}
	}

	return oneFound && zeroFound
}

func isPrivateIP(ip string) bool {
	parts := strings.Split(ip, ".")
	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[1])

	if first == 10 {
		return true
	}

	if first == 172 && second >= 16 && second <= 31 {
		return true
	}

	if first == 192 && second == 168 {
		return true
	}

	return false
}
