package round5

import "testing"

func assertEqual(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestFizzBuzz(t *testing.T) {
	t.Run("test_show_raw_number", func(t *testing.T) {
		fb := FizzBuzz{1}
		want := "1"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
	t.Run("test_show_raw_number", func(t *testing.T) {
		fb := FizzBuzz{2}
		want := "2"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
	t.Run("test_show_fizz_if_divisible_by_3", func(t *testing.T) {
		fb := FizzBuzz{3}
		want := "Fizz"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
	t.Run("test_show_buzz_if_divisible_by_5", func(t *testing.T) {
		fb := FizzBuzz{5}
		want := "Buzz"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
	t.Run("test_show_fizzbuzz_if_divisible_by_3_and_5", func(t *testing.T) {
		fb := FizzBuzz{15}
		want := "FizzBuzz"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
	t.Run("test_show_fizz_or_buzz_if_contains_3_or_5", func(t *testing.T) {
		fb := FizzBuzz{13}
		want := "Fizz"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
	t.Run("test_show_fizz_or_buzz_if_contains_3_or_5", func(t *testing.T) {
		fb := FizzBuzz{52}
		want := "Buzz"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
	t.Run("test_show_fizz_or_buzz_if_contains_3_or_5", func(t *testing.T) {
		fb := FizzBuzz{53}
		want := "FizzBuzz"
		got := fb.Convert()
		assertEqual(t, want, got)
	})
}
