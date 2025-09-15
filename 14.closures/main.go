package main

import "fmt"

// closures in golang

func counter() func(inc_num int) int {
	var count int = 100
	return func(inc_num int) int {
		count += inc_num
		return count
	}
}

func main() {

	res := counter()
	fmt.Println(res(11)) //111
	fmt.Println(res(11)) //122
	fmt.Println(res(11)) //133

}
