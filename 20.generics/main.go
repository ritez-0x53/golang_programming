package main

import "fmt"

func PrintSlice[T comparable](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

type Stack[T interface{}] struct {
	elements []T
}

func main() {

	items1 := []int{1, 2, 3, 4, 5, 6}
	items2 := []string{"golang", "javascript", "c++", "arduino"}

	PrintSlice(items1)
	PrintSlice(items2)

	myStack := Stack[string]{
		elements: []string{"Golang", "TypeScript"},
	}

	fmt.Println(myStack)
	fmt.Println(myStack.elements)

}
