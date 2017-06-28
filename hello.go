package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// START OMIT
type LineTicker struct {
	C <-chan string
}

func NewLineTicker(reader io.ReadCloser) *LineTicker {
	lineChan := make(chan string) // HLchan
	l := &LineTicker{C: lineChan} // HLescape
	go func() {                   // HLgo
		defer reader.Close() // HLdefer
		ticker := time.Tick(200 * time.Millisecond)
		scanner := bufio.NewScanner(reader)
		for _ = range ticker {
			if !scanner.Scan() {
				close(lineChan) // HLchan
				return
			}
			lineChan <- scanner.Text() // HLchan
		}
	}() // HLgo
	return l
}

// END OMIT

func main() {
	worldsFile, _ := os.Open("world.txt")
	worldTicker := NewLineTicker(worldsFile)
	for world := range worldTicker.C {
		fmt.Println("Hello, " + world)
	}
}
