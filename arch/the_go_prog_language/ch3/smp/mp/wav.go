package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing wav music", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Microsecond)
		fmt.Print("...")
		p.progress += 10
	}

	fmt.Println("Played", source)
}
