package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLangs() (string, string, string) {
	return "c", "go", "js"
}

func logName(name string) {
	fmt.Println(name)
}

// * passing function as a parameter like in JS - callback
func processIt(fn func(a int) int, val int) {
	result := fn(val)
	fmt.Println(result)
}

func myFunc() func(a int) int {
	return func (a int) int  {
		return a
	}
} 


func main() {
	res := add(2, 5)
	fmt.Println(res)

	lang1, lang2, lang3 := getLangs()
	fmt.Println(lang1, lang2, lang3)
	logName("riteswar")

	fn := func(val int) int {
		return val
	}
	processIt(fn, 6000)

	run := myFunc()
	res1 := run(200)
	fmt.Println(res1)

}
