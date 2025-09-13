package main

import "fmt"

func canHalve(x int) (int, bool) {
	half := x / 2
	if x%2 == 0 {
		return half, true
	} else {
		return half, false
	}
}

func greatestNum(nums ...int) int {
	var greatest int

	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}

	return greatest
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		i += 2
		return i
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	fmt.Println(canHalve(1))
	fmt.Println(canHalve(2))

	fmt.Println(greatestNum(1, 2, 3, 4, 5))

	oddGenerator := makeOddGenerator()
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())

	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
