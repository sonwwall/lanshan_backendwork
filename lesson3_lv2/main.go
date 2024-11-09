package main

import (
	"fmt"
	"lesson3_lv2/utils"
)

func main() {
	result := utils.Reverse("hello")
	fmt.Println(result) //验证Reverse函数正确
	fmt.Println(utils.IsPalindrome("121"))
	fmt.Println(utils.IsPalindrome("12345"))//验证回文函数

}
