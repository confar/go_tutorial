package main

import "fmt"

func main() {
	for i:=0; i<3; i++{
		fmt.Println("you man bruh")
	}
	m := 1
	for {
		fmt.Println("You mate", m)
		m++
		if m > 15 {
			break
		}
		}

	m = 4
	p := &m
	//p = "string"
	fmt.Println(p)

	a := 23
	for x:=1; a<30 && x<7; x++ {
		fmt.Println(x, a)
		a++
	}
}
