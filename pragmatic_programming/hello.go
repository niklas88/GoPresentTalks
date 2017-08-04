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
	C chan string // HLstruct
	reader io.ReadCloser // HLstruct
}

func (l *LineTicker) Start() {
	l.C = make(chan string) // HLchan
	go func() {                   // HLgo
		defer close(l.C) // HLdefer
		ticker := time.Tick(200 * time.Millisecond)
		scanner := bufio.NewScanner(l.reader)
		for range ticker {
			if !scanner.Scan() {
				return
			}
			l.C <- scanner.Text() // HLchan
		}
	}() // HLgo
}
// END OMIT

// START NOTICKER OMIT
func NewLineTicker(input io.ReadCloser) *LineTicker {
	l := &LineTicker{reader: input} // HLescape
	return l
}

func main() {
	worldsFile, err := os.Open("world.txt")
	if err != nil {
		return
	}
	defer worldsFile.Close() // HLdefer
	worldTicker := NewLineTicker(worldsFile)
	worldTicker.Start()
	for world := range worldTicker.C {
		fmt.Println("Hello, " + world)
	}
}
// END NOTICKER OMIT
