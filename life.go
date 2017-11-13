package main

import (
  "./grid"
  "./shell"
  "time"
)

func main() {
  size := shell.GetNumLines() - 4
  startGrid := grid.MakeGrid(size)
  shell.ClearScreen()
  DoGeneration(startGrid)
}

func DoGeneration(aGrid [][]bool) {
  shell.PrintGrid(aGrid)
  newGrid := grid.DetermineNextGen(aGrid)
  time.Sleep(1 * time.Second)
  DoGeneration(newGrid)
}