package round6

import (
	"fmt"
	"strings"
)

func FizzBuzz(input int) string {
	if isRelatedTo(input, 3) && isRelatedTo(input, 5) {
		return "FizzBuzz"
	}
	if isRelatedTo(input, 3) {
		return "Fizz"
	}
	if isRelatedTo(input, 5) {
		return "Buzz"
	}
	return fmt.Sprintf("%d", input)
}

func isRelatedTo(input, n int) bool {
	if input%n == 0 || strings.Contains(fmt.Sprintf("%d", input), fmt.Sprintf("%d", n)) {
		return true
	}
	return false
}
