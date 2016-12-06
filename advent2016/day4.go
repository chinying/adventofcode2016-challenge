package advent2016

import (
  "fmt"
  "io/ioutil"
  "strings"
  "sort"
  "strconv"
)

type FreqPair struct {
  Key rune
  Freq int
}

type FreqMap []FreqPair
func (f FreqMap) Len() int {return len(f)}
func (f FreqMap) Less(i, j int) bool {
  if f[i].Freq == f[j].Freq {
    return f[i].Key > f[j].Key // opp. sign because of sort.Reverse
  }
  return f[i].Freq < f[j].Freq
}
func (f FreqMap) Swap(i, j int) { f[i], f[j] = f[j], f[i]}

func sortByFreq(m map[rune]int) FreqMap {
  fm := make(FreqMap, len(m))
  i := 0
  for k, v := range m {
    fm[i] = FreqPair{k, v}
    i += 1
  }
  sort.Sort(sort.Reverse(fm))
  return fm
}

func genCheckSum(s string) (string, string) {
  sectorSep := strings.LastIndex(s, "-")
  sector := s[sectorSep+1:]
  freq := make(map[rune]int)
  for _, c := range s[:sectorSep] {
    i, ok := freq[c]
    if !ok {
      freq[c] = 1
    } else {
      freq[c] = i + 1
    }
  }
  delete(freq, '-')
  fm := sortByFreq(freq)
  five := make([]rune, 5)
  for i:=0; i<5; i++ {
    five[i] = fm[i].Key
  }
  return string(five), sector
}

func Day4a() {
  raw, err := ioutil.ReadFile("input/4.txt")
  check(err)
  data := strings.Split(string(raw), "\n")
  ans := 0
  for _, d := range data[:len(data)-1] {
    checksum, sector := genCheckSum(d[:len(d)-7])
    actualChecksum := d[len(d)-7:][1:6]
    if actualChecksum == checksum {
      sectorInt, ok := strconv.Atoi(sector)
      check(ok)
      ans += sectorInt
    }
  }
  fmt.Println(ans)
}

func crackCaesarCipher(cipher string, mod int) string {
  // split by dash, join by space at end
  message := make([]rune, 0)
  words := strings.Split(cipher, "-")
  for _, word := range words {
    for _, c := range word {
      pos := (int(c) + (mod % 26))
      if pos > 122 {
        pos -= 26
      }
      message = append(message, rune(pos))
    }
    message = append(message, ' ')
  }
  return (string(message))
}

func Day4b() {
  raw, err := ioutil.ReadFile("input/4.txt")
  check(err)
  data := strings.Split(string(raw), "\n")
  for _, d := range data {
    if len(d) == 0 {
      break
    }
    sectorSep := strings.LastIndex(d[:len(d)-7], "-")
    sector, ok := strconv.Atoi(d[sectorSep+1:len(d)-7])
    check(ok)
    fmt.Println(crackCaesarCipher(d[:sectorSep], sector), sector)
  }
}
