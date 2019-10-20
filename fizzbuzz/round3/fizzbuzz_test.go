package round3

import "testing"

func assertEqual(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestFizzBuzz(t *testing.T) {
	t.Run("test_3", func(t *testing.T) {
		f := FizzBuzz{input: 3}
		want := "Fizz"
		assertEqual(t, want, f.Convert())
	})

	t.Run("test_5", func(t *testing.T) {
		f := FizzBuzz{input: 5}
		want := "Buzz"
		assertEqual(t, want, f.Convert())
	})

	t.Run("test_15", func(t *testing.T) {
		f := FizzBuzz{input: 15}
		want := "FizzBuzz"
		assertEqual(t, want, f.Convert())
	})

	t.Run("test_4", func(t *testing.T) {
		f := FizzBuzz{input: 4}
		want := "4"
		assertEqual(t, want, f.Convert())
	})
}
