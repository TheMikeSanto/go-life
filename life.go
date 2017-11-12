package main

import (
  "bytes"
  "fmt"
  "math/rand"
  "os"
  "os/exec"
  "strconv"
  "time"
)

func main() {
  // messages := make(chan string)
  // go func() { messages <- "ping" }()
  // msg := <- messages
  // fmt.Println(msg)
  size := 20
  sleepLen := 2 * time.Second
  clearScreen()
  doGeneration(size)
  time.Sleep(sleepLen)
  doGeneration(size)
  time.Sleep(sleepLen)
  doGeneration(size)
  time.Sleep(sleepLen)
  doGeneration(size)
}

func doGeneration(gridSize int) {
  grid := initGrid(gridSize, gridSize)
  printGrid(grid)
}

func getNumLines() int {
  cmd := exec.Command("stty", "size")
  cmd.Stdin = os.Stdin
  out, _ := cmd.Output()
  lines, _ := strconv.Atoi(string(out)[0:2])
  return lines
}

func initGrid(rows int, columns int) [][]bool {
  rand.Seed(time.Now().UTC().UnixNano())
  grid := make([][]bool, rows)

  for i := 0; i < rows; i++ {
    grid[i] = make([]bool, columns)
    for j := 0; j < columns; j++ {
      grid[i][j] = rand.Intn(2) == 0
    }
  }

  return grid
}
func clearScreen() {
  lines := getNumLines()
  for i := 0; i < lines; i++ {
    fmt.Printf("\n")
  }
}

func printGrid(grid [][]bool) {
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