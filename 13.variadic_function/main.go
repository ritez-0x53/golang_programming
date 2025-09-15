package main

import "fmt"

func sum(nums ...int) int {

	total := 0

	for _, val := range nums {
		total = total + val
	}
	return total
}

func main() {
	res := sum(2, 3, 4, 5, 6, 7, 1, 22)

	nums := []int {22,50,8}
	res2 := sum(nums...)
	fmt.Println(res ,res2)
}
