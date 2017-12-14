package main

import ("fmt"
        )


func main() {
  x := 15
  a := &x //memory adress
  fmt.Println(a, *a)
  *a = 5 // это короч дикий указатель такой на объект, который находится в ячейке памяти. шик в том, что
  x = 16
  fmt.Println(a, x)
  *a = *a**a
  fmt.Println(a, x)
  }
