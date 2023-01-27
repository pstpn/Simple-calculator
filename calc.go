package main

import (
	"errors"
	"fmt"
	"my_project/calc/my_packages/operations"
	"os"
)

// Главная функция программы проекта
func main() {

	if len(os.Args) != 2 {
		fmt.Println(errors.New("error: incorrect args"))
		return
	}

	expression := os.Args[1]

	postfixExpr, err := operations.GetPostfixExpression(expression)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	answer, err := operations.CalcPostfixExpr(postfixExpr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Expression value: ", answer)
}
