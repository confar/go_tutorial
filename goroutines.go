package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
}

func say (s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 2{
			panic("Oh dear")
		}
	}
}

func main() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say( "there")
	//wg.Wait()
	
	//wg.Add(1)
	//go say( "hi")
	wg.Wait()
	//foo()
}

//func foo()  {
//	defer fmt.Println("Done")
//	defer fmt.Println("Are we done?")
//	fmt.Println("SOme stuff here")
//
//}