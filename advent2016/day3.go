package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func checkSides(s []string) bool {
  sides := make([]int,0)
  for _, side := range(s) {
    conv, err := strconv.Atoi(side)
    check(err)
    sides = append(sides, conv)
  }
  if len(sides) < 3 { // for that pesky last line
    return false
  }
  // shall just brute force this since it's only permutations of 3
  if sides[0] + sides[1] <= sides[2] {
    return false
  }
  if sides[0] + sides[2] <= sides[1] {
    return false
  }
  if sides[1] + sides[0] <= sides[2] {
    return false
  }
  if sides[1] + sides[2] <= sides[0] {
    return false
  }
  if sides[2] + sides[0] <= sides[1] {
    return false
  }
  if sides[2] + sides[1] <= sides[0] {
    return false
  }
  return true
}

func check3b(a, b, c string) bool {
  x, err := strconv.Atoi(a)
  check(err)
  y, err := strconv.Atoi(b)
  check(err)
  z, err := strconv.Atoi(c)
  check(err)
  if x + y <= z {
    return false
  }
  if x + z <= y {
    return false
  }
  if y + z <= x {
    return false
  }
  if y + x <= z {
    return false
  }
  if z + x <= y {
    return false
  }
  if z + y <= x {
    return false
  }
  return true
}

func Day3a() {
  raw, err := ioutil.ReadFile("input/3.txt")
  check(err)
  data := strings.Split(string(raw), "\n")
  ans := 0
  for _, triangle := range(data) {
    sides := strings.Fields(triangle)
    if checkSides(sides) {
      ans += 1
    }
  }
  fmt.Println(ans)
}

func Day3b() {
  raw, err := ioutil.ReadFile("input/3.txt")
  check(err)
  data := strings.Split(string(raw), "\n") // I know I'm supposed to do a matrix transposition but I'm lazy
  // checked len(data) % 3 == 0
  ans := 0
  for i:=0; i<len(data); i+=3 {
    if i + 3 >= len(data) {
      break
    }
    triangle := data[i:i+3]
    row0 := strings.Fields(triangle[0])
    row1 := strings.Fields(triangle[1])
    row2 := strings.Fields(triangle[2])
    if (check3b(row0[0], row1[0], row2[0])) {
      ans += 1
    }
    if (check3b(row0[1], row1[1], row2[1])) {
      ans += 1
    }
    if (check3b(row0[2], row1[2], row2[2])) {
      ans += 1
    }
  }
  fmt.Println(ans)
}
