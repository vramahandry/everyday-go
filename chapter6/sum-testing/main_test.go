package main

import "testing"

func Test_sum(t *testing.T) {
	tables := []struct {
		x    int
		y    int
		want int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}
	for _, table := range tables {
		got := sum(table.x, table.y)
		if got != table.want {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.",
				table.x, table.y, got, table.want)
		}
	}
}
