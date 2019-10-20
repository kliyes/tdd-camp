package round5

import (
	"fmt"
	"strings"
)

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz) Convert() string {
	if fb.isRelatedTo(3) && fb.isRelatedTo(5) {
		return "FizzBuzz"
	}
	if fb.isRelatedTo(3) {
		return "Fizz"
	}
	if fb.isRelatedTo(5) {
		return "Buzz"
	}
	return fmt.Sprintf("%d", fb.input)
}

func (fb FizzBuzz) isRelatedTo(n int) bool {
	return fb.input%n == 0 || strings.Contains(fmt.Sprintf("%d", fb.input), fmt.Sprintf("%d", n))
}
