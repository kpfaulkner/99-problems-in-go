package problems

import "fmt"

func last(l []int) int {
  return l[len(l)-1]
}

func prob1() {
  fmt.Printf("prob1\n")

  l := []int{1, 1, 2, 3, 5, 8}
  fmt.Printf("last is %d\n", last(l))
}


