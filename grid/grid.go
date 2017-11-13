package grid

import (
  "fmt"
  "math/rand"
  "time"
)

func DetermineNextGen(grid [][]bool) [][]bool {
  // For each row r, send {r-1, r, r+1 }
  // If first row, r-1 = last row in grid
  // If last row, r+1 = first row in grid
  nextGrid := make([][]bool, len(grid))
  numRows := len(grid)

  for index, row := range grid {
    var nextRow, prevRow []bool

    if index == (numRows - 1) {
      nextRow = grid[0]
    } else {
      nextRow = grid[index + 1]
    }

    if index == 0 {
      prevRow = grid[numRows - 1]
    } else {
      prevRow = grid[index - 1]
    }

    nextGrid[index] = DetermineRow(row, nextRow, prevRow)
  }

  return nextGrid
}

func DetermineRow(row []bool, nextRow []bool, prevRow []bool) []bool {
  fmt.Sprintf(string(len(row)))
  fmt.Sprintf(string(len(nextRow)))
  fmt.Sprintf(string(len(prevRow)))
  newRow := make([]bool, len(row))
  for index, cell := range row {
    numAlive := 0
    neighbors := GetNeighbors(index, row, nextRow, prevRow)
    for _, neighbor := range neighbors {
      if neighbor {
        numAlive += 1
      }
    }
    // For a living cell
    if cell && (numAlive > 3 || numAlive < 2) {
      newRow[index] = false
    }
    if cell && (numAlive == 2 || numAlive == 3) {
      newRow[index] = cell
    }
    // For a dead cell
    if !cell && (numAlive == 3) {
      newRow[index] = true
    }
  }
  return newRow
}

func GetNeighbors(index int, row []bool, nextRow []bool, prevRow []bool) [8]bool {
  var nw,n,ne,w,e,sw,s,se bool
  numCells := len(row)

  n = prevRow[index]
  s = nextRow[index]

  if index == 0 {
    nw = prevRow[numCells - 1]
    w = row[numCells - 1]
    sw = nextRow[numCells -1]
  } else {
    nw = prevRow[index - 1]
    w = row[index - 1]
    sw = nextRow[index - 1]
  }

  if index == (numCells - 1) {
    ne = prevRow[0]
    e = row[0]
    se = nextRow[0]
  } else {
    ne = prevRow[index + 1]
    e = row[index + 1]
    se = nextRow[index + 1]
  }

  neighbors := [8]bool{nw, n, ne, w, e, sw, s, se}
  return neighbors
}

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
