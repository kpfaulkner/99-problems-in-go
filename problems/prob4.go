package problems

import "fmt"

func length(l []int) int {
  return len(l)
}

func prob4() {
  fmt.Printf("prob4\n")

  l := []int{1, 1, 2, 3, 5, 8}
  fmt.Printf("length is %d\n", length(l))
}

