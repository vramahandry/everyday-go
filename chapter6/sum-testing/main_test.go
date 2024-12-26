package main

import "testing"

func Test_sum_odd_numbers(t *testing.T) {
	x := 5
	y := 5
	want := 10
	got := sum(x, y)
	t.Logf("Testing a sum of: %d %d", x, y)
	if want != got {
		t.Errorf("Sum was incorrect want: %d, but got: %d", want, got)
	}
}
func Test_sum_negative_numbers(t *testing.T) {
	x := -5
	y := -5
	want := -10
	got := sum(x, y)
	t.Logf("Testing a sum of: %d %d", x, y)
	if want != got {
		t.Errorf("Sum was incorrect want: %d, but got: %d", want, got)
	}
}
