package round4

import "testing"

func assertEqual(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestFizzBuzz(t *testing.T) {
	t.Run("test_3", func(t *testing.T) {
		fb := FizzBuzz{input: 3}
		want := "Fizz"
		assertEqual(t, want, fb.Convert())
	})
	t.Run("test_13", func(t *testing.T) {
		fb := FizzBuzz{input: 13}
		want := "Fizz"
		assertEqual(t, want, fb.Convert())
	})
	t.Run("test_5", func(t *testing.T) {
		fb := FizzBuzz{input: 5}
		want := "Buzz"
		assertEqual(t, want, fb.Convert())
	})
	t.Run("test_25", func(t *testing.T) {
		fb := FizzBuzz{input: 25}
		want := "Buzz"
		assertEqual(t, want, fb.Convert())
	})
	t.Run("test_4", func(t *testing.T) {
		fb := FizzBuzz{input: 4}
		want := "4"
		assertEqual(t, want, fb.Convert())
	})
}
