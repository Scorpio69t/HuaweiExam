package main

import (
	"fmt"
	"sort"
)

// Student 学生结构体
type Student struct {
	Name  string // 学生姓名
	Score int    // 学生成绩
	Index int    // 输入顺序，用于稳定排序
}

// StudentList 学生列表管理
type StudentList struct {
	Students  []Student // 学生列表
	SortOrder int       // 排序方式：0-降序，1-升序
}

// Add 添加学生
func (sl *StudentList) Add(name string, score int, index int) {
	student := Student{
		Name:  name,
		Score: score,
		Index: index,
	}
	sl.Students = append(sl.Students, student)
}

// Sort 执行稳定排序
func (sl *StudentList) Sort() {
	sort.SliceStable(sl.Students, func(i, j int) bool {
		// 如果成绩相同，按输入顺序排序（保持稳定性）
		if sl.Students[i].Score == sl.Students[j].Score {
			return sl.Students[i].Index < sl.Students[j].Index
		}

		// 根据排序方式决定升序或降序
		if sl.SortOrder == 0 {
			// 降序：成绩大的在前
			return sl.Students[i].Score > sl.Students[j].Score
		} else {
			// 升序：成绩小的在前
			return sl.Students[i].Score < sl.Students[j].Score
		}
	})
}

// Display 输出排序结果
func (sl *StudentList) Display() {
	for _, student := range sl.Students {
		fmt.Printf("%s %d\n", student.Name, student.Score)
	}
}

func main() {
	var n, op int

	// 读取学生数量和排序方式
	fmt.Scan(&n)
	fmt.Scan(&op)

	// 创建学生列表
	studentList := StudentList{
		Students:  make([]Student, 0, n),
		SortOrder: op,
	}

	// 读取学生信息
	for i := 0; i < n; i++ {
		var name string
		var score int
		fmt.Scan(&name, &score)

		// 添加学生，记录输入顺序
		studentList.Add(name, score, i)
	}

	// 执行稳定排序
	studentList.Sort()

	// 输出排序结果
	studentList.Display()
}
