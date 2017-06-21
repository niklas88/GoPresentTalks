package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func worldTicker() <-chan string {
	ticker := time.Tick(200 * time.Millisecond)
	worlds := make(chan string)
	go func() {
		worldsFile, _ := os.Open("world.txt")
		defer worldsFile.Close()
		scanner := bufio.NewScanner(worldsFile)
		for _ = range ticker {
			if !scanner.Scan() {
				close(worlds)
				return
			}
			worlds <- scanner.Text()
		}
	}()
	return worlds
}

func main() {
	worlds := worldTicker()
	for world := range worlds {
		fmt.Println("Hello, " + world)
	}
}
