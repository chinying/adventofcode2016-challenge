package advent2016

import (
  "fmt"
  "strings"
  "strconv"
)

var d1input = "L5, R1, R4, L5, L4, R3, R1, L1, R4, R5, L1, L3, R4, L2, L4, R2, L4, L1, R3, R1, R1, L1, R1, L5, R5, R2, L5, R2, R1, L2, L4, L4, R191, R2, R5, R1, L1, L2, R5, L2, L3, R4, L1, L1, R1, R50, L1, R1, R76, R5, R4, R2, L5, L3, L5, R2, R1, L1, R2, L3, R4, R2, L1, L1, R4, L1, L1, R185, R1, L5, L4, L5, L3, R2, R3, R1, L5, R1, L3, L2, L2, R5, L1, L1, L3, R1, R4, L2, L1, L1, L3, L4, R5, L2, R3, R5, R1, L4, R5, L3, R3, R3, R1, R1, R5, R2, L2, R5, L5, L4, R4, R3, R5, R1, L3, R1, L2, L2, R3, R4, L1, R4, L1, R4, R3, L1, L4, L1, L5, L2, R2, L1, R1, L5, L3, R4, L1, R5, L5, L5, L1, L3, R1, R5, L2, L4, L5, L1, L1, L2, R5, R5, L4, R3, L2, L1, L3, L4, L5, L5, L2, R4, R3, L5, R4, R2, R1, L5"

var bearings = []int {0, 90, 180, 270}

func abs(n int) int {
  if n < 0 {return -n}
  return n
}

func getBearing(n int) int {
  ptr := 0
  if n < 0 {
    ptr = (4 - abs(n % 4)) % 4 // this fixes the 4 - 0 = 4 index out of range
  } else {
    ptr = n % 4
  }
  return bearings[ptr]
}

func Day1a() {
  steps := strings.Split(d1input, ", ")
  var up, right int = 0, 0
  var bearing = 0

  for _, step := range(steps) {
    dir := step[:1]
    magnitude, _ := strconv.Atoi(strings.Join([]string{step[1:]}, ""))
    switch dir {
      case "L":
        bearing -= 1
      case "R":
        bearing += 1
      default:
        bearing += 0
    }
    switch getBearing(bearing) {
      case 0:
        up += magnitude
      case 90:
        right += magnitude
      case 180:
        up -= magnitude
      case 270:
        right -= magnitude
    }
  }
  fmt.Println(up, right)
}

func Day1b() { // there has to be an easier way to do this using vectors / comp geometry
  steps := strings.Split(d1input, ", ")
  grid := make([][]bool, 800)
  for i:=0; i<len(grid); i++ {
    grid[i] = make([]bool, 800)
  }
  //0,0 at grid 400,400
  x, y := 400, 400
  bearing := 0
  flag := true
  grid[x][y] = true
  for _, step := range(steps) {
    if !flag {
      break
    }
    dir := step[:1]
    magnitude, _ := strconv.Atoi(strings.Join([]string{step[1:]}, ""))
    switch dir {
      case "L":
        bearing -= 1
      case "R":
        bearing += 1
      default:
        bearing += 0
    }
    //fmt.Println(x,y)
    switch getBearing(bearing) {
      case 0:
        for j:=y+1; j<y+magnitude; j++ {
          if grid[j][x] == false {
            grid[j][x] = true
          } else {
            fmt.Printf("collision at %d %d\n", j, x)
            flag = false
          }
        }
        y += magnitude
      case 90:
        for i:=x+1; i<x+magnitude; i++ {
          if grid[y][i] == false {
            grid[y][i] = true
          } else {
            fmt.Printf("collision at %d %d\n", y, i)
            flag = false
          }
        }
        x += magnitude
      case 180:
        for j:=y-1; j>=y-magnitude; j-- {
          if grid[j][x] == false {
            grid[j][x] = true
          } else {
            fmt.Printf("collision at %d %d\n", j, x)
            flag = false
          }
        }
        y -= magnitude
      case 270:
        for i:=x-1; i>=x-magnitude; i-- {
          if grid[y][i] == false {
            grid[y][i] = true
          } else {
            fmt.Printf("collision at %d %d\n", y, i)
            flag = false
          }
        }
        x -= magnitude
    }
  }
  fmt.Println(x, y) // note this isn't answer
}

