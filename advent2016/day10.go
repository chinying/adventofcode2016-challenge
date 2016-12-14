package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type Bot struct {
  Id int
  Out_Lo int
  Out_Hi int
  Chips []int
}

func (b *Bot) SetLo(lo int) { b.Out_Lo = lo }
func (b *Bot) SetHi(hi int) { b.Out_Hi = hi }
func (b *Bot) AddChip(x int) { b.Chips = append(b.Chips, x) }
func (b *Bot) ClearChips() {
  newSlice := make([]int, 0)
  b.Chips = newSlice
}

type Pair struct {
  First int
  Second int
}

func Day10() {
  raw, err := ioutil.ReadFile("input/10.txt")
  check(err)
  lines := strings.Split(string(raw), "\n")

  // check notes.md on how this value was obtained
  var maxN int = 210
  g := make([]Bot, maxN) // this map indexed by bot ID

  for i:=0; i<maxN; i++ {
    chips := make([]int, 0)
    g[i] = Bot{i, -1, -1, chips}
  }

  valCount := 0

  for _, line := range(lines) {
    if line == "" {
      continue
    }
    tokens := strings.Fields(line)
    if tokens[0] == "bot" {
      // check if it is giving to a bot or output bin
      from, err := strconv.Atoi(tokens[1])
      check(err)

      toLo, err := strconv.Atoi(tokens[6])
      check(err)
      if tokens[5] == "bot" {
        g[from].SetLo(toLo)
      } else if tokens[5] == "output" { // assign to -2
        g[from].SetLo(-(2+toLo))
      }

      toHi, err := strconv.Atoi(tokens[11])
      check(err)
      if tokens[10] == "bot" {
        g[from].SetHi(toHi)
      } else if tokens[10] == "output" {
        g[from].SetHi(-(2+toHi))
      }
    } else if tokens[0] == "value" {
      // TODO
      val, err := strconv.Atoi(tokens[1])
      check(err)
      bot, err := strconv.Atoi(tokens[5])
      check(err)
      // assign val to bot
      g[bot].AddChip(val)
      valCount++
    }
  }

  activeBots := make(map[int]bool)

  for i:=0; i<maxN; i++ {
    if len(g[i].Chips) == 2 {
      activeBots[i] = true
    }
  }

  endBin := make(map[int]int)
  //alternative := -1 // may need > 1?
  for len(endBin) < valCount {
    for k := range activeBots {
      current := g[k]
      lo, hi := current.Chips[0], current.Chips[1]
      if lo > hi {
        hi ^= lo
        lo ^= hi
        hi ^= lo
      }

      if lo == 17 && hi == 61 {
        fmt.Println(current.Id) // *** 10a ans here ***
      }

      toLo := current.Out_Lo
      toHi := current.Out_Hi
      if toLo < -1 {
        endBin[toLo] = lo
      } else {
        g[toLo].AddChip(lo)
      }

      if toHi < -1 {
        endBin[toHi] = hi
      } else {
        g[toHi].AddChip(hi)
      }

      current.ClearChips()
      delete(activeBots, current.Id)

      if toHi > -1 && len(g[toHi].Chips) == 2 {
        activeBots[toHi] = true
      }

      if toLo > -1 && len(g[toLo].Chips) == 2 {
        activeBots[toLo] = true
      }

    }
  }
  fmt.Printf("10b : %d\n", endBin[-2] * endBin[-3] * endBin[-4])
}
