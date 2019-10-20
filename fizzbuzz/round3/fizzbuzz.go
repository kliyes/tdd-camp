package round3

import "fmt"

type FizzBuzz struct {
	input int
}

func (f FizzBuzz) DividedBy(n int) bool {
	return f.input%n == 0
}

func (f FizzBuzz) Convert() string {
	var result string
	if f.DividedBy(3) {
		result += "Fizz"
	}
	if f.DividedBy(5) {
		result += "Buzz"
	}
	if result == "" {
		result += fmt.Sprintf("%d", f.input)
	}
	return result
}
