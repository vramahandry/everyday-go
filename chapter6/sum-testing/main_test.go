package main

import "testing"

func Test_sum(t *testing.T) {
	x := 5
	y := 5
	want := 10
	got := sum(x, y)
	if want != got {
		t.Logf("Testing a sum of: %d %d", x, y)
		t.Errorf("Sum was incorrect want: %d, but got: %d", want, got)

	}
}
