package tabletest

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct{ A, B, Expected int }{
		{1, 1, 2},
		{1, -1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d + %d", tc.A, tc.B), func(t *testing.T) {
			actual := tc.A + tc.B
			if actual != tc.Expected {
				t.Fatal("failure")
			}
		})
	}
}
