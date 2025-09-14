package main

import "fmt"

func main() {

	//* Iterating over data structure

	numArr := []int{12, 34, 22, 66}
	fmt.Println(numArr)

	for i := 0; i < len(numArr); i++ {
		// fmt.Println(numArr[i])
	}

	for index, val := range numArr {
		fmt.Println(index, val)
	}

	person := map[string]string{"name": "rx", "location": "ghy", "gender": "male"}
	for k, v := range person {
		fmt.Println(k, "-", v)
	}

	// unicode code point rune
	// starting byte of rune 
	for i,c := range "golang" {
		fmt.Println(i, string(c))
	}

}
