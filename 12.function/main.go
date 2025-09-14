package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLangs() (string, string, string) {
	return "c", "go", "js"
}

func main() {
	res := add(2, 5)
	fmt.Println(res)

	lang1, lang2, lang3 := getLangs()
	fmt.Println(lang1, lang2, lang3)

}
