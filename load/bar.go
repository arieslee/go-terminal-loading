package load

import (
	"fmt"
	"os"
	"time"
)

type Bar struct {
	GraphProgress []string
	Delay         time.Duration
}

func New() *Bar {
	bar := &Bar{
		GraphProgress: []string{`\`, `|`, `-`, `/`, `-`},
		Delay:         time.Millisecond * 300,
	}
	return bar
}

// clearCurrentLine 清空当前行
func (bar *Bar) clearCurrentLine() {
	fmt.Printf("\033[0;") // clear current line
	fmt.Printf("\033[2K\r%d", 0)
	fmt.Fprint(os.Stdout, "\033[y;0H")
	fmt.Fprint(os.Stdout, "\033[K")
	fmt.Print("\x1b[2k") // erase the current line
}

// Loading  loading
func (bar *Bar) Loading(styles ...int) {
	var style int
	if len(styles) > 0 {
		style = styles[0]
	} else {
		style = 0
	}
	if style == 2 {
		fmt.Printf("loading:%s\r", bar.getProgress())
	} else {
		dot := `.`
		loadingBar := ""
		for i := 0; i < 5; i++ {
			loadingBar += dot
			if style == 0 {
				fmt.Printf("\rloading:%s", loadingBar)
			} else {
				fmt.Printf("loading:%s\n", loadingBar)
			}
			time.Sleep(bar.Delay)
		}
		if style == 0 {
			bar.clearCurrentLine()
		}
	}
}

var index = 0

func (bar *Bar) getProgress() string {
	index++
	if index == 5 {
		index = 0
	}
	time.Sleep(bar.Delay)
	return bar.GraphProgress[index]
}
