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
  size := 20
  sleepLen := 2 * time.Second
  util.ClearScreen()
  doGeneration(size)
  time.Sleep(sleepLen)
  doGeneration(size)
  time.Sleep(sleepLen)
  doGeneration(size)
  time.Sleep(sleepLen)
  doGeneration(size)
}

func doGeneration(gridSize int) {
  grid := util.MakeGrid(gridSize)
  util.PrintGrid(grid)
}