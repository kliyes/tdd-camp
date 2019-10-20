package round4

import (
	"fmt"
	"strings"
)

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz) Contains(i string) bool {
	return strings.Contains(fmt.Sprintf("%d", fb.input), i)
}

func (fb FizzBuzz) DividedBy(n int) bool {
	return fb.input%n == 0
}

func (fb FizzBuzz) Convert() string {
	var result string
	if fb.DividedBy(3) || fb.Contains("3") {
		result += "Fizz"
	}
	if fb.DividedBy(5) || fb.Contains("5") {
		result += "Buzz"
	}
	if result == "" {
		result += fmt.Sprintf("%d", fb.input)
	}
	return result
}
