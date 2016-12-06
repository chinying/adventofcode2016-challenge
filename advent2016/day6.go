package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func Day6a() {
  raw, err := ioutil.ReadFile("input/6.txt")
  check(err)
  lines := strings.Split(string(raw), "\n")
  lines = lines[:len(lines)-1] // strip last blank line

  ans := make([]rune, len(lines[0]))

  for j:=0; j<len(lines[0]); j++ {
    freq := make(map[rune]int)
    for i:=0; i<len(lines); i++ {
      c := rune(lines[i][j])
      x, ok := freq[c]
      if !ok {
        freq[c] = 1
      } else {
        freq[c] = x + 1
      }
    }
    fm := sortByFreq(freq)
    ans[j] = fm[0].Key
  }
  fmt.Println(string(ans))
}

func Day6b() {
  raw, err := ioutil.ReadFile("input/6.txt")
  check(err)
  lines := strings.Split(string(raw), "\n")
  lines = lines[:len(lines)-1] // strip last blank line

  ans := make([]rune, len(lines[0]))

  for j:=0; j<len(lines[0]); j++ {
    freq := make(map[rune]int)
    for i:=0; i<len(lines); i++ {
      c := rune(lines[i][j])
      x, ok := freq[c]
      if !ok {
        freq[c] = 1
      } else {
        freq[c] = x + 1
      }
    }
    fm := sortByFreq(freq)
    ans[j] = fm[len(fm)-1].Key
  }
  fmt.Println(string(ans))
}
