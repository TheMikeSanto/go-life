package shell

import (
  "bytes"
  "fmt"
  "os"
  "os/exec"
  "strconv"
)
const aliveCell string = "\033[42m  "
const deadCell  string = "\033[104m  "
const moveCursor string = "\033[0;0H"
const termReset  string = "\033[0m"
const termSpacer string = "\033[49m "

func GetNumLines() int {
  cmd := exec.Command("stty", "size")
  cmd.Stdin = os.Stdin
  out, _ := cmd.Output()
  lines, _ := strconv.Atoi(string(out)[0:2])
  return lines
}

func ClearScreen() {
  lines := GetNumLines()
  for i := 0; i < lines; i++ {
    fmt.Printf("\n")
  }
}

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
