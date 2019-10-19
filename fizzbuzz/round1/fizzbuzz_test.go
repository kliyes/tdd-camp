package round1

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("test_3", func(t *testing.T) {
		want := "Fizz"
		got := FizzBuzz(3)
		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("test_5", func(t *testing.T) {
		want := "Buzz"
		got := FizzBuzz(5)
		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("test_15", func(t *testing.T) {
		want := "FizzBuzz"
		got := FizzBuzz(15)
		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("test_others", func(t *testing.T) {
		want := "4"
		got := FizzBuzz(4)
		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
