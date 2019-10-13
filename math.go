package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func main() {
	fmt.Println(add(13, 12))
	fmt.Println(divide(8, 2))
	fmt.Println(multiply(3, 3))
}
