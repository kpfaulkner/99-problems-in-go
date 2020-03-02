package problems

import "fmt"


func pack( l []string) [][]string {

  lll := [][]string{}
  ll := []string{}

  lastChar := ""
  for _,c := range l {
    if c != lastChar {
      if lastChar != "" {
        lll = append(lll, ll)
        ll = []string{c}
      } else {
        ll = append(ll,c)
      }
    } else {
      ll = append(ll,c)
    }
    lastChar = c
  }
  return lll
}

func pack2( l []string) [][]string {

  lll := [][]string{}
  ll := []string{}

  lastChar := l[0]
  ll = append(ll, lastChar)
  for i:=1; i< len(l); i++ {
    if l[i] != lastChar {
      lll = append(lll, ll)
      ll = []string{}
    }
    ll = append(ll,l[i])
    lastChar = l[i]
  }
  lll = append(lll, ll)
  return lll
}

func pack3( l []string) [][]string {

  lll := [][]string{}
  ll := []string{}

  lastChar := l[0]
  ll = append(ll, lastChar)
  for _,c := range l[1:] {
    if c != lastChar {
      lll = append(lll, ll)
      ll = []string{}
    }
    ll = append(ll,c)
    lastChar = c
  }
  lll = append(lll, ll)
  return lll
}

func Prob9() {
  fmt.Printf("prob9\n")

  l := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}

  fmt.Printf("pack %v\n", pack3(l))
}

