package problems

import "fmt"

func nth(l []int, index int) int {
  return l[index]
}

func prob3() {
  fmt.Printf("prob3\n")

  l := []int{1, 1, 2, 3, 5, 8}
  fmt.Printf("nth is %d\n", nth(l,2))
}

