package main

import (
  "fmt"
  "math/rand"
  // "time"
)

func main() {
  // messages := make(chan string)
  // go func() { messages <- "ping" }()
  // msg := <- messages
  // fmt.Println(msg)
  rows, columns := 256, 256
  matrix := initMatrix(rows, columns)
  fmt.Println(matrix)
}

func initMatrix(rows int, columns int) [][]bool {
  matrix := make([][]bool, rows)

  for i := 0; i < rows; i++ {
    matrix[i] = make([]bool, columns)
    for j := 0; j < columns; j++ {
      matrix[i][j] = rand.Intn(2) == 0
    }
  }

  return matrix
}