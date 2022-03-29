package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		Using 'if' statements, determine the following and print out the answer
		for each query.
		(1) Given the numbers are 17 and 24 and
		(2) Using preceeding conditionals for at least 1 of the numbers

		(a) which number is bigger?
		(b) by how much is it bigger by?
	*/

	// number1 := 17
	// number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

	if number1, number2 := 17, 24; number1 > number2 {
		bigger := number1 - number2
		resultMessage = strconv.Itoa(number1)
		resultMessage += " is bigger than "
		resultMessage += strconv.Itoa(number2)
		resultMessage += " by"
		resultMessage += strconv.Itoa(bigger)
		fmt.Println(resultMessage)
	}

	if number1, number2 := 17, 24; number1 < number2 {
		bigger := number2 - number1
		resultMessage = strconv.Itoa(number2)
		resultMessage += " is bigger than "
		resultMessage += strconv.Itoa(number1)
		resultMessage += " by "
		resultMessage += strconv.Itoa(bigger)
		fmt.Println(resultMessage)
	}

}
