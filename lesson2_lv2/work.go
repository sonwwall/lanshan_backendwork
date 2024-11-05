package main

import "fmt"

func main() {
	results := Calculator("multiply")

	fmt.Println(results(4, 5))
}
func Calculator(optype string) func(int, int) int {
	switch optype {
	case "add":
		return func(x, y int) int {
			return x + y
		}
	case "subtract":
		return func(x, y int) int {
			return x - y
		}
	case "multiply":
		return func(x, y int) int {
			return x * y

		}
	case "divide":
		return func(x, y int) int {
			return x / y
		}
	default:
		return nil
	}
}
