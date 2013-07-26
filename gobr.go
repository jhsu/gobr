package main

import (
	"github.com/jhsu/gobr/local"
	"github.com/nsf/termbox-go"
)

func redraw(branches []string) {
	for line, branch := range branches {
		draw_line(line, branch)
	}
}

func draw_line(line int, text string) {
	x := 0
	for _, c := range text {
		termbox.SetCell(x, line, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
}

func select_line(line int, text string) {
	x := 0
	for _, c := range text {
		termbox.SetCell(x, line, c, termbox.ColorBlack, termbox.ColorWhite)
		x++
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	branches := gobr.Branches()
	cb := 0
	redraw(branches)
	select_line(cb, branches[cb])
	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 113:
				break loop
			case 106: // down
				if cb < (len(branches) - 1) {
					cb++
					redraw(branches)
					select_line(cb, branches[cb])
				}
			case 107: // up
				if cb > 0 {
					cb--
					redraw(branches)
					select_line(cb, branches[cb])
				}
			}

			switch ev.Key {
			case termbox.KeyEnter:
				gobr.SetBranch(branches[cb])
				break loop
			}

			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
