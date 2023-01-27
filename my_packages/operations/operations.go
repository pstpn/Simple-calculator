package operations

import (
	"errors"
	"my_project/calc/my_packages/stack"
	"strconv"
	"strings"
	"unicode"
)

// GetPostfixExpression - Функция, реализующая преобразование выражения
// из инфиксной формы в постфиксную при помощи стека
func GetPostfixExpression(inExpr string) (string, error) {

	var outExpr string
	var signStack stack.Stack

	signStack.Data = make([]string, 10)
	signStack.LastIndex = 0
	lenInExpr := len(inExpr)

	for i := 0; i < lenInExpr; {
		if inExpr[i] == '(' {
			stack.PushInStack(&signStack, string(inExpr[i]))
			i++
		} else if inExpr[i] == ')' {
			popElem, _ := stack.PopFromStack(&signStack)

			for signStack.LastIndex > 0 && popElem != "(" {
				outExpr += string(popElem) + " "
				popElem, _ = stack.PopFromStack(&signStack)
			}

			if signStack.LastIndex == 0 && popElem != "(" {
				return "", errors.New("error: incorrect expression")
			}
			i++
		} else if inExpr[i] == '-' || inExpr[i] == '+' {
			popElem, _ := stack.PopFromStack(&signStack)

			for signStack.LastIndex > 0 && popElem != "(" {
				outExpr += string(popElem) + " "
				popElem, _ = stack.PopFromStack(&signStack)
			}

			stack.PushInStack(&signStack, popElem)
			stack.PushInStack(&signStack, string(inExpr[i]))
			i++
		} else if inExpr[i] == '*' || inExpr[i] == '/' {
			popElem, _ := stack.PopFromStack(&signStack)

			for signStack.LastIndex > 0 && popElem != "(" &&
				popElem != "+" && popElem != "-" {
				outExpr += string(popElem) + " "
				popElem, _ = stack.PopFromStack(&signStack)
			}

			stack.PushInStack(&signStack, popElem)
			stack.PushInStack(&signStack, string(inExpr[i]))
			i++
		} else if unicode.IsDigit(rune(inExpr[i])) {
			outExpr += string(inExpr[i])
			i++

			for i < lenInExpr && unicode.IsDigit(rune(inExpr[i])) {
				outExpr += string(inExpr[i])
				i++
			}
			outExpr += " "
		} else {
			return "", errors.New("error: incorrect expression")
		}
	}

	popElem, _ := stack.PopFromStack(&signStack)

	if popElem != "(" && popElem != ")" {
		outExpr += string(popElem) + " "
	}
	for signStack.LastIndex != 0 {
		popElem, _ = stack.PopFromStack(&signStack)

		if popElem != "(" && popElem != ")" {
			outExpr += string(popElem) + " "
		}
	}

	return outExpr, nil
}

// CalcPostfixExpr - Функция, реализующая вычисление значения выражения,
// записанного в строке в постфиксном виде при помощи стека
func CalcPostfixExpr(inPostExpr string) (string, error) {

	var numsStack stack.Stack

	numsStack.Data = make([]string, 10)
	numsStack.LastIndex = 0

	answer := 0

	splitExpr := strings.Split(inPostExpr, " ")

	for i := 0; splitExpr[i] != ""; i++ {
		_, err := strconv.Atoi(splitExpr[i])

		if err != nil {
			firstStrNum, numErr := stack.PopFromStack(&numsStack)
			if numErr != nil {
				return "", errors.New("error: incorrect expression")
			}

			secondStrNum, numErr := stack.PopFromStack(&numsStack)
			if numErr != nil {
				return "", errors.New("error: incorrect expression")
			}

			firstNum, _ := strconv.Atoi(firstStrNum)
			secondNum, _ := strconv.Atoi(secondStrNum)

			if splitExpr[i] == "+" {
				answer = firstNum + secondNum
			} else if splitExpr[i] == "-" {
				answer = secondNum - firstNum
			} else if splitExpr[i] == "*" {
				answer = firstNum * secondNum
			} else if splitExpr[i] == "/" {
				answer = secondNum / firstNum
			}

			stack.PushInStack(&numsStack, strconv.Itoa(answer))
		} else {
			stack.PushInStack(&numsStack, splitExpr[i])
		}
	}

	return stack.PopFromStack(&numsStack)
}
