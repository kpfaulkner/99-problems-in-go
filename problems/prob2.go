package problems

import "fmt"

func lastButOne(l []int) int {
  return l[len(l)-2]
}

func prob2() {
  fmt.Printf("prob2\n")

  l := []int{1, 1, 2, 3, 5, 8}
  fmt.Printf("last is %d\n", lastButOne(l))
}


