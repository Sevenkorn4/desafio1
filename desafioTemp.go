package main

import "fmt"

const temperaturaF = 212.0
const kelvinBase = 273.15

func main() {

	tempC := (temperaturaF - 32) * 5 / 9
	tempK := tempC + kelvinBase

	fmt.Println("A temperatura em Fahrenheit é:", temperaturaF)
	fmt.Println("A temperatura em Celsius é:", tempC)
	fmt.Println("A temperatura em Kelvin é:", tempK)
}
