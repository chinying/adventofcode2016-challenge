package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func Day12() {
  raw, err := ioutil.ReadFile("input/12.txt")
  check(err)
  lines := strings.Split(string(raw), "\n")
  lines = lines[:len(lines)-1]
  registers := make(map[string]int)
  for i:='a'; i<='d'; i++ {
    registers[string(i)] = 0
  }

  registers["c"] = 1 // *** comment out for part A ***

  for i:=0; i<len(lines); i++ {
    tokens := strings.Fields(lines[i])
    action := tokens[0]
    switch action {
      case "cpy":
        val, reg := tokens[1], tokens[2]
        if v, err := strconv.Atoi(val); err == nil {
          registers[reg] = v
        } else {
          v, _ := registers[val]
          registers[reg] = v
        }
      case "jnz":
        x, y := tokens[1], tokens[2]
        if v, err := strconv.Atoi(x); err == nil { // x is a number
          if v != 0 {
            jmp, _ := strconv.Atoi(y)
            i += jmp-1 //-1 needed because of i++ at counter
          }
        } else { // x is a register, hence you need to fetch value of register
          val, _ := registers[x]
          if val != 0 {
            jmp, _ := strconv.Atoi(y)
            i += jmp-1
          }
        }
      case "dec":
        reg := tokens[1]
        val, _ := registers[reg]
        registers[reg] = val - 1
      case "inc":
        reg := tokens[1]
        val, _ := registers[reg]
        registers[reg] = val + 1
    }
  }
  ans, _ := registers["a"]
  fmt.Println(ans)
}
