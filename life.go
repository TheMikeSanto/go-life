package main

import (
  "./grid"
  "./shell"
  "time"
)

func main() {
  // messages := make(chan string)
  // go func() { messages <- "ping" }()
  // msg := <- messages
  // fmt.Println(msg)
  size := shell.GetNumLines() - 4
  sleepLen := 1 * time.Second
  shell.ClearScreen()
  for true {
    doGeneration(size)
    time.Sleep(sleepLen)
  }
}

func doGeneration(gridSize int) {
  newGrid := grid.MakeGrid(gridSize)
  grid.PrintGrid(newGrid)
}