package main

import (
  "./util"
  "time"
)

func main() {
  // messages := make(chan string)
  // go func() { messages <- "ping" }()
  // msg := <- messages
  // fmt.Println(msg)
  size := util.GetNumLines() - 4
  sleepLen := 1 * time.Second
  util.ClearScreen()
  for true {
    doGeneration(size)
    time.Sleep(sleepLen)
  }
}

func doGeneration(gridSize int) {
  grid := util.MakeGrid(gridSize)
  util.PrintGrid(grid)
}