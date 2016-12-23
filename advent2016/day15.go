package advent2016

import (
  "fmt"
)

func gcd(a, b int) int {
  if b == 0 {
    return a
  } else {
    return gcd(b, a % b)
  }
}

func lcm(a, b int) int {
  return (a * b) / gcd(a, b)
}


func Day15() {
  disk := []int{17, 19, 7, 13, 5, 3, 11}
  pos := []int{5, 8, 1, 7, 1, 0, 0}
  // gcd disk = 1; ie disk is pairwise coprime
  /*running := gcd(disk[0], disk[1])

  for i := 2; i<len(disk); i++ {
    running = gcd(running, disk[i])
  }

  fmt.Println(running)*/
  t := 0
  for true {
    mask := 0
    for i:=0; i<len(disk); i++ {
      if (t + i + pos[i]) % disk[i] == 0 {
        mask = mask | (1 << uint32(i))
      }
    }
    if mask == 127 {
      fmt.Println(t-1)
      break
    }
    t++
  }
}
