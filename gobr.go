package gobr

import "github.com/nsf/termbox-go"
import "time"
import "fmt"

type key struct {
	ch rune
}

const hello_world = "Hello, World!"

func main() {
	time.Sleep(1 * time.Second)
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	x := 0
	for _, c := range hello_world {
		termbox.SetCell(x, 0, c, termbox.ColorWhite, termbox.ColorDefault)
		x++
	}
	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
      fmt.Println(ev.Ch)
      if ev.Ch == 113 {
        break loop
      }
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
