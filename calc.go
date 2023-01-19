package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

type Stack struct {
	LastIndex int
	Data      []uint8
}

func PushInStack(stack *Stack, sign uint8) {
	stack.Data[stack.LastIndex] = sign
	stack.LastIndex++
}

func PopFromStack(stack *Stack) (uint8, error) {
	if stack.LastIndex == 0 {
		return 0, errors.New("error: empty stack")
	}

	stack.LastIndex--
	return stack.Data[stack.LastIndex], nil
}

func GetPostfixExpression(inExpr string) (string, error) {
	var outExpr string
	var signStack Stack

	signStack.Data = make([]uint8, 10)
	signStack.LastIndex = 0

	for i := 0; i < len(inExpr); {
		if inExpr[i] == '(' {
			PushInStack(&signStack, inExpr[i])
			i++
		} else if inExpr[i] == ')' {
			popElem, _ := PopFromStack(&signStack)

			for signStack.LastIndex > 0 && popElem != '(' {
				outExpr += string(popElem)
				popElem, _ = PopFromStack(&signStack)
			}

			if signStack.LastIndex == 0 && popElem != '(' {
				return "", errors.New("error: incorrect expression")
			}
			i++
		} else if inExpr[i] == '-' || inExpr[i] == '+' {
			popElem, _ := PopFromStack(&signStack)

			for signStack.LastIndex > 0 && popElem != '(' {
				outExpr += string(popElem)
				popElem, _ = PopFromStack(&signStack)
			}

			PushInStack(&signStack, popElem)
			PushInStack(&signStack, inExpr[i])
			i++
		} else if inExpr[i] == '*' || inExpr[i] == '/' {
			popElem, _ := PopFromStack(&signStack)

			for signStack.LastIndex > 0 && popElem != '(' &&
				popElem != '+' && popElem != '-' {
				outExpr += string(popElem)
				popElem, _ = PopFromStack(&signStack)
			}

			PushInStack(&signStack, popElem)
			PushInStack(&signStack, inExpr[i])
			i++
		} else if unicode.IsDigit(rune(inExpr[i])) {
			outExpr += string(inExpr[i])
			i++

			for unicode.IsDigit(rune(inExpr[i])) {
				outExpr += string(inExpr[i])
				i++
			}
			outExpr += " "
		} else {
			return "", errors.New("error: incorrect expression")
		}
	}

	popElem, _ := PopFromStack(&signStack)

	if popElem != '(' && popElem != ')' {
		outExpr += string(popElem)
	}
	for signStack.LastIndex != 0 {
		popElem, _ = PopFromStack(&signStack)

		if popElem != '(' && popElem != ')' {
			outExpr += string(popElem)
		}
	}

	return outExpr, nil
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	expression := os.Args[1]
	fmt.Println("In expr: ", expression)

	postfixExpr, err := GetPostfixExpression(expression)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Post expr: ", postfixExpr)
}
