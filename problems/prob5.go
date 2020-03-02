package problems

import "fmt"

func reverse(l []int) []int {
  ll := make([]int,len(l), len(l))
  for i,_ := range l {
    newIndex := len(l) - 1 - i
    ll[newIndex] = l[i]
  }
  return ll
}

func reverse2(l []int) []int {
  ll := make([]int,len(l), len(l))
  for i,j :=0, len(l)-1 ; i < len(l); i,j = i+1, j-1 {
    ll[j] = l[i]
  }
  return ll
}


func prob5() {
  fmt.Printf("prob5\n")

  l := []int{1, 1, 2, 3, 5, 8}
  fmt.Printf("rev is %v\n", reverse(l))
  fmt.Printf("rev2 is %v\n", reverse2(l))
}

