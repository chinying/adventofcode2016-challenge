package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func Day8() {
  raw, err := ioutil.ReadFile("input/8.txt")
  check(err)
  lines := strings.Split(string(raw), "\n")
  lines = lines[:len(lines)-1]
  wide, tall := 50, 6
  screen := make([][]bool, tall)
  for i:=0; i<tall; i++ { // i = y
    screen[i] = make([]bool, wide)
    for j:=0; j<wide; j++ { // j = x
      screen[i][j] = false // false means off
    }
  }

  for _, line := range(lines) {
    tokens := strings.Fields(line)
    action := tokens[0]
    switch action {
      case "rect":
        xy := strings.Split(tokens[1], "x")
        x, err := strconv.Atoi(xy[0])
        check(err)
        y, err := strconv.Atoi(xy[1])
        check(err)
        for i:=0; i<y; i++ {
          for j:=0; j<x; j++ {
            screen[i][j] = true
          }
        }
      case "rotate":
        //fmt.Println("rotate", tokens[2], tokens[4])
        // tokens[2] is start index
        // tokens[4] is number of blocks to translate by, need to wrap around
        idx, ok := strconv.Atoi(tokens[2][2:])
        check(ok)
        shift, ok := strconv.Atoi(tokens[4])
        check(ok)
        if tokens[2][0] == 'x' {
          newLine := make([]bool, tall)
          for i:=0; i<shift; i++ {
            newLine[i] = screen[tall-shift+i][idx]
          }
          j := 0
          for i:=shift; i<tall; i++ {
            newLine[i] = screen[j][idx]
            j++
          }
          for i:=0; i<tall; i++ {
            screen[i][idx] = newLine[i]
          }
        } else if tokens[2][0] == 'y' {
          newLine := make([]bool, wide)
          for i:=0; i<shift; i++ {
            newLine[i] = screen[idx][wide-shift+i]
          }
          j := 0
          for i:=shift; i<wide; i++ {
            newLine[i] = screen[idx][j]
            j++
          }
          for i:=0; i<wide; i++ {
            screen[idx][i] = newLine[i]
          }
        }
    }
  }
  ans := 0
  for i:=0; i<tall; i++ { // i = y
    for j:=0; j<wide; j++ { // j = x
      if screen[i][j] {
        ans += 1
        fmt.Print("*")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
  fmt.Println(ans)
}
