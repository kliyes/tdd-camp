package round6

import "testing"

func TestFizzBuzz(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_show_raw_number",
			args{1},
			"1",
		},
		{
			"test_show_raw_number",
			args{2},
			"2",
		},
		{
			"test_show_fizz",
			args{3},
			"Fizz",
		},
		{
			"test_show_buzz",
			args{5},
			"Buzz",
		},
		{
			"test_show_fizzbuzz",
			args{15},
			"FizzBuzz",
		},
		{
			"test_show_fizz_if_contains_3",
			args{13},
			"Fizz",
		},
		{
			"test_show_buzz_if_contains_5",
			args{52},
			"Buzz",
		},
		{
			"test_show_fizzbuzz_if_contains_3_and_5",
			args{51},
			"FizzBuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.input); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
