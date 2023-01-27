package stack

import "errors"

// Stack - Тип данных, реализующий стек
type Stack struct {
	LastIndex int
	Data      []string
}

// PushInStack - Функция добавления элемента в стек
func PushInStack(stack *Stack, elem string) {

	stack.Data[stack.LastIndex] = elem
	stack.LastIndex++
}

// PopFromStack - Функция удаления элемента из стека
func PopFromStack(stack *Stack) (string, error) {

	if stack.LastIndex == 0 {
		return "", errors.New("error: empty stack")
	}

	stack.LastIndex--
	return stack.Data[stack.LastIndex], nil
}
