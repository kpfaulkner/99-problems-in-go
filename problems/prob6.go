package problems

import "fmt"

func isPalindrome(l []int) bool {
  for i,j :=0, len(l)-1 ; i < len(l); i,j = i+1, j-1 {
    if l[j] != l[i]  {
      return false
    }
  }

  return true
}


func prob6() {
  fmt.Printf("prob6\n")

  l := []int{1, 2, 3, 2, 1}
  fmt.Printf("ispal is %v\n", isPalindrome(l))
}

