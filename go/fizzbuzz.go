package main

import "fmt"

// FizzBuzz return fizz, buzz, FizzBuzz or the number
func FizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "fizzbuzz"
	case n%5 == 0:
		return "buzz"
	case n%3 == 0:
		return "fizz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

// N generate empty struct containing specified number of arrays
func N(size int) []struct{} {
	return make([]struct{}, size)
}

const size = 100

func main() {
	for i := range N(size) {
		fmt.Println(FizzBuzz(i + 1))
	}
}
