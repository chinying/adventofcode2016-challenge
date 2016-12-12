package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type RepeatCluster struct {
  Text []byte
  Count int
}

func (r *RepeatCluster) Length() int { return len(r.Text)}
func expandRepeatCluster(r RepeatCluster) int {
  cnt := 0
  for i:=0; i<len(r.Text); i++ {
    if r.Text[i] == '(' {
      i++
      startMarker := i
      for r.Text[i] != ')' {
        i++
      }

      xy := strings.Split(string(r.Text[startMarker:i]), "x")
      x, _ := strconv.Atoi(xy[0])
      rep, _ := strconv.Atoi(xy[1])
      r1 := RepeatCluster{r.Text[i+1:i+x+1], rep}
      cnt += expandRepeatCluster(r1) * rep
      i += x
    } else {
      cnt ++
    }
  }
  return cnt
}

func Day9a() {
  raw, err := ioutil.ReadFile("input/9.txt")
  check(err)
  raw = []byte(strings.TrimSpace(string(raw)))
  cnt := 0
  for i:=0; i<len(raw); i++ {
    if raw[i] == '(' {
      i++
      startMarker := i
      for raw[i] != ')' {
        i++
      }
      xy := strings.Split(string(raw[startMarker:i]), "x")
      x, _ := strconv.Atoi(xy[0])
      rep, _ := strconv.Atoi(xy[1])
      i += x
      cnt += (x * rep)
    } else {
      cnt++
    }
  }
  fmt.Println(cnt)
}

func Day9b() {
  raw, err := ioutil.ReadFile("input/9.txt")
  check(err)
  raw = []byte(strings.TrimSpace(string(raw)))
  fmt.Println(expandRepeatCluster(RepeatCluster{raw, 1}))
}
