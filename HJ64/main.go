package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var commands string

	fmt.Scanln(&n)
	fmt.Scanln(&commands)

	// 模拟MP3播放器状态
	player := NewMP3Player(n)

	// 执行命令
	for _, cmd := range commands {
		switch cmd {
		case 'U':
			player.MoveUp()
		case 'D':
			player.MoveDown()
		}
	}

	// 输出结果
	player.PrintResult()
}

// MP3Player 模拟MP3播放器
type MP3Player struct {
	totalSongs int // 总歌曲数
	cursor     int // 当前光标位置（1-based）
	pageSize   int // 每页显示歌曲数
	startSong  int // 当前页第一首歌曲
	endSong    int // 当前页最后一首歌曲
}

// NewMP3Player 创建新的MP3播放器
func NewMP3Player(totalSongs int) *MP3Player {
	player := &MP3Player{
		totalSongs: totalSongs,
		cursor:     1,
		pageSize:   4,
		startSong:  1,
		endSong:    min(4, totalSongs),
	}
	return player
}

// MoveUp 向上移动光标
func (p *MP3Player) MoveUp() {
	if p.totalSongs <= p.pageSize {
		// 歌曲总数<=4，不需要翻页
		if p.cursor == 1 {
			p.cursor = p.totalSongs
		} else {
			p.cursor--
		}
	} else {
		// 歌曲总数>4，需要处理翻页
		if p.cursor == p.startSong {
			// 光标在当前页第一首，需要翻页
			if p.startSong == 1 {
				// 当前是第一页，翻到最后一页
				p.startSong = p.totalSongs - p.pageSize + 1
				p.endSong = p.totalSongs
				p.cursor = p.totalSongs
			} else {
				// 翻到上一页
				p.startSong--
				p.endSong--
				p.cursor--
			}
		} else {
			// 光标不在第一首，直接向上移动
			p.cursor--
		}
	}
}

// MoveDown 向下移动光标
func (p *MP3Player) MoveDown() {
	if p.totalSongs <= p.pageSize {
		// 歌曲总数<=4，不需要翻页
		if p.cursor == p.totalSongs {
			p.cursor = 1
		} else {
			p.cursor++
		}
	} else {
		// 歌曲总数>4，需要处理翻页
		if p.cursor == p.endSong {
			// 光标在当前页最后一首，需要翻页
			if p.endSong == p.totalSongs {
				// 当前是最后一页，翻到第一页
				p.startSong = 1
				p.endSong = p.pageSize
				p.cursor = 1
			} else {
				// 翻到下一页
				p.startSong++
				p.endSong++
				p.cursor++
			}
		} else {
			// 光标不在最后一首，直接向下移动
			p.cursor++
		}
	}
}

// PrintResult 打印当前页面和光标位置
func (p *MP3Player) PrintResult() {
	// 输出当前页面显示的歌曲
	var songs []string
	for i := p.startSong; i <= p.endSong; i++ {
		songs = append(songs, fmt.Sprintf("%d", i))
	}
	fmt.Println(strings.Join(songs, " "))

	// 输出当前光标位置
	fmt.Println(p.cursor)
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
