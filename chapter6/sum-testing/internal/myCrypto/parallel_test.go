package myCrypto

import (
	"testing"
	"time"
)

func Test_Slow(t *testing.T) {
	cases := []struct {
		Name     string
		Duration time.Duration
	}{
		{Name: "Case 1", Duration: time.Second * 1},
		{Name: "Case 2", Duration: time.Second * 1},
		{Name: "Case 3", Duration: time.Second * 1},
		{Name: "Case 4", Duration: time.Second * 1},
		{Name: "Case 5", Duration: time.Second * 1},
	}
	for _, tc := range cases {

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			t.Logf("%s sleeping..", tc.Name)
			time.Sleep(tc.Duration)
			t.Logf("%s slept", tc.Name)

		})

	}
}
