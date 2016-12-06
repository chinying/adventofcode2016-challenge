package advent2016

import (
  "fmt"
  "crypto/md5"
  "strings"
  "strconv"
  "io"
  "encoding/hex"
)

func Day5a() {
  roomid := "ffykfhsq"
  l := 0
  var i int64 = 0
  for l < 8 {
    h := md5.New()
    io.WriteString(h, roomid + strconv.FormatInt(i, 10))
    hash := hex.EncodeToString(h.Sum(nil))
    if (strings.HasPrefix(hash, "00000")) {
      fmt.Println(hash)
      l += 1
    }
    i += 1
  }
}

func Day5b() {
  roomid := "ffykfhsq"
  var i int64 = 0
  pos := make(map[byte]byte)
  for len(pos) < 8 {
    h := md5.New()
    io.WriteString(h, roomid + strconv.FormatInt(i, 10))
    hash := hex.EncodeToString(h.Sum(nil))
    if (strings.HasPrefix(hash, "00000")) {
      if (hash[5] > 47 && hash[5] < 56) {
        position, val := hash[5], hash[6]
        _, ok := pos[position]
        if !ok {
          pos[position] = val
        }
      }
    }
    i += 1
  }
  fmt.Println(pos)
  ans := make([]byte, 8)
  cnt := 0
  for j:=48; j<56; j++ { // 0 to 7
    ans[cnt] = pos[byte(j)]
    cnt += 1
  }
  fmt.Println(string(ans))
}
