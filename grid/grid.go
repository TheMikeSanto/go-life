package grid

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

const aliveCell string = "\033[40m  "
const deadCell  string = "\033[47m  "
const moveCursor string = "\033[0;0H"
const termReset  string = "\033[0m"
const termSpacer string = "\033[49m "
func PrintGrid(grid [][]bool) {
  var buffer bytes.Buffer
  buffer.WriteString(moveCursor + "\n")
  for _, row := range grid {
    buffer.WriteString(termReset)
    for i := 0; i < 2; i++ {
      buffer.WriteString(termSpacer)
    }
    for _, cell := range row {
      if cell {
        buffer.WriteString(aliveCell)
      } else {
        buffer.WriteString(deadCell)
      }
    }
    buffer.WriteString("\n")
  }

  buffer.WriteString(termReset + "\n")
  fmt.Printf(buffer.String())
}