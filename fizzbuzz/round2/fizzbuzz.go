package round2

import "fmt"

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

func FizzBuzz(input int) string {
	var result string
	if input%3 == 0 {
		result += Fizz
	}
	if input%5 == 0 {
		result += Buzz
	}
	if result == "" {
		result += fmt.Sprintf("%d", input)
	}
	return result
}
