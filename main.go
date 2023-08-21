package main

import "fmt"

func main()  {
	fmt.Println("Hello, Go")
	//string
	var varOne string = "this is first variable"
	var varTwo = "Second Variable"
	varThree := "Third Variable"
	var varFour string
	fmt.Println(varOne, varTwo, varThree, varFour)
	varFour = "Four updated"
	fmt.Println(varOne, varTwo, varThree, varFour)
	
	//int
	var intOne int = 20
	var intTwo int8 = 100
	var intTree int16 = 1000
	var intFoure float32 = 5000.34
	var intFive  float64 = 59684459034.0968
	int60 := 456
	var intSeven uint = 245
	fmt.Println(intOne, intTwo, intTree, intFoure, intFive, int60, intSeven)
}