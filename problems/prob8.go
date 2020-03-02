package problems

import "fmt"


func compress( l []string) []string {

  ll := []string{}
  lastChar := ""
  for _,c := range l {
    if c != lastChar {
      ll = append(ll,c)
      lastChar = c
    }
  }
  return ll
}

func prob8() {
  fmt.Printf("prob8\n")

  l := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}

  fmt.Printf("compress %v\n", compress(l))
}

