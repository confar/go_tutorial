package main

import ("fmt"
        )

func add(x, y float64) float64 {
  return x + y
}

func multiple(a, b string) (string,string){
  return a, b


}

func main() {
  // fmt.Println("The aquare root of 16 is", math.Sqrt(15))
  // fmt.Println("The random number between 1 and 100", rand.Intn(1000))
  // var num1,num2 float64 = 5.5, 6.8
  // var num2 float64 = 6.7
  // fmt.Println(add(num1,num2))
  w1, w2 := "Hey", "mate"
  fmt.Println(multiple(w1,w2))
}
