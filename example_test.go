// example_test.go
// `go test example.go example_test.go`
package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1 - 2", -1},
		{"0", 0},
		{"1", 1},
		{"1 + 1", 2},
		{"1 - 1", 0},
		{"1 * 1", 1},
		{"1 / 1", 1},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"4 / 2", 2},
		{"1 + 2 * 3", 7},
		{"1+2*3", 7},
		{"  1\t+\n2\r*3   ", 7},
		{"123", 123},
		{"123 * 456", 56088},
		{"5 * (6 + 8)", 70},
		{"5 * (6 + 8)/2", 35},
		{"12 * (34 - 56)", -264},
	}

	for _, test := range tests {
		if got := calc(test.input); got != test.want {
			t.Errorf("calc(%q) = %d want %d", test.input, got, test.want)
		}
	}
}

/*

go test -v example.go example_test.go
=== RUN   TestCalc
--- PASS: TestCalc (0.00s)
PASS
ok      command-line-arguments  0.062s

*/
