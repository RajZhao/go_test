package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}
func div(a, b int) (int, int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	//fmt.Println(eval(3, 4, "*"))
	//fmt.Println(div(13, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
