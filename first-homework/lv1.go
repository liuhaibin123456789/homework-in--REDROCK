package main

import (
	"fmt"
)

func countTwoNumber(number1, number2 float64, operationSymbol string) float64 {
	if operationSymbol!="+"&&operationSymbol!="-"&&operationSymbol!="*"&&operationSymbol!="/" {
		return 0
	}
	switch operationSymbol {
	case "+":
		return number1+number2
	case "-":
		return number1-number2
	case "*":
		return number1*number2
	case "/":
		return number1/number2
	default:
	}
	return 1
}
func main()  {
	var (
		number1,number2 float64
		operationSymbol string
		isContinue bool=true
	)
	for isContinue {
		fmt.Println("enter a number:")
		fmt.Scan(&number1)
		fmt.Println("enter another number:")
		fmt.Scan(&number2)
		fmt.Println("enter a way to count (+-*/):")
		fmt.Scan(&operationSymbol)
		fmt.Printf("result is:%.2f\n",countTwoNumber(number1,number2,operationSymbol))
		fmt.Println("isContinued? (true/false)")
		fmt.Scan(&isContinue)
	}
}

