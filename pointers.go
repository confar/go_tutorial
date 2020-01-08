package main

import ("fmt"
        )


func main() {
  x := 15
  a := &x //memory adress
  fmt.Println(a, *a)
  *a = 5 // указатель на объект в ячейке памяти
  x = 16
  fmt.Println(a, x)
  *a = *a**a
  fmt.Println(a, x)
  }
