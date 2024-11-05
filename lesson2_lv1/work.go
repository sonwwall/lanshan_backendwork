package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	lst1 := []int{1, 2, 2, 4, 4, 4, 6, 2, 2, 9}
	fmt.Println(calculate(lst1))

}
func calculate(lst []int) map[int]int {
	var num int

	result := map[int]int{}
	for i := 0; i < len(lst); i++ {

		num = lst[i]

		k, exists := result[num]
		if exists {
			result[num] = k + 1
		} else {
			result[num] = 1
		}
	}
	return result
}
