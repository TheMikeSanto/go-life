package util

import (
  "fmt"
  "os"
  "os/exec"
  "strconv"
)

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