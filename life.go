package main

import (
  "./grid"
  "./shell"
  "time"
)

func main() {
  gridSize := shell.GetNumLines() - 4
  shell.ClearScreen()
  startGrid := grid.MakeGrid(gridSize)
  DoGeneration(startGrid)
}

func DoGeneration(aGrid [][]bool) {
  shell.PrintGrid(aGrid)
  newGrid := grid.DetermineNextGen(aGrid)
  time.Sleep(250 * time.Millisecond)
  DoGeneration(newGrid)
}