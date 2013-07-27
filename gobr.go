package main

import (
	"bytes"
	"github.com/nsf/termbox-go"
	"os/exec"
	"strings"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	branches := branches()
	cb := 0
	redraw(branches)
	selectLine(cb, branches[cb])
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
					selectLine(cb, branches[cb])
				}
			case 107: // up
				if cb > 0 {
					cb--
					redraw(branches)
					selectLine(cb, branches[cb])
				}
			}

			switch ev.Key {
			case termbox.KeyEnter:
				setBranch(branches[cb])
				break loop
			}

			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func redraw(branches []string) {
	for line, branch := range branches {
		drawLine(line, branch)
	}
}

func drawLine(line int, text string) {
	x := 0
	for _, c := range text {
		termbox.SetCell(x, line, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
}

func selectLine(line int, text string) {
	x := 0
	for _, c := range text {
		termbox.SetCell(x, line, c, termbox.ColorBlack, termbox.ColorWhite)
		x++
	}
}

// Branches gets a range of local git branches.
func branches() []string {
	cmd := exec.Command("git", "branch")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	branches := strings.Fields(out.String())
	names := []string{}
	for _, name := range branches {
		if name != "" && name != "*" {
			names = append(names, name)
		}
	}
	return names
}

// SetBranch changes the current git branch.
func setBranch(branch string) {
	cmd := exec.Command("git", "checkout", branch)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
