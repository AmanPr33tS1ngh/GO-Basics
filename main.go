package main

import (
	"fmt"
	"unicode/utf8"
)

// to run the file : go run main.go

func main(){
	// Signed Integers
	var intNum int
	// var intNum int8
	// var intNum int16
	var intNum32 int32
	// var intNum int64

	fmt.Println(intNum)

	// Unsigned Integers
	var uIntNum uint
	// var uIntNum uint8
	// var uIntNum uint16
	// var uIntNum uint32
	// var uIntNum uint64
	fmt.Println(uIntNum)

	// Float
	var floatNum float32
	// var floatNum float16
	var floatNum32 float32
	// var floatNum float64
	fmt.Println(floatNum)
	fmt.Println(floatNum32)
	
	var result = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var myStr string = "Hello World"
	fmt.Println(len(myStr)) // doesn't prints length of string but the byte value of the string

	fmt.Println(utf8.RuneCountInString(myStr)) // this actually returns length of string, yeah weird

	var myBool bool
	fmt.Println(myBool)

	// two ways of declaring variables

	
	// var myVar int = 5 // it is better to declare variables like this
	myVar := 5
	fmt.Println(myVar)

	// declaring multiple variables in one line
	myStr, myInt := "STRING", 5
	fmt.Println(myStr, myInt)

	const myConst string = "CONSTANT"
	fmt.Println(myConst)

	printFun()

	// dividedVal := division(5, 2)
	// fmt.Println(dividedVal)

	var dividedVal int = division(5, 2)
	fmt.Println(dividedVal)

	dVal, rVal, err := divisionResultAndRemainder(5, 2)
	if err!=nil{
		err.Error()
	}else if(rVal == 0){
		fmt.Println("remainder == 0")
	}
	fmt.Println(dVal, rVal)

	var d1Val, r1Val, err1 = divisionResultAndRemainder(5, 0)
	if err1!=nil{
		err1.Error()
	}else if(rVal == 0){
		fmt.Println("remainder == 0")
	}else{
		fmt.Println("the d1Val is: %v and r1Val is : %v", d1Val, r1Val) // printf similar to c
	}


	switch {
		case err1!=nil:
			err1.Error()
		case rVal == 0:
			fmt.Println("remainder == 0")
		default:
			fmt.Println("the d1Val is: %v and r1Val is : %v", d1Val, r1Val) // printf similar to c
	}


	switch rVal{
		case 0:
			fmt.Println("remainder == 0")
		case 1, 2:
			fmt.Println("remainder 1 or 2")
		default:
			fmt.Println("default") // printf similar to c
	}


}
func printFun(){
	fmt.Println("printing....")
}
func division(numerator int, denominator int) int{
	return numerator /  denominator
}

func divisionResultAndRemainder(numerator int, denominator int) (int, int, error){
	var err error
	if denominator==0 {
		fmt.Println("Cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}