package main

import "fmt"

func main()  {
	name:="Ismail"
	age:=28
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

	//Print
	fmt.Print("Hello,")
	fmt.Print("Go")
	fmt.Print("\n")
	fmt.Print("This is new Line \n")

	//PrintLn
	fmt.Println("My name is :",name,"and I am ",age, "years old")
	//Printf

	fmt.Printf("My name is : %v and I am %v years old\n",name,age)

	//Sprintf

	var str = fmt.Sprintf("My name is : %v and I am %v years old",name,age)
	fmt.Println("The save string is: ",str)



}