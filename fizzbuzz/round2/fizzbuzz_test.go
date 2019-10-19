package round2

import "testing"

func assertEqual(t *testing.T, want, got string) {
	t.Helper()

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestFizzBuzz(t *testing.T) {
	t.Run("test_3", func(t *testing.T) {
		assertEqual(t, Fizz, FizzBuzz(3))
	})

	t.Run("test_5", func(t *testing.T) {
		assertEqual(t, Buzz, FizzBuzz(5))
	})

	t.Run("test_15", func(t *testing.T) {
		assertEqual(t, Fizz+Buzz, FizzBuzz(15))
	})

	t.Run("test_4", func(t *testing.T) {
		assertEqual(t, "4", FizzBuzz(4))
	})
}
