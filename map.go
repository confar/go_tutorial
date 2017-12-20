package main

import "fmt"

func main() {
	grades := make(map[string]float64)

	grades["Timmy"] = 42
	grades["Jessica"] = 92
	grades["Matt"] = 70

	fmt.Println(grades)

	//TimsGrade := grades["Timmy"]
	delete(grades, "Timmy")
	//fmt.Println(TimsGrade)
	fmt.Println(grades)

	for k,v := range grades {
		fmt.Println(k, v)
	}
}