package util

import (
  "bytes"
  "fmt"
  "math/rand"
  "time"
)
func MakeGrid(size int) [][]bool {
  rand.Seed(time.Now().UTC().UnixNano())
  grid := make([][]bool, size)

  for i := 0; i < size; i++ {
    grid[i] = make([]bool, size)
    for j := 0; j < size; j++ {
      grid[i][j] = rand.Intn(2) == 0
    }
  }

  return grid
}

func PrintGrid(grid [][]bool) {
  var buffer bytes.Buffer
  buffer.WriteString("\033[0;0H\n")
  for _, row := range grid {
    buffer.WriteString("\t")
    for _, cell := range row {
      if cell {
        buffer.WriteString("\033[40m  ")
      } else {
        buffer.WriteString("\033[47m  ")
      }
    }
    buffer.WriteString("\n")
  }

  buffer.WriteString("\033[0m\n")
  fmt.Printf(buffer.String())
}