// Pointers in golang

package main

import "fmt"

func main() {
	val1 := 40
	val2 := val1
	fmt.Println(&val1, &val2) //* memory addresses are different

	slc1 := []int{1, 2, 3, 4}
	slc2 := slc1
	fmt.Println(&slc1[0], &slc2[0]) //* memory addresses are same ,it means pointing the same reference

	// regualar variable
	var r int = 25
	fmt.Println(r, &r)

	// pointer variable
	var p *int = &r
	fmt.Println(*p, p)
	*p = 44
	fmt.Println(r, *p)
}
