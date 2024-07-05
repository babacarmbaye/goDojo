package main

import "fmt"

func add(x, y int) int {
	//panic("unimplemented")
	return x + y
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum_some_nums(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Basic function")
	result := add(4, 3)
	fmt.Println("result =", result)

	fmt.Println("Multiple return values")
	a, b := swap("hello", "world")
	fmt.Println("a, b =", a, b)

	fmt.Println("Named return values")
	sum := 100
	x, y := split(sum)
	fmt.Printf("The split of %d is: x=%d, y=%d", sum, x, y)

	fmt.Println("Variadic function")
	sum_result := sum_some_nums(1, 2, 3, 4, 5)
	fmt.Println("sum_result =", sum_result)
}
