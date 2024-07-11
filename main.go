package main

import (
	"fmt"
	"go-grass-cli/cmd"
	"log"
	"os"
	"sync"
	"time"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

func MoveCursor(x, y int) string {
	return fmt.Sprintf("\033[%d;%dH", y, x*2)
}

var width int = 0
var height int = 0
var fd = int(os.Stdout.Fd())

func getTermSize() (int, int) {
	w, h, _ := term.GetSize(fd)
	return w / 2, h
}

func main() {
	width, height = getTermSize()

	gameState := cmd.NewGameState(width, height)
	gameState.Setup()

	quit := make(chan bool)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go inputLoop(&waitGroup, quit)
	go gameLoop(&waitGroup, quit, gameState)

	waitGroup.Wait()
	close(quit)
}

func gameLoop(waitGroup *sync.WaitGroup, quit <-chan bool, gameState *cmd.GameState) {
	lastLoopTime := time.Now()
	for {
		select {
		case <-quit:
			waitGroup.Done()
			return
		default:
			newWidth, newHeight := getTermSize()
			if newWidth != width || newHeight != height {
				fmt.Println("resize")
				width = newWidth
				height = newHeight
				gameState.Resize(width, height)
			}
			diff := time.Since(lastLoopTime).Seconds()
			gameState.Loop(diff)
			lastLoopTime = time.Now()
		}
	}
}

func inputLoop(waitGroup *sync.WaitGroup, quit chan bool) {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer tty.Close()

	for {
		readRune, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		// if rune is escape key
		if readRune == 27 {
			quit <- true
			break
		}
	}
	waitGroup.Done()
}
