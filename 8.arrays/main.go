package main

import "fmt"

func main() {

	// Zeroid values
	// int -> 0 , string -> "" , bool -> false

	// number sequenced of specific element

	var nums [5]int
	fmt.Println(nums)
	// array length
	fmt.Println(len(nums))
	nums[0] = 23
	fmt.Println(nums)

	var strs [4]string
	strs[1] = "hi"
	fmt.Println(strs)

	var vals [5]bool
	fmt.Println(vals)
	vals[2] = true
	fmt.Println(vals)

	// declaring in a single line
	langs := [3]string{"go","TS","python"}
	fmt.Println(langs)

	// 2d array
	arr2d := [2][2]int{{1,2},{3,4}} 
	fmt.Println(arr2d)

}
