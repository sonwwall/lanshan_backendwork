package main

import "fmt"

func main() {

	var tem Temperature
	tem = Temperature{
		10,
		10,
	}
	ToFahrenbheit(&tem)
	fmt.Println(tem)
	tem = Temperature{
		20,
		20,
	}
	ToCelsius(&tem)
	fmt.Println(tem)

}

type Temperature struct {
	Fahrenheit float64
	Celsius    float64
}

func ToFahrenbheit(tem *Temperature) {
	tem.Fahrenheit = tem.Celsius*1.8 + 32
}
func ToCelsius(tem *Temperature) {
	tem.Celsius = (tem.Fahrenheit - 32) / 1.8
}
