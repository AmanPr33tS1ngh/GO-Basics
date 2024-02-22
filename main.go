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

	// if else conditionals

	if err1!=nil{
		err1.Error()
	}else if(rVal == 0){
		fmt.Println("remainder == 0")
	}else{
		fmt.Println("the d1Val is: %v and r1Val is : %v", d1Val, r1Val) // printf similar to c
	}

	//switch 

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

	// arrays

	var intArr [3]int32 // declaration of array
	fmt.Println(intArr[0])
	fmt.Println(intArr[1])
	fmt.Println(intArr[0: 3])

	fmt.Println(&intArr[0]) // returns address of stored val
	fmt.Println(&intArr[1])

	// another way to declare array
	var intArray [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArray)

	// another way to declare array with colon equal declaration method
	intArrayWithCol := [3]int32{1,2,3}
	fmt.Println(intArrayWithCol)

	// another way to declare array with len not defining
	intArrayWithNoLen := [...]int32{1,2,3}
	fmt.Println(intArrayWithNoLen)
	
	// Slice
	var intSlice []int32 = []int32{4,5,6} // declaration of slice
	fmt.Println("after", intSlice, len(intSlice), cap(intSlice)) // returns length 3 and cap 3
	intSlice = append(intSlice, 7)
	fmt.Println("after", intSlice, len(intSlice), cap(intSlice)) // returns length 4 which is basically elements in an array and cap is new increased size of slice 

	var intSlice2 []int32 = []int32{9,10,11}
	intSlice = append(intSlice, intSlice2...) // another way of appending elements in a slice
	fmt.Println(intSlice)


	// looping in slices
	for index, value := range intSlice{
		fmt.Println(index, value)
	}

	// without index
	for _, value := range intSlice{
		fmt.Println(value)
	}

	var i int = 0 // while loop
	for i<10{
		fmt.Println(i)
		i += 1
	}

	for i=0;i<10;i++{
		fmt.Println(i)
	}


	// maps => similar to dict in python
	var myMap map[string]uint8 = make(map[string]uint8) // declaration of map
	fmt.Println(myMap)
	var myMap2 = map[string]uint8{"age": 19}
	fmt.Println(myMap2)
	fmt.Println(myMap2["age"])

	var myMap2Val, exists = myMap2["age"] // myMap2Val will be having value of myMap2["age"] item and exists is a boolean if value exists it will be true
	if exists{
		fmt.Println("value is", myMap2Val)
	}

	// looping in maps
	for age, val := range myMap2{
		fmt.Println(age, val)
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