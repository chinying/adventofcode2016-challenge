package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
  "regexp"
)

func Day7a() { // must have ABBA in seq and not have ABBA within square brackets
  raw, err := ioutil.ReadFile("input/7.txt")
  check(err)
  lines := strings.Split(string(raw), "\n")
  ans := 0
  reg, err := regexp.Compile("\\[(.*?)\\]")
  check(err)
  for _, lin := range lines {
    flag := false
    _lines := reg.ReplaceAllString(lin, "-") // there are many ways this can go wrong
    brackets := reg.FindAllString(lin, -1)
    //fmt.Println(lin, line)
    _line := strings.Split(_lines, "-")
    for _, line := range _line {
      for i:=0; i<len(line); i++ {
        if i + 3 < len(line) {
          if line[i+3] == line[i] && line[i+1] == line[i+2] && line[i] != line[i+1] {
            flag = true
          }
        }
      }
    }

    for _, s := range brackets {
      if !flag { break }
      for i := range s {
        if !flag { break }
        if i + 3 < len(s) {
          if s[i+3] == s[i] && s[i+1] == s[i+2] && s[i] != s[i+1] {
            flag = false
          }
        }
      }
    }


    if flag {
      ans += 1
    }
  }
  fmt.Println(ans)
}

func check_aba_match(aba, bab []string) bool {
  set := make(map[string]bool)
  for _, k := range aba {
    set[string(k)] = true
  }
  for _, b := range bab {
    testString := make([]byte, 3)
    testString[0], testString[1], testString[2] = b[1], b[0], b[1]
    _, ok := set[string(testString)]
    if ok {
      return true
    }
  }
  return false
}

func Day7b() {
  raw, err := ioutil.ReadFile("input/7.txt")
  check(err)
  lines := strings.Split(string(raw), "\n")
  ans := 0
  reg, err := regexp.Compile("\\[(.*?)\\]")
  check(err)
  for _, lin := range lines {
    flag, flag2 := false, false
    _lines := reg.ReplaceAllString(lin, "-") // there are many ways this can go wrong
    brackets := reg.FindAllString(lin, -1)
    _line := strings.Split(_lines, "-")
    aba := make([]string, 0)
    for _, line := range _line {
      for i:=0; i<len(line); i++ {
        if i + 2 < len(line) {
          if line[i+2] == line[i] && line[i] != line[i+1] {
            flag = true
            aba = append(aba, line[i:i+3])
          }
        }
      }
    }

    bab := make([]string, 0)
    for _, s := range brackets {
      for i := range s {
        if i + 2 < len(s) {
          if s[i+2] == s[i] && s[i] != s[i+1] {
            bab = append(bab, s[i:i+3])
          }
        }
      }
    }

    flag2 = check_aba_match(aba, bab)

    if flag && flag2 {
      ans += 1
    }
  }
  fmt.Println(ans)
}
