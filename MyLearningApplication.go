package main

import (
	"Go-lang/first_program"
	"Go-lang/utils"
	"fmt"
)

func main() {
	fmt.Println("MAIN said:: Hello, World!")
	fmt.Println(first_program.TestPrint())

	result := utils.Filter([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(result)
}
